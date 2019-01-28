package main

import (
		"fmt"
		"net/http"
		"github.com/jinzhu/gorm"
		_ "github.com/mattn/go-sqlite3"
		"github.com/qor/admin"
       )

var db *gorm.DB                                         // declaring the db globally
var err error
var loguser string

type Login struct {
	Name string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		City string `json:"city"`
}

type Ques struct {
	Genre string `json:"genre"`
		Question string `json:"question"`
		Opt1 string `json:"opt1"`
		Opt2 string `json:"opt2"`
		Opt3 string `json:"opt3"`
}

type Answers struct {
	Ans1 string `json:"ans1"`
		Ans2 string `json:"ans2"`
		Ans3 string `json:"ans3"`
		Ans4 string `json:"ans4"`
		Ans5 string `json:"ans5"`
		Ans6 string `json:"ans6"`
		Ans7 string `json:"ans7"`
		Ans8 string `json:"ans8"`
}

type HighScores struct {
	Username string `json:"username"`
		DN int `json:"dn"`
		CG int `json:"cg"`
		IP int `json:"ip"`
		USP int `json:"usp"`
}

func main() {
	db, _ := gorm.Open("sqlite3", "../src/gorm.db")
		db.AutoMigrate(&Login{}, &Ques{} , &HighScores{})

		// Initalize
		Admin := admin.New(&admin.AdminConfig{DB: db})

		// Allow to use Admin to manage User, Product
		Admin.AddResource(&Login{})
		Admin.AddResource(&Ques{})
		Admin.AddResource(&HighScores{})

		// initalize an HTTP request multiplexer
		mux := http.NewServeMux()

		// Mount admin interface to mux
		Admin.MountTo("/admin", mux)

		fmt.Println("Listening on: 9000")
		http.ListenAndServe(":9000", mux)
}
