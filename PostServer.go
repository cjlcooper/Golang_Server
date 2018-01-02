package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

//type struct
type ResponseDict struct {
	Name    string
	ID      string
	Code    string
}

func main() {
	http.HandleFunc("/test",test)
	http.HandleFunc("/getJson",getJson)
	http.ListenAndServe(":8080",nil)
}

func test(w http.ResponseWriter,r *http.Request) {
	defer r.Body.Close()
	var dict map[string]interface{}
	con, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(con,&dict)
	fmt.Println("Request->")
	fmt.Println(dict)
	fmt.Fprintln(w, dict)
}

func getJson(w http.ResponseWriter,r *http.Request) {
	defer r.Body.Close()
	var dict map[string]interface{}
	con, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(con,&dict)
	var responseDict = ResponseDict{}
	responseDict.Name = "Hello Client"
	responseDict.ID = "001"
	responseDict.Code = "0"
	var data,_ = json.Marshal(responseDict)
	fmt.Fprintln(w, string(data))
}