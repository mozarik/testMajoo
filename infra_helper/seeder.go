package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/mozarik/testMajoo/api_model/models"
)

var users = []models.User{
	models.User{
		Username:    "mozarik",
		Password:    "password1",
		NamaLengkap: "Zein Fahrozi",
	},
	models.User{
		Username:    "mozarik2",
		Password:    "password2",
		NamaLengkap: "Zein Fahrozi2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
