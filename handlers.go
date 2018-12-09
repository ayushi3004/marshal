package main

import (
	"encoding/json"
	"net/http"
	"github.com/blevesearch/bleve"
	"github.com/gorilla/mux"
)

func GetIssues(w http.ResponseWriter, r *http.Request) {
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
			MssgShort:       res.Fields["mssgShort"].(string),
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
