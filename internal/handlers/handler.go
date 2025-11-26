package handler

import (
	"encoding/json"
	"fmt"
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
	format :=  time.Now()
	res := Time_Response{Time: format}
	w.Header().Add("Content-Type","application/json")
	w.Write([]byte("⏰ Time now"))
	json.NewEncoder(w).Encode(&res)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) { 
	res := Status_Response{Status: "OK[200]"}
	w.Header().Add("Content-Type","application/json")
	json.NewEncoder(w).Encode(&res)
}


//accept anything json 
type Data struct { 
	Data map[string]interface{}
}
type Password struct { 
	Password string `json:"password"`
}

// task from "DAY-2️⃣"
func AcceptAndGiveJSON(w http.ResponseWriter, r *http.Request) { 
var password_DATA Password
	if err := json.NewDecoder(r.Body).Decode(&password_DATA); err != nil { 
		fmt.Println("🚨Decoding Error: ",err)
		return
	}
	w.Header().Add("Content-Type","application/json")
	if err := json.NewEncoder(w).Encode(&password_DATA); err != nil { 
		fmt.Println("🚨Encoding Error: ",err)
		return
		} 
	}


	func DynamicJSON(w http.ResponseWriter, r *http.Request) { 
		var data Data
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil { 
				fmt.Println("🚨Decoding Error: ",err)
				return
			}
		w.Header().Add("Content-Type","application/json")
		json.NewEncoder(w).Encode(data)
	}