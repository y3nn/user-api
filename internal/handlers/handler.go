package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type Status_Response struct { 
	Status string `json:"Stauts"`
}
type Hello_Response struct  { 
	Status string `json:"Status"`
}

type Time_Response struct { 
	Time time.Time `json:"Time"`
}


func HelloHandler(w http.ResponseWriter, r * http.Request) {
	res := Hello_Response{Status: "Hello 🔥"}
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(&res)
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	res := Time_Response{Time: time.Now()}
	w.Header().Add("Content-Type","application/json")
	w.Write([]byte("⏰ Time now"))
	json.NewEncoder(w).Encode(&res)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) { 
	res := Status_Response{Status: "OK[200]"}
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(&res)
}
