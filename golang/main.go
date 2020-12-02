package main

import (
	"fmt"

	"golang/Config"
	"golang/Models"
	"golang/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDbConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()
	r.Run()
}
