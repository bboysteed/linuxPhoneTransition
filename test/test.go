package main

import (
	"github.com/creack/pty"
	"io"
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("/bin/bash")
	f, err := pty.Start(c)
	if err != nil {
		panic(err)
	}

	//go func() {
		f.Write([]byte("ls\r"))
		f.Write([]byte("pwd\n"))
		f.Write([]byte("whoami\n"))
		//f.Write([]byte{4}) // EOT
	//}()
	io.Copy(os.Stdout, f)

















	//var err error
	//_, err = gorm.Open("mysql", "root:root@tcp(192.168.50.49:30006)/chatdb?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	fmt.Printf("open mysql failed with err: %v\n",err)
	//	panic("connect mysql failed!")
	//}else{
	//
	//}

	//response, err := http.Get("http://linux2phone:30000")
	//if err != nil {
	//	log.Print("err is",err)
	//	return
	//}
	//robots, err1 := io.ReadAll(response.Body)
	//response.Body.Close()
	//if err1 != nil {
	//	log.Fatal(err1)
	//}
	//log.Printf("%s", robots)
}
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
