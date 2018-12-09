package main

type Issue struct {
	Status          string   `json:"status"`
	Area            string   `json:"area"`
	Address         string   `json:"address"`
	PhoneNo         string   `json:"phoneNo"`
	DisasterType    string   `json:"disasterType"`
	DetailedMssg   	string 	 `json:"detailedMssg"`
	Acknowledgement bool     `json:"acknowledgement"`
}
