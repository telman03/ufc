package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/telman03/ufc/db"
	"github.com/telman03/ufc/models"
)


func ScrapeUpcomingEvents() {
	baseURL := "http://ufcstats.com/statistics/events/upcoming"
	c := colly.NewCollector()

	c.OnHTML("table.b-statistics__table-events tbody tr", func(e *colly.HTMLElement) {
		eventName := strings.TrimSpace(e.ChildText("td.b-statistics__table-col a"))
		eventURL := e.ChildAttr("td.b-statistics__table-col a", "href")
		date := strings.TrimSpace(e.ChildText("span.b-statistics__date")) // Extract date correctly
		location := strings.TrimSpace(e.ChildText("td.b-statistics__table-col:nth-child(2)")) // Extract location correctly


		if eventName == "" || eventURL == "" {
			return
		}

		event := models.Event{
			Name:     eventName,
			Date:     date,
			Location: location,
			URL:      eventURL,
		}

		// Save to DB
		if err := db.DB.Create(&event).Error; err != nil {
			log.Println("Failed to save event:", err)
		} else {
			fmt.Println("Event saved:", event)
		}
	})

	c.Visit(baseURL)
}
