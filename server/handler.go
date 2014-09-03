package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
)

// func setCacheHeaders(file *mgo.GridFile, w http.ResponseWriter) {
// 	w.Header().Set("Etag", file.MD5())
// 	w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", ImageCacheDuration))
// 	d, _ := time.ParseDuration(fmt.Sprintf("%ds", ImageCacheDuration))
//
// 	expires := file.UploadDate().Add(d)
//
// 	w.Header().Set("Last-Modified", file.UploadDate().Format(time.RFC1123))
// 	w.Header().Set("Expires", expires.Format(time.RFC1123))
// 	w.Header().Set("Date", file.UploadDate().Format(time.RFC1123))
// }

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi, welcome my server")
}

func Upload(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("file") // the FormFile function takes in the POST input id file
	defer file.Close()

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	ext := filepath.Ext(header.Filename)

	fileInfo := FileInfo{Name: header.Filename, Type: ext, Data: data, Size: len(data), Date: time.Now()}
	Save(fileInfo)

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, header.Filename)

}

func Load(w http.ResponseWriter, r *http.Request) {

	key := mux.Vars(r)["key"]

	fmt.Println(key)

	file := Find(key)

	w.Header().Set("Content-Type", "image/gif")

	w.Write(file.Data)
}
