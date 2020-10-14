package mailmanager

import (
    "log"
    "net/smtp"
)

const (
    from     = "example@domain.com"
    password = "password"
    smtpHost = "smtp.host.com"
    smtpPort = "portNumber"
)

func SendMessage(body string, emailTo string) {

    msg := []byte("From: " + from + "\n" +
        "To: " + emailTo + "\n" +
        "Subject: avitoPriceTracking\n\n" +
        body)

    auth := smtp.PlainAuth("", from, password, smtpHost)

    err := smtp.SendMail(
        smtpHost+":"+smtpPort,
        auth,
        from,
        []string{emailTo},
        msg,
    )
    if err != nil {
        log.Fatal(err)
    }
}
