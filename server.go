package main

import (
	"encoding/json"
	"encoding/hex"
	"net/http"
	"bytes"
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


func WriteError(w http.ResponseWriter, err error) {
	resp := HTTPResponse{"", err.Error()}
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

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in verif")

	args := conv(r)
	addr := args["addr"]
	hash := args["hash"]
	sig := args["sig"]

	fmt.Println(addr, hash, sig)
	
	client := &http.Client{}
	jsonBytes, err := json.Marshal(args)
	if err != nil{
		WriteError(w, err)
		return 
	}
	buf := bytes.NewBuffer(jsonBytes)
	req, err := http.NewRequest("POST", "http://localhost:4767/verify", buf)
	if err != nil{
		WriteError(w, err)
		return 
	}
	resp, err := client.Do(req)
	if err != nil{
		WriteError(w, err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		WriteError(w, err)
		return
	}
	resp.Body.Close()
	fmt.Println("Body:", string(b))
	fmt.Println(resp.Status)
	WriteResult(w, string(b))

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
	var bytes = make([]byte, n)
	rand.Read(bytes)
	s := hex.EncodeToString(bytes)
	return s
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/nonce", nonceHandler);
	mux.HandleFunc("/verify", verifyHandler);
	fmt.Println(http.ListenAndServe(":8080", mux))
}
