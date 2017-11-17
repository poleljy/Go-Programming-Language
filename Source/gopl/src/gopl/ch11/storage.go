package ch11

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(userName string) int64 { return 0 }

const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"
const template = `Warning: you are using %d bytes of storage, %d%% of your quota.`

var notifyUser = func(userName, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)

	err := smtp.SendMail(hostname+":587", auth, sender, []string{userName}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", userName, err)
	}
}

func CheckQuota(userName string) {
	used := bytesInUse(userName)

	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(userName, msg)
}
