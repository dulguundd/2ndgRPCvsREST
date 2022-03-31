package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	Name        string `json:"name"`
	Age         int32  `json:"age"`
	ID          int32  `json:"ID"`
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
	router := mux.NewRouter()

	//wiring

	//define routes
	router.HandleFunc("/user", NewUser).Methods(http.MethodPost)

	log.Printf("listening 8080 port")

	// starting server
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	log.Println("GET request")
	var newuser newUser
	start := time.Now()

	err := json.NewDecoder(r.Body).Decode(&newuser)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)
	} else {
		user, err := createNewUser(newuser)
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, err)
		} else {
			writeResponse(w, http.StatusCreated, user)
			serviceLatencyLogger(start)
		}

	}
}

func createNewUser(req newUser) (*User, error) {
	log.Printf("Recieved: %v", req.Name)
	var user_id int32 = int32(rand.Intn(1000))
	return &User{
		Name:        req.Name,
		Age:         req.Age,
		ID:          user_id,
		INACTENDD:   "01/01/2038",
		Error:       "",
		RETAILER:    "2",
		CLASS:       "PRE_Hybrid_14900_N",
		ACTENDD:     "01/01/2038",
		ADMINST:     "1",
		CREDITVIOCE: "5490500",
		CODE:        "0",
		PHONE:       "94300048",
		RBAL:        "5490500",
	}, nil
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func serviceLatencyLogger(start time.Time) {
	elapsed := time.Since(start)
	logMessage := fmt.Sprintf("response latencie %s", elapsed)
	log.Println(logMessage)
}
