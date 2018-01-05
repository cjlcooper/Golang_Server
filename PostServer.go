package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"time"
)

//type struct
type ResponseDict struct {
	Name    string
	ID      string
	Code    string
}

func main() {
	// print start string
	fmt.Println(returnValueTest())
	//
	http.HandleFunc("/test",test)
	http.HandleFunc("/getJson",getJson)
	http.HandleFunc("/concurrent",concurrent)
	//start Server
	http.ListenAndServe(":8080",nil)
}

// test
func test(w http.ResponseWriter,r *http.Request) {
	defer r.Body.Close()
	var dict map[string]interface{}
	con, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(con,&dict)
	fmt.Println("Request->")
	fmt.Println(dict)
	fmt.Fprintln(w, dict)
}

// get json
func getJson(w http.ResponseWriter,r *http.Request) {
	defer r.Body.Close()
	var dict map[string]interface{}
	con, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(con,&dict)
	//
	var responseDict = ResponseDict{}
	responseDict.Name = "Hello Client"
	responseDict.ID = "001"
	responseDict.Code = "0"
	var data,_ = json.Marshal(responseDict)
	fmt.Fprintln(w, string(data))
}

// concurrent
func concurrent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	go saveInDB()
}

func saveInDB() {
	fmt.Println("Hello,goroutine")
}

// return value
func returnValueTest() (string) {
	fmt.Println(time.Now())
	return "Start Server ..."
}
