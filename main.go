package main

import (
    "log"
    "net/smtp"
)

func main() {
    to_email     := "to mail address"
    from_email   := "from mail address"
    subject_body := "subject" + "\n\n" + "body"
    status       := smtp.SendMail("server:port", nil, from_email, []string{to_email}, []byte(subject_body))
    if status != nil {
        log.Printf("Error from SMTP Server: %s", status)
    }
    log.Print("Email Sent Successfully")
}
// 参考
// https://netcorecloud.com/tutorials/send-email-through-gmail-smtp-server-using-go/

// どうやらAUTHをサポートしていないらしい…
// 2021/04/05 13:31:11 Error from SMTP Server: smtp: server doesn't support AUTH