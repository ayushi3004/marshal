package main

import (
    "fmt"
    "log"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/blevesearch/bleve"
    "github.com/blevesearch/bleve/document"
)

var issues []Issue

func main() {
    // issues = append(issues, Issue{MssgShort: "Need help. 4 people...", Status: "Open", Area: "Velachery", Address: "blah blah", PhoneNo: "9765829276", DisasterType: "Flood", SquashedMssgs: []string{"Need help. 4 people stuck in flood. No food or water. Please send boats", "Stuck in flood. 4 people. Send help urgently"}, Acknowledgement: true})
    // issues = append(issues, Issue{MssgShort: "Fire! Fire! family...", Status: "Dispatched", Area: "Perungudi", Address: "blu blah", PhoneNo: "222222222", DisasterType: "Fire", SquashedMssgs: []string{"Fire! Fire! family stuck. send fire brigade "}, Acknowledgement: true})
    // issues = append(issues, Issue{MssgShort: "Flood. House drowning...", Status: "Closed", Area: "OMR", Address: "boom ba", PhoneNo: "333333333", DisasterType: "Flood", SquashedMssgs: []string{"Flood. House drowning. We are 4 people. 2 kids. Send help. Also send food and torch. No electricity.", "We are all going to die. All of us."}, Acknowledgement: false})

    // var issue = Issue{MssgShort: "Need help. 4 people...", Status: "Open", Area: "Velachery", Address: "blah blah", PhoneNo: "9765829276", DisasterType: "Flood", SquashedMssgs: []string{"Need help. 4 people stuck in flood. No food or water. Please send boats", "Stuck in flood. 4 people. Send help urgently"}, Acknowledgement: true}
    // mapping := bleve.NewIndexMapping()
	// index, err := bleve.New("test.bleve", mapping)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// index some data
	// index.Index(issue.SquashedMssgs[0], issue)
    
    index, _ := bleve.Open("test.bleve")
    query := bleve.NewQueryStringQuery("Stuck in flood. 4 people. Send help urgently")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(searchResults)

    // fmt.Println(searchResults.Hits[0].Fields["Address"].(string))
    BleveSearchDocs := getBleveDocsFromSearchResults(searchResults, index)
    fmt.Println(BleveSearchDocs)

    router := mux.NewRouter()
	router.HandleFunc("/issues", GetIssues).Methods("GET")
	router.HandleFunc("/issues/{phoneNo}", GetIssue).Methods("GET")
	router.HandleFunc("/issues/{phoneNo}", CreateIssue).Methods("POST")
	router.HandleFunc("/issues/{phoneNo}", DeleteIssue).Methods("DELETE")
	router.PathPrefix("/app").Handler(http.StripPrefix("/app", http.FileServer(http.Dir("./app/"))))
	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBleveDocsFromSearchResults(
	results *bleve.SearchResult,
	index bleve.Index,
) [][]byte {
	docs := make([][]byte, 0)

	for _, val := range results.Hits {
		id := val.ID
		doc, _ := index.Document(id)

		rv := struct {
			ID     string                 `json:"id"`
			Fields map[string]interface{} `json:"fields"`
		}{
			ID:     id,
			Fields: map[string]interface{}{},
		}
		for _, field := range doc.Fields {
			var newval interface{}
			switch field := field.(type) {
			case *document.TextField:
				newval = string(field.Value())
			case *document.NumericField:
				n, err := field.Number()
				if err == nil {
					newval = n
				}
			existing, existed := rv.Fields[field.Name()]
			if existed {
				switch existing := existing.(type) {
				case []interface{}:
					rv.Fields[field.Name()] = append(existing, newval)
				case interface{}:
					arr := make([]interface{}, 2)
					arr[0] = existing
					arr[1] = newval
					rv.Fields[field.Name()] = arr
				}
			} else {
				rv.Fields[field.Name()] = newval
			}
		}
		j2, _ := json.MarshalIndent(rv, "", "    ")
		docs = append(docs, j2)
	}
    }
return docs
}
