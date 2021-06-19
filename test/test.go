package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//func main() {
//	//fmt.Println(commom.GetFileList("./files"))
//
//	db, err := gorm.Open("mysql", "root:root@/chatdb?charset=utf8&parseTime=True&loc=Local")
//	defer db.Close()
//	if err != nil {
//		fmt.Printf("open mysql failed with err: %v\n", err)
//	}
//	db.AutoMigrate(&models.Chat{})
//
//
//	db.Create(&models.Chat{
//		Content: "8888",
//		CreateTime: time.Now().Unix(),
//		Role:    "mobile",
//	})
//	//msgs := struct {
//	//	msg []models.Chat
//	//}{}
//	msgs := make([]models.Chat,0)
//	db.Where("creattime > ?", time.Now().Unix()-30*60).Find(&msgs)
//	fmt.Printf("msgs is :%v\n",msgs)
//	msg := models.Chat{}
//	db.Last(&msg)
//	fmt.Println(msg)
//
//}
