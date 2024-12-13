package main

import (
	"fmt"
	"log"
	"net/smtp"
)

// BackupResult represents the outcome of a backup operation
type BackupResult struct {
	Successful bool
	Error      error
}

func main() {
	// Simulate backup
	result := performBackup()

	// Notify user of the result
	notifyBackupResult(result)
}

// performBackup simulates a backup operation
func performBackup() BackupResult {
	// Simulated operation; for demonstration, we'll randomly promote success or failure
	// In a real application, this would include actual backup logic.
	success := true // This could depend on actual backup code.
	if !success {
		return BackupResult{Successful: false, Error: fmt.Errorf("backup failed due to connectivity issue")}
	}
	return BackupResult{Successful: true}
}

// notifyBackupResult sends a notification based on the backup result
func notifyBackupResult(result BackupResult) {
	if result.Successful {
		fmt.Println("Backup completed successfully.")
		sendEmail("Backup completed successfully.")
	} else {
		fmt.Println("Backup failed:", result.Error)
		sendEmail(fmt.Sprintf("Backup failed: %s", result.Error.Error()))
	}
}

// sendEmail simulates sending an email notification
func sendEmail(message string) {
	from := "your-email@example.com"
	to := "recipient@example.com"
	subject := "Backup Notification"
	body := "Subject: " + subject + "\n\n" + message

	// Set up authentication information.
	auth := smtp.PlainAuth("", from, "your-email-password", "smtp.example.com")

	// Connect to the remote SMTP server.
	err := smtp.SendMail("smtp.example.com:587", auth, from, []string{to}, []byte(body))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Notification sent!")
}
