package main

import (
	"log"
	"net/http"

	"github.com/blevesearch/bleve"
	"github.com/gorilla/mux"
)

var issues []Issue
var index bleve.Index

func main() {
	// issues = append(issues, Issue{MssgShort: "Need help. 4 people...", Status: "Open", Area: "Velachery", Address: "blah blah", PhoneNo: "9765829276", DisasterType: "Flood", SquashedMssgs: []string{"Need help. 4 people stuck in flood. No food or water. Please send boats", "Stuck in flood. 4 people. Send help urgently"}, Acknowledgement: true})
	// issues = append(issues, Issue{MssgShort: "Fire! Fire! family...", Status: "Dispatched", Area: "Perungudi", Address: "blu blah", PhoneNo: "222222222", DisasterType: "Fire", SquashedMssgs: []string{"Fire! Fire! family stuck. send fire brigade "}, Acknowledgement: true})
	// issues = append(issues, Issue{MssgShort: "Flood. House drowning...", Status: "Closed", Area: "OMR", Address: "boom ba", PhoneNo: "333333333", DisasterType: "Flood", SquashedMssgs: []string{"Flood. House drowning. We are 4 people. 2 kids. Send help. Also send food and torch. No electricity.", "We are all going to die. All of us."}, Acknowledgement: false})

	//Comment it out after indexing
	// var issue1 = Issue{Status: "Open", Area: "Velachery", Address: "blah blah", PhoneNo: "9765829276", DisasterType: "Flood", DetailedMssg: "Need help. 4 people stuck in flood. No food or water. Please send boats", Acknowledgement: true}
	// var issue2 = Issue{Status: "Dispatched", Area: "Perungudi", Address: "blu blah", PhoneNo: "222222222", DisasterType: "Fire", DetailedMssg: "Fire! Fire! family stuck. send fire brigade ", Acknowledgement: true}
	// var issue3 = Issue{Status: "Closed", Area: "OMR", Address: "boom ba", PhoneNo: "333333333", DisasterType: "Flood", DetailedMssg: "Flood. House drowning. We are 4 people. 2 kids. Send help. Also send food and torch. No electricity.", Acknowledgement: false}
	// var issue4 = Issue{Status: "Closed", Area: "OMR", Address: "boom ba", PhoneNo: "444444444", DisasterType: "Flood", DetailedMssg: "Flood. House drowning. We are 4 people. 2 kids. Send help. Also send food and torch. No electricity.", Acknowledgement: false}

	// mapping := bleve.NewIndexMapping()
	// index, err := bleve.New("test.bleve", mapping)
	// if err != nil {
	// 	panic(err)
	// }
	// index.Index(issue1.DetailedMssg, issue1)
	// index.Index(issue2.DetailedMssg, issue2)
	// index.Index(issue3.DetailedMssg, issue3)
	// index.Index(issue4.DetailedMssg, issue4)

	var err error
	index, err = bleve.Open("test.bleve")
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/issues/{query}", GetIssues).Methods("GET")
	router.HandleFunc("/issues/{phoneNo}", GetIssue).Methods("GET")
	router.HandleFunc("/issues/{phoneNo}", CreateIssue).Methods("POST")
	router.HandleFunc("/issues/{phoneNo}", DeleteIssue).Methods("DELETE")
	router.PathPrefix("/app").Handler(http.StripPrefix("/app", http.FileServer(http.Dir("./app/"))))
	log.Fatal(http.ListenAndServe(":8000", router))

}
