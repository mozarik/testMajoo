package models

type (
	User struct {
		ID          int    `gorm:"primary_key;auto_increment" json:"id"`
		Username    int    `gorm:"size:255;not null;unique" json:"username"`
		Password    string `gorm:"size:100;not null;" json:"password"`
		NamaLengkap string `gorm:"size:255;not null" json:"nama_lengkap"`
	}
)
