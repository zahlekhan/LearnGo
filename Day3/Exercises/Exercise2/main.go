package main

import (
	"awesomeProject1/LearnGo/Day3/Exercises/Exercise2/Config"
	"awesomeProject1/LearnGo/Day3/Exercises/Exercise2/Models"
	"awesomeProject1/LearnGo/Day3/Exercises/Exercise2/Routes"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	var err error
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Status:", err)
	}

	student := Models.Student{
		Model:    gorm.Model{ID: 2},
		Name:     "jj",
		LastName: "lk",
		DOB:      "2019-10-6",
		Address:  "mads",
		Scores: []Models.Score{
			{
				SubjectID: 1,
				Marks:     90,
			},
			{
				SubjectID: 2,
				Marks:     68,
			},
		},
	}
	_ = Config.DB.AutoMigrate(&Models.Subject{})
	_ = Config.DB.AutoMigrate(&Models.Student{})
	_ = Config.DB.AutoMigrate(&Models.Score{})

	_ = Config.DB.Create(&student)

	r := Routes.SetupRouter()
	//running
	_ = r.Run()
}
