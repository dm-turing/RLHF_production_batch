package main

import (
	"fmt"
	"time"
)

// BackupStatus represents the status of a backup job
type BackupStatus string

const (
	Successful BackupStatus = "successful"
	Failed     BackupStatus = "failed"
)

// BackupJob contains information about a backup job
type BackupJob struct {
	ID           string
	StartTime    time.Time
	EndTime      time.Time
	Status       BackupStatus
	ErrorMessage string
}

// Notifier interface for sending notifications
type Notifier interface {
	SendNotification(job BackupJob)
}

// EmailNotifier implements the Notifier interface to send emails
type EmailNotifier struct {
	To string
}

func (n *EmailNotifier) SendNotification(job BackupJob) {
	message := fmt.Sprintf("Backup Job %s %s at %s. Duration: %s.",
		job.Status,
		job.ID,
		job.StartTime.Format(time.RFC3339),
		job.EndTime.Sub(job.StartTime),
	)
	if job.Status == Failed {
		message += fmt.Sprintf(" Error: %s", job.ErrorMessage)
	}
	fmt.Println("Sending email:", message)
	// Actual email sending logic would go here
}

// MonitorBackups periodically checks backup statuses and sends notifications
func MonitorBackups(notifier Notifier, interval time.Duration) {
	for {
		// Simulate checking backup status
		jobs := []BackupJob{
			{
				ID:           "job1",
				StartTime:    time.Now().Add(-1 * time.Hour),
				EndTime:      time.Now(),
				Status:       Failed,
				ErrorMessage: "Disk space insufficient",
			},
			{
				ID:        "job2",
				StartTime: time.Now().Add(-2 * time.Hour),
				EndTime:   time.Now().Add(-1 * time.Hour),
				Status:    Successful,
			},
		}

		for _, job := range jobs {
			notifier.SendNotification(job)
		}

		time.Sleep(interval)
	}
}

func main() {
	notifier := &EmailNotifier{To: "admin@example.com"}
	MonitorBackups(notifier, 10*time.Second)
}
