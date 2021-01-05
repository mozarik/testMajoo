package models

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		ID          int    `gorm:"primary_key;auto_increment" json:"id"`
		Username    string `gorm:"size:255;not null;unique" json:"username"`
		Password    string `gorm:"size:100;not null;" json:"password"`
		NamaLengkap string `gorm:"size:255;not null" json:"nama_lengkap"`
	}
)

// TODO Implement Has()
func (u *User) BeforeSave() error {
	passowordHashed, err := Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(passowordHashed)
	return nil
}

func (u *User) Prep() {
	u.ID = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.NamaLengkap = html.EscapeString(strings.TrimSpace(u.NamaLengkap))
}

func (u *User) Validasi(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			return errors.New("Butuh Username")
		}
		if u.Password == "" {
			return errors.New("Butuh Password")
		}
		if u.NamaLengkap == "" {
			return errors.New("Butuh Nama Lengkap")
		}

		return nil

	case "login":
		if u.Password == "" {
			return errors.New("Butuh Password")
		}
		if u.Username == "" {
			return errors.New("Butuh Email")
		}
		return nil

	default:
		if u.Username == "" {
			return errors.New("Butuh Email")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {

	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {

	err := db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":     u.Password,
			"username":     u.Username,
			"nama_lengkap": u.NamaLengkap,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
