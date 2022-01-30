package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var accessToken = os.Getenv("ACCESS_TOKEN")
var roomID = os.Getenv("ROOM_ID")

func main() {

	http.HandleFunc("/_wedding/submit", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			rw.WriteHeader(200)
			return
		}
		if err := r.ParseMultipartForm(50*1024); err != nil {
			rw.WriteHeader(500)
			log.Printf("failed to parse form: %s", err)
			return
		}
		body, err := json.Marshal(r.PostForm)
		if err != nil {
			rw.WriteHeader(500)
			log.Printf("failed to marshal: %s", err)
			return
		}
		fmt.Printf("received form %+v \n", string(body))
		now := time.Now()
		timestr := now.Format("2006-01-02-15-04-05.000")
		filename := fmt.Sprintf("./forms/%s.json", timestr)
		if err = ioutil.WriteFile(filename, body, os.ModePerm); err != nil {
			rw.WriteHeader(500)
			log.Printf("failed to write: %s", err)
			return
		}

		client := http.Client{
			Timeout: 10 * time.Second,
		}
		matrixURL := fmt.Sprintf(
			"https://matrix.org/_matrix/client/v3/rooms/%s/send/m.room.message/%s",
			url.PathEscape(roomID), url.PathEscape(filename),
		)
		matrixBody := map[string]interface{}{
			"msgtype": "m.notice",
			"data": string(body),
			"body": fmt.Sprintf(
				"New wedding submission: %v coming? %v", r.PostForm["email"], r.PostForm["ifcoming"],
			),
		}
		matrixBodyJSON, err := json.Marshal(matrixBody)
		if err != nil {
			rw.WriteHeader(500)
			log.Printf("failed to marshal matrix message: %s", err)
			return
		}
		notif, err := http.NewRequest("PUT", matrixURL, bytes.NewBuffer(matrixBodyJSON))
		if err != nil {
			rw.WriteHeader(500)
			log.Printf("failed to create matrix request: %s", err)
			return
		}
		notif.Header.Set("Authorization", "Bearer " + accessToken)
		notif.Header.Set("Content-Type", "application/json")
		res, err := client.Do(notif)
		if err != nil {
			rw.WriteHeader(500)
			log.Printf("failed to send matrix request: %s", err)
			return
		}
		if res.StatusCode >= 300 {
			rw.WriteHeader(500)
			log.Printf("failed to send matrix request: http %s", res.Status)
			return
		}
		rw.WriteHeader(200)
	})
	fmt.Println("listening")
	if err := http.ListenAndServe("0.0.0.0:1906", nil); err != nil {
		panic(err)
	}
}