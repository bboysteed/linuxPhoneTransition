package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"linuxPhoneTransition/models"
)

func ConnectMysql() {
	var err error
	db, err = gorm.Open("mysql", "root:root@/chatdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("open mysql failed with err: %v\n",err)
	}else{
		//defer db.Close()
		db.AutoMigrate(models.Chat{})
		fmt.Println("success connect mysql db!")
	}

}
