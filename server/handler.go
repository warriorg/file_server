package server

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
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

	fileInfo := FileInfo{
		Name: header.Filename,
		Type: ext,
		Data: data,
		Size: len(data),
		Date: time.Now()}

	Save(&fileInfo)

	fmt.Fprintf(w, "File uploaded successfully : ")
	fmt.Fprintf(w, fileInfo.Key)

}

func Load(w http.ResponseWriter, r *http.Request) {

	key := mux.Vars(r)["key"]

	width := r.FormValue("width")
	height := r.FormValue("height")
	quality := r.FormValue("quality")

	fmt.Println(width)

	file := Find(key)

	if file.Key == "" {
		fmt.Fprintln(w, "Resource is not found")
		return
	}

	w.Header().Set("Content-Type", "image/"+strings.TrimLeft(file.Type, "."))
	w.Header().Set("Content-Length", strconv.Itoa(file.Size))
	w.Header().Set("File-Name", file.Name)

	if width != "" || height != "" || quality != "" {

		// newdx, uerr := strconv.Atoi(url[1])

		originalImage, _, err := image.Decode(bytes.NewReader(file.Data))
		if err != nil {
			fmt.Println(err)
		}

		targetWidth := 100
		targetHeight := 100

		dst := resize.Resize(uint(targetWidth), uint(targetHeight), originalImage, resize.Lanczos3)

		jpeg.Encode(w, dst, &jpeg.Options{JpegMaximumQuality})
	} else {
		w.Write(file.Data)
	}
}
