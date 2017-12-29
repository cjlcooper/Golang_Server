package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/myHandle",myHandle)
	http.ListenAndServe(":8080",nil)
}

func myHandle(w http.ResponseWriter,r *http.Request) {
	defer r.Body.Close()
	var dict map[string]interface{}
	con, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(con,&dict)
	fmt.Println("Request->")
	fmt.Println(dict)
	fmt.Fprintln(w, dict)
	//w.Write("ok")
}