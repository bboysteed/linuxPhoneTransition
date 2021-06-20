package commom

import (
	"fmt"
	"io/ioutil"
	"linuxPhoneTransition/models"
	"log"
	"os"
	"path/filepath"
)

func GetFileList(path string) []models.File {
	fileList := make([]models.File, 0)
	fs, _ := ioutil.ReadDir(path)
	for _, file := range fs {
		if file.IsDir() {
			continue
			//fmt.Println(path+file.Name())
			//GetFileList(path+file.Name()+"/")
		} else {
			//fmt.Println(path + file.Name())
			fileList = append(fileList, models.File{
				FileName: file.Name(),
				Size:     GetFileSize(filepath.Join(path, file.Name())),
			})
		}
	}
	return fileList
}

//return xx MB
func GetFileSize(path string) string {
	absPath, _ := filepath.Abs(path)
	fn, err := os.Stat(absPath)
	if err != nil {
		log.Printf("file stat err,err is: %v\n", err)
	}
		return fmt.Sprintf("%.1f MB", float32(fn.Size())/1024/1024)
}
