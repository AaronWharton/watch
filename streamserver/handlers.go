package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid

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

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
