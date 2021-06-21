package main

import (

	"encoding/json"
	"fmt"
	"io"
	"linuxPhoneTransition/commom"
	"linuxPhoneTransition/models"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//func chatWsHandler(w http.ResponseWriter, r *http.Request) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Printf("ws upgrade filed,err is: %v", err)
//		return
//	}
//	var CurMsg = models.Chat{}
//	conn.ReadJSON(&CurMsg)
//	if CurMsg.Content == "" {
//		return
//	}
//	CurMsg.CreateTime = time.Now().Unix()
//	db.Create(&CurMsg)
//	conn.WriteJSON(&struct {
//		Msg string `json:"msg"`
//	}{
//		Msg: "send ok",
//	})
//}

func getMsgsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	msgs := make([]models.Chat, 0)
	db.Where("creattime > ?", time.Now().Unix()-30*60).Find(&msgs)
	//fmt.Printf("[mysql] find records are :%v\n",msgs)
	resp, err := json.Marshal(msgs)
	if err != nil {
		fmt.Printf("mashall msgs failed with err:%v\n", err)
		return
	}
	w.Write(resp)
}

//func sendMsgHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	if r.Method != "POST" {
//		remsg, _ := json.Marshal(&models.MsgST{
//			M: "error",
//			C: 500,
//			R: "请求方式错误",
//		})
//		w.Write(remsg)
//		return
//	}
//	var CurMsg = models.Chat{}
//	msg, _ := ioutil.ReadAll(r.Body)
//	err := json.Unmarshal(msg, &CurMsg)
//	if err != nil {
//		fmt.Printf("unMashal msg from http body failed with err :%v\n", err)
//		return
//	}
//	CurMsg.CreateTime = time.Now().Unix()
//	db.Create(&CurMsg)

//}

func getFileHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fileList, err := json.Marshal(commom.GetFileList(storePath))
	if err != nil {
		log.Printf("json marshal err,err is: %v", err)
	}
	w.Write(fileList)

}

func uploadFileHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	contentType := req.Header.Get("content-type")
	contentLen := req.ContentLength
	//fmt.Printf("upload content-type:%s\tcontent-length:%d\n", contentType, contentLen)
	if !strings.Contains(contentType, "multipart/form-data") {
		w.Write([]byte("content-type must be multipart/form-data"))
		return
	}
	if contentLen >= 1000*1024*1024 { // 1000 MB
		w.Write([]byte("file to large,limit 100MB"))
		return
	}
	err := req.ParseMultipartForm(100 * 1024 * 1024)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("ParseMultipartForm error:" + err.Error()))
		return
	}

	if len(req.MultipartForm.File) == 0 {
		w.Write([]byte("not have any file"))
		return
	}

	for _, files := range req.MultipartForm.File {
		//fmt.Printf("req.MultipartForm.File,name=%s\n", name)
		for _, f := range files {
			handle, err_ := f.Open()
			if err_ != nil {
				w.Write([]byte(fmt.Sprintf("unknown error,fileName=%s,fileSize=%d,err:%s", f.Filename, f.Size, err_.Error())))
				return
			}
			abStorePath,_:=filepath.Abs(storePath)
			os.MkdirAll(abStorePath,0777)
			path := filepath.Join(storePath, f.Filename)
			dst, err1 := os.Create(path)
			if err1 != nil {
				fmt.Printf("create file failed,err is :%v\n", err1)
			}
			_, cpErr := io.Copy(dst, handle)
			if cpErr != nil {
				fmt.Printf("io copy failed,err is : %v", cpErr)
			}
			dst.Close()
			log.Printf("[-UPLOAD-]: FILE: %s\tSIZE: %.2f MB\tSAVE_PATH: %s\n", f.Filename, float64(contentLen)/1024/1024, path)
			remsg, _ := json.Marshal(&models.MsgST{
				M: "ok",
				C: 200,
				R: "",
			})
			w.Write(remsg)

		}

	}
}

func downloadFileHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//if req.RequestURI == "/favicon.ico" {
	//	return
	//}
	filename, _ := url.QueryUnescape(strings.Replace(req.RequestURI, "/download", "", -1))

	log.Printf("[-DOWNLOAD-] FILE：%v", filename)

	enEscapeUrl, err := url.QueryUnescape(filename[1:])
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	f, err := os.Open("./files/" + enEscapeUrl)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	info, err := f.Stat()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	_, contentType := getContentType(filename)
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	//w.Header().Set("Content-Type", http.DetectContentType(fileHeader))
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", strconv.FormatInt(info.Size(), 10))

	f.Seek(0, 0)
	io.Copy(w, f)
}

func getContentType(fileName string) (extension, contentType string) {
	arr := strings.Split(fileName, ".")

	// see: https://tool.oschina.net/commons/
	if len(arr) >= 2 {
		extension = arr[len(arr)-1]
		switch extension {
		case "jpeg", "jpe", "jpg":
			contentType = "image/jpeg"
		case "png":
			contentType = "image/png"
		case "gif":
			contentType = "image/gif"
		case "mp4":
			contentType = "video/mpeg4"
		case "mp3":
			contentType = "audio/mp3"
		case "wav":
			contentType = "audio/wav"
		case "pdf":
			contentType = "application/pdf"
		case "doc", "":
			contentType = "application/msword"
		}
	}
	// .*（ 二进制流，不知道下载文件类型）
	contentType = "application/octet-stream"
	return
}

func allowOrigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	r.ParseForm()
	fmt.Println("收到客户端请求: ", r.Form)
}
