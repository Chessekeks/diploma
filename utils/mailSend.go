package utils

import (
	"MurihKursovoi/model"
	"bytes"
	"net/textproto"
	"github.com/jordan-wright/email"
	"html/template"
	"log"
	"net/smtp"
)

func parseTemplate(filename string, data interface{}) []byte {
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(model.ErrorLog("handler", "parseTemplate:parseFile", err))
		return nil
	}
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, data)
	if err != nil {
		log.Println(model.ErrorLog("handler", "parseTemplate:execute", err))
		return nil
	}
	return buffer.Bytes()
}

func sendMail(to []string, templateName string, items interface{}) bool {
	//config.Read()
	/*body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	SMTP := fmt.Sprintf("%s:%d", "smtp.gmail.com", 587)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", "chessekeks@gmail.com", "Satpaevcityprod1", "smtp.gmail.com"), "chessekeks@gmail.com", r.to, []byte(body)); err != nil {
		log.Println(model.ErrorLog("handler", "sendMail", err))
		return false
	}
	return true*/
	template := parseTemplate(templateName, items)
	if template != nil {
		e := email.Email{
			To:      to,
			From:    "Nurzhan Ilyassov <chessekeks@gmail.com>",
			Subject: "Confirmation Code",
			HTML:    template,
			Headers: textproto.MIMEHeader{},
		}
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "chessekeks@gmail.com", "Satpaevcityprod1", "smtp.gmail.com"))
		return true
	}
	return false
}

func Send(to []string, templateName string, items interface{}) {
	if ok := sendMail(to, templateName, items); ok {
		log.Println("Email has been sent")
	} else {
		log.Println("Failed to send the email")
	}
}

/*func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Println(model.ErrorLog("handler", "Read", err))
	}
}*/
