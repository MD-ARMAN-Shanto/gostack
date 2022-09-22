package seed

import (
	"github.com/MD-ARMAN-Shanto/gostack/api/models"
	"github.com/jinzhu/gorm"
	"log"
)

var users = []models.User{
	models.User{
		Nickname: "Arman Shanto",
		Email:    "shanto377@gmail.com",
		Password: "arman",
	},
	models.User{
		Nickname: "Arman Hossain",
		Email:    "Hossain.arman@northsouth.edu",
		Password: "arman1",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Titel 1",
		Content: "Hello world",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Welcome to go world",
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate tables: %v", err)
	}
	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorId = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
