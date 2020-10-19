package mailmanager

import (
    "log"
    "net/smtp"
    "os"
)

var (
    from     = os.Getenv("FROMEMAIL")
    password = os.Getenv("PASSWORD")
    smtpHost = os.Getenv("SMTPHOST")
    smtpPort = os.Getenv("SMTPPORT")
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
