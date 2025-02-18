package scheduler

import (
	"log"
	"time"

	"github.com/telman03/ufc/scraper"
)

func StartScheduler() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Running scheduled scraping...")
		scraper.ScrapeUpcomingEvents()
		log.Println("Scheduled scraping completed.")
	}
}
