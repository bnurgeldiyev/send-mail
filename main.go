package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"mail/config"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gorilla/mux"
)

func sendMailMessage(to []string, body string) {

	auth := smtp.PlainAuth(
		"",
		config.Conf.Mail,
		config.Conf.Password,
		config.Conf.Host,
	)

	err := smtp.SendMail(
		config.Conf.Host+":"+config.Conf.Port,
		auth,
		config.Conf.Mail,
		to,
		[]byte(body),
	)

	if err != nil {
		log.Fatal("error in smtp.SendMail")
		return
	}
}

func HandleMessage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		mails := r.FormValue("mails") // []string
		msg := r.FormValue("message") // string
		title := r.FormValue("title") // string

		bytes := []byte(mails)
		var mailList []string
		err := json.Unmarshal(bytes, &mailList)
		if err != nil {
			eMsg := "An error occurred on json.Unmarshal(coords)"
			log.Fatal(eMsg)
			return
		}

		header := make(map[string]string)
		header["From"] = config.Conf.Mail
		header["Subject"] = title
		header["MIME-Version"] = "1.0"
		header["Content-Type"] = "text/plain; charset=\"utf-8\""
		header["Content-Transfer-Encoding"] = "base64"

		message := ""
		for k, v := range header {
			message += fmt.Sprintf("%s: %s\r\n", k, v)
		}

		message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(msg))

		go sendMailMessage(mailList, message)

		fmt.Println("Successfully sended mail for", mailList)

	} else {
		log.Fatal("only post requests")
	}
}

func main() {

	err := config.ReadConfig("config.json")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/send/message", HandleMessage)

	srv := &http.Server{
		Handler:      r,
		Addr:         config.Conf.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("<--START-SERVER-->")
	log.Fatal(srv.ListenAndServe())

}
