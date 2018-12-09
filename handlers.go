package main

import (
	"encoding/json"
	"net/http"

	"github.com/blevesearch/bleve"
	"github.com/gorilla/mux"
)

func GetIssues(w http.ResponseWriter, r *http.Request) {
	query := bleve.NewQueryStringQuery(helpParam)
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"*"}
	searchResults, err := index.Search(search)
	if err != nil {
		panic(err)
	}

	var issues []Issue
	for _, res := range searchResults.Hits {
		issue := Issue{
			Status:          res.Fields["status"].(string),
			Area:            res.Fields["area"].(string),
			Address:         res.Fields["address"].(string),
			PhoneNo:         res.Fields["phoneNo"].(string),
			DisasterType:    res.Fields["disasterType"].(string),
			DetailedMssg:    res.Fields["detailedMssg"].(string),
			Acknowledgement: res.Fields["acknowledgement"].(bool),
		}
		issues = append(issues, issue)
	}
	json.NewEncoder(w).Encode(issues)
}

func GetIssue(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	queryParam := params["query"]
	query := bleve.NewQueryStringQuery(queryParam)
	search := bleve.NewSearchRequest(query)
	search.Fields = []string{"*"}
	searchResults, err := index.Search(search)
	if err != nil {
		panic(err)
	}

	var issues []Issue
	for _, res := range searchResults.Hits {
		issue := Issue{
			Status:          res.Fields["status"].(string),
			Area:            res.Fields["area"].(string),
			Address:         res.Fields["address"].(string),
			PhoneNo:         res.Fields["phoneNo"].(string),
			DisasterType:    res.Fields["disasterType"].(string),
			DetailedMssg:    res.Fields["detailedMssg"].(string),
			Acknowledgement: res.Fields["acknowledgement"].(bool),
		}
		issues = append(issues, issue)
	}
	json.NewEncoder(w).Encode(issues)
}

// func GetIssue(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for _, issue := range issues {
// 		if issue.PhoneNo == params["phoneNo"] {
// 			json.NewEncoder(w).Encode(issue)
// 		} else {
// 			json.NewEncoder(w).Encode("No. not found")
// 		}
// 		return
// 	}
// }

func CreateIssue(w http.ResponseWriter, r *http.Request) {
	var issue Issue
	_ = json.NewDecoder(r.Body).Decode(&issue)
	index.Index(issue.DetailedMssg, issue)
	json.NewEncoder(w).Encode("Indexed")
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
