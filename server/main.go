package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	ResSuccess  int = 0
	ResUNknown  int = 1
	ResTimeout  int = 2
	ResInternal int = 3
	ResOtherErr int = 4
)
const (
	StatusOK            int = 200
	StatusBadRequest    int = 400
	StatusUnauthorized  int = 401
	StatusInternalError int = 500
)

type MyHandler struct{}
type ResPonse struct {
	ErrCode int    `json:"errCode"`
	Message string `json:"message"`
}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	response := ResPonse{ResUNknown, "wrong request"}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.WriteHeader(StatusBadRequest)
	w.Write(b)
	fmt.Printf("!!!error:unrecognized url %s!!!\n", r.URL.Path)
}

func normalCase(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength <= 0 || r.ContentLength > 1024*1024*2 {
		response := ResPonse{ResTimeout, "data length error"}
		b, _ := json.Marshal(response)
		w.WriteHeader(StatusBadRequest)
		w.Write(b)
		fmt.Printf("!!!error: data length error!!!\n")
		return
	}

	var oneMessage OneMessage
	data_buf := make([]byte, r.ContentLength)
	io.ReadFull(r.Body, data_buf)
	err := json.Unmarshal(data_buf, &oneMessage)
	if err != nil {
		response := ResPonse{ResUNknown, err.Error()}
		b, _ := json.Marshal(response)
		w.WriteHeader(StatusBadRequest)
		w.Write(b)
		fmt.Printf("!!!error: %s!!!\n", err.Error())
		return
	}

	err = oneMessage.print()
	if err != nil {
		response := ResPonse{ResUNknown, err.Error()}
		b, _ := json.Marshal(response)
		w.WriteHeader(StatusBadRequest)
		w.Write(b)
		fmt.Printf("!!!error: %s!!!\n", err.Error())
		return
	}
	response := ResPonse{ResSuccess, ""}
	b, _ := json.Marshal(response)
	w.WriteHeader(StatusOK)
	w.Write(b)
	return
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", MyHandler{})
	mux.HandleFunc("/v1/pos/operate/backflow", normalCase)

	server := &http.Server{
		Addr:         ":8023",
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	server.SetKeepAlivesEnabled(false)
	fmt.Println("httpserver started,listen 8023 port", "accept http url /v1/pos/operate/backflow")
	server.ListenAndServe()
}
