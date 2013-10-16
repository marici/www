package main

import (
	"os"
	"log"
	"net/http"
	"net/http/cgi"
	"text/template"
	"encoding/json"
	"sendmail"
)

func config(jsonpath string, obj interface{}) {
	f, _ := os.Open(jsonpath)
	decoder := json.NewDecoder(f)
	decoder.Decode(obj)
}

func render(w http.ResponseWriter, tmplpath string, obj interface{}) {
	tmpl, err := template.ParseFiles(tmplpath)
	if err != nil {
		log.Fatal("Template could not be opened.: ", err)
	}
	err = tmpl.Execute(w, obj)
	if err != nil {
		log.Fatal("Template could not be executed.: ", err)
	}
}

func parse(req *http.Request) (data map[string]string) {
	params := []string{"corp", "department", "position", "industry",
		"name", "name_kana", "email", "url", "postal1", "postal2",
		"address", "tel", "fax", "content"}
	d := make(map[string]string)
	for _, v := range params {
		d[v] = req.FormValue(v)
	}
	return d
}

func SendEmail(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data := parse(req)
		if req.FormValue("stage") == "confirm" {
			render(w, "check.html", data)
		} else {
			smtpServer := sendmail.SMTPServer{}
			email := sendmail.Email{}
			config("smtp.json", &smtpServer)
			config("email.json", &email)
			err := smtpServer.SendMailWithAuth(
				email.From,
				email.To,
				email.Subject,
				"email.txt",
				data)
			if err != nil {
				log.Fatal("Email error: ", err)
			}
			http.Redirect(w, req, "./send.html", http.StatusFound)
		}
	} else {
		http.Redirect(w, req, "./contactindex.html", http.StatusFound)
	}
}

func main() {
	http.HandleFunc("/contact/contact.cgi", SendEmail)
	cgi.Serve(nil)
}
