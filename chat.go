package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//SendMessage sends a messsage on whatsapp with twillio
func SendMessage(message string, number string) {

	TWILIO_ACCOUNT := os.Getenv("TWILIO_ACCOUNT")
	TWILIO_AUTH := os.Getenv("TWILIO_AUTH")

	whatsappNumber := fmt.Sprintf("whatsapp:+91%s", number)

	data := url.Values{}
	data.Set("From", "whatsapp:+14155238886")
	data.Add("To", whatsappNumber)
	data.Add("Body", message)

	reqTemplate := "https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json"

	reqURL := fmt.Sprintf(reqTemplate, TWILIO_ACCOUNT)

	req, err := http.NewRequest("POST", reqURL, strings.NewReader(data.Encode()))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(TWILIO_ACCOUNT, TWILIO_AUTH)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
