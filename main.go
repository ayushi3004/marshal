package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var issues []Issue

func main() {
	issues = append(issues, Issue{MssgShort: "Need help. 4 people...", Status: "Open", Area: "Velachery", Address: "blah blah", PhoneNo: "9765829276", DisasterType: "Flood", SquashedMssgs: []string{"Need help. 4 people stuck in flood. No food or water. Please send boats", "Stuck in flood. 4 people. Send help urgently"}, Acknowledgement: true})
	issues = append(issues, Issue{MssgShort: "Fire! Fire! family...", Status: "Dispatched", Area: "Perungudi", Address: "blu blah", PhoneNo: "222222222", DisasterType: "Fire", SquashedMssgs: []string{"Fire! Fire! family stuck. send fire brigade "}, Acknowledgement: true})
	issues = append(issues, Issue{MssgShort: "Flood. House drowning...", Status: "Closed", Area: "OMR", Address: "boom ba", PhoneNo: "333333333", DisasterType: "Flood", SquashedMssgs: []string{"Flood. House drowning. We are 4 people. 2 kids. Send help. Also send food and torch. No electricity.", "We are all going to die. All of us."}, Acknowledgement: false})

	router := mux.NewRouter()
	router.HandleFunc("/issues", GetIssues).Methods("GET")
	router.HandleFunc("/issues/{phoneNo}", GetIssue).Methods("GET")
	router.HandleFunc("/issues/{phoneNo}", CreateIssue).Methods("POST")
	router.HandleFunc("/issues/{phoneNo}", DeleteIssue).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}
