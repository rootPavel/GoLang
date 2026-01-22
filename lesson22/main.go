package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"uniqueIndex"`
	LastName  string `gorm:"uniqueIndex"`
	Email     string `gorm:"not null"`
	Country   string `gorm:"not null"`
	Age       int    `gorm:"not null;size:3"`
}

func main() {
	pgConnect := "host=localhost user=postgres password=password dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(pgConnect), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
		return
	}

	newUser := User{
		FirstName: "Leha",
		LastName:  "Lehin",
		Email:     "Leha@mail.ru",
		Country:   "Russia",
		Age:       29,
	}

	result := db.Create(&newUser)
	if result.Error != nil {
		log.Println("Error create Serega", result.Error)
		return
	}

	// var users []User
	// result := db.Where("first_name = ?", "Leha").Find(&users)
	// if result.Error != nil {
	// 	log.Println("Error - ", result.Error)
	// 	return
	// }
	// fmt.Println(users)

	// var user User
	// result := db.First(&user)
	// if result.Error != nil {
	// 	log.Println("Error - ", result.Error)
	// 	return
	// }
	// user.Age = 35
	// result = db.Save(&user)
	// if result.Error != nil {
	// 	log.Println("Error - ", result.Error)
	// 	return
	// }

	// var user User
	// result := db.First(&user)
	// if result.Error != nil {
	// 	log.Println("Error - ", result.Error)
	// 	return
	// }
	// result = db.Delete(&user)
	// if result.Error != nil {
	// 	log.Println("Error - ", result.Error)
	// } else if result.RowsAffected == 0 {
	// 	log.Println("No rows affexted")
	// } else {
	// 	log.Println("Deleted susessfully")
	// }
}
