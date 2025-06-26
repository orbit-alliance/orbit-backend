package main

import (
	"log"

	"github.com/orbit-alliance/orbit-backend/internal/domain/shared"
	"github.com/robfig/cron/v3"
)

func dailyJobs(c *cron.Cron, jobs []shared.Job) {
	for _, job := range jobs {
		// “1 0 * * *”  →  minuto 1, hora 0, todos os dias
		if _, err := c.AddFunc("1 0 * * *", job); err != nil {
			log.Printf("Failed to schedule job %T: %v", job, err)
		} else {
			log.Printf("Scheduled job %T", job)
		}
	}
}
