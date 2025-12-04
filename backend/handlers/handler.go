package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DynamicJSON(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println("🚨Decoding Error: ", err)
		return
	}
	fmt.Println("recieved: ", data)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
