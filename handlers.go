package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetIssues(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryParam := r.Form.Get("query")
	fmt.Println(queryParam)
	json.NewEncoder(w).Encode(issues)
}

func GetIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, issue := range issues {
		if issue.PhoneNo == params["phoneNo"] {
			json.NewEncoder(w).Encode(issue)
		} else {
			json.NewEncoder(w).Encode("No. not found")
		}
		return
	}
}

func CreateIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var issue Issue
	_ = json.NewDecoder(r.Body).Decode(&issue)
	issue.PhoneNo = params["phoneNo"]
	issues = append(issues, issue)
	json.NewEncoder(w).Encode(issue)
}

func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, issue := range issues {
		if issue.PhoneNo == params["phoneNo"] {
			issues = append(issues[:index], issues[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(issues)
	}
}
