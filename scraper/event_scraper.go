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
		date := strings.TrimSpace(e.ChildText("span.b-statistics__date"))
		location := strings.TrimSpace(e.ChildText("td.b-statistics__table-col:nth-child(2)"))

		if eventName == "" || eventURL == "" {
			return
		}

		event := models.Event{
			Name:     eventName,
			Date:     date,
			Location: location,
			URL:      eventURL,
		}
		
		if err := db.DB.Create(&event).Error; err != nil {
			log.Println("Failed to save event:", err)
		} else {
			fmt.Println("Event saved:", event)
		}
	})

	c.Visit(baseURL)
}

// ScrapeFightCards scrapes fight card data for all saved events
func ScrapeFightCards() {
	var events []models.Event
	if err := db.DB.Find(&events).Error; err != nil {
		log.Println("Failed to fetch events:", err)
		return
	}

	for _, event := range events {
		scrapeEventFights(event.URL, event.ID)
	}
}

// scrapeEventFights scrapes fight details for a specific event
func scrapeEventFights(eventURL string, eventID uint) {
	c := colly.NewCollector()

	c.OnHTML("table.b-fight-details__table tbody tr", func(e *colly.HTMLElement) {
		fighters := e.ChildTexts("td.b-fight-details__table-col a")

		if len(fighters) < 2 {
			log.Println("Skipping fight due to missing fighter names:", fighters)
			return
		}

		fighter1 := strings.TrimSpace(fighters[0])
		fighter2 := strings.TrimSpace(fighters[1])
		weightClass := strings.TrimSpace(e.ChildText("td.b-fight-details__table-col:nth-child(7)"))

		fmt.Printf("Fighter 1: [%s], Fighter 2: [%s], Weight Class: [%s]\n", fighter1, fighter2, weightClass)

		if fighter1 == "" || fighter2 == "" || weightClass == "" {
			return
		}

		fight := models.Fight{
			EventID:    eventID,
			Fighter1:   fighter1,
			Fighter2:   fighter2,
			WeightClass: weightClass,
		}

		if err := db.DB.Create(&fight).Error; err != nil {
			log.Println("Failed to save fight:", err)
		} else {
			fmt.Println("Fight saved:", fight)
		}
	})

	c.Visit(eventURL)
}