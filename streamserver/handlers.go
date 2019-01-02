package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/gommon/log"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// server->client
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VideoDir + vid

	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error!")
		return
	}

	defer video.Close()

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
}

// client->server
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big!")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error!")
		return
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error : %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error!")
		return
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VideoDir+fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error!")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Upload file successfully!")
}
