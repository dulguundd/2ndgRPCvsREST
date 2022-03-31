package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	host = "http://172.22.2.215"
	port = ":8080"
)

type User struct {
	Name        string
	Age         int32
	ID          int32
	INACTENDD   string `json:"INACTENDD"`
	Error       string `json:"Error"`
	RETAILER    string `json:"RETAILER"`
	CLASS       string `json:"CLASS"`
	ACTENDD     string `json:"ACTENDD"`
	ADMINST     string `json:"ADMINST"`
	CREDITVIOCE string `json:"CREDITVIOCE"`
	CODE        string `json:"CODE"`
	PHONE       string `json:"PHONE"`
	RBAL        string `json:"RBAL"`
}

type newUser struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func main() {
	start := time.Now()
	var user User

	count := 1

	var new_users = make(map[string]int32)
	new_users["RESTdomain"] = 43
	new_users["Testing"] = 30
	URL := fmt.Sprintf(host + port + "/user")

	runningcount := 0
	for runningcount < count {
		for name, age := range new_users {
			data, _ := json.Marshal(&newUser{Name: name,
				Age: age})
			r, err := http.Post(URL, "applecation/json", bytes.NewBuffer(data))
			if err != nil {
				log.Fatal(err)
			}
			err = json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf(`User Detail:
NAME: %s
AGE: %d
ID: %d`, user.Name, user.Age, user.ID)
		}
		runningcount++
	}
	serviceLatencyLogger(start)
}

func serviceLatencyLogger(start time.Time) {
	elapsed := time.Since(start)
	logMessage := fmt.Sprintf("response latencie %s", elapsed)
	log.Println(logMessage)
}
