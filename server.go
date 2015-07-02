package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"crypto/rand"
	"io/ioutil"
)

//var listenAddr = "0.0.0.0:8080"


func rootHandler(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Path[1:]

	if _, err := os.Stat(f); err != nil {
		f = "index.html"
	}
	http.ServeFile(w, r, f)
}

type HTTPResponse struct {
	Response string
	Error string
}

func WriteResult(w http.ResponseWriter, result string) {
	resp := HTTPResponse{result, ""}
	b, _ := json.Marshal(resp)
	w.Write(b)
}

func nonceHandler(w http.ResponseWriter, r *http.Request) {
	args := conv(r)
	val := args["nonce"]
	
	if val == "TRUE" {
		res := randString(10)
		WriteResult(w, fmt.Sprintf("%v", res))
		return
	}
}

func conv(r *http.Request) (args map[string]string) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	
	if err = json.Unmarshal(body, &args); err != nil {
		return
	}
	return
}

func randString(n int) string {
	const alphanum = "0123456789abcdefABCDEF"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b % byte(len(alphanum))]
	}
	return string(bytes)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/nonce", nonceHandler);
//	fmt.Println(randString(10))
	fmt.Println(http.ListenAndServe(":8080", mux))
}
