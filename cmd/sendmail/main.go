package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"

	"github.com/jacobalberty/smtpoauth2"
)

func main() {
	from := flag.String("from", "", "Your email address")
	totmp := flag.String("to", "", "Destination email address")
	body := flag.String("body", "", "message email")
	smtpHost := flag.String("server", "smtp.gmail.com", "SMTP server address")
	smtpPort := flag.String("port", "587", "SMTP server port")
	flag.Parse()
	if *from == "" || *totmp == "" || *body == "" {
		fmt.Printf("%s, %s, %s", from, totmp, body)
		log.Fatalf("Must provide all arguments see -help")
	}
	to := []string{*totmp}

	tokenJSON, err := ioutil.ReadFile("token.json")
	if err != nil {
		log.Fatalf("Error opening 'token.json': %v", err)
	}
	token := make(map[string]string)
	err = json.Unmarshal(tokenJSON, &token)
	if err != nil {
		log.Fatalf("Error unmarshaling token: %v", err)
	}

	auth := smtpoauth2.Oauth2(*from, token["token_type"], token["access_token"])
	err = smtp.SendMail(*smtpHost+":"+*smtpPort, auth, *from, to, []byte(*body))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
