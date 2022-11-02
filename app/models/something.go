package models

import (
	"gin_sample_project/app/config"

	"gorm.io/gorm"
)

/* TODO: replace all "something"s with an actual object as an example */
type Something struct {
	gorm.Model
	Number  int    `gorm:"" json:"number"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

var database *gorm.DB

func init() {
	config.Connect()
	database = config.GetDatabase()
	database.AutoMigrate(&Something{})
}

// TODO: proper error handling for all the functions below

func (something *Something) Create() (uint, error) {
	result := database.Create(&something)
	return something.ID, result.Error
}

func GetAllSomethings() []Something {
	var Somethings []Something
	database.Find(&Somethings)
	return Somethings
}

func GetSomethingByID(ID uint) (*Something, *gorm.DB) {
	var getSomething Something
	db := database.Where("ID = ?", ID).Find(&getSomething)
	return &getSomething, db
}

func (something *Something) Update() {
	database.Save(&something)
}

func (something *Something) Delete() {
	database.Delete(something)
}
