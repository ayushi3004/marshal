package main

type Issue struct {
	MssgShort       string   `json:"mssgShort"`
	Status          string   `json:"status"`
	Area            string   `json:"area"`
	Address         string   `json:"address"`
	PhoneNo         string   `json:"phoneNo"`
	DisasterType    string   `json:"disasterType"`
	SquashedMssgs   []string `json:"squashedMssgs"`
	Acknowledgement bool     `json:"acknowledgement"`
}
