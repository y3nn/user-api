package middleware

import (
	"fmt"
	"net/http"
)

// example middlware function
func HelloMiddleWare(fn http.HandlerFunc) http.HandlerFunc { 
	var counter int 
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("✅ Отправлен запрос handler -> /hello | 📡[total request: %d]\n",counter)
		counter++
		fn.ServeHTTP(w,r)
	})
}