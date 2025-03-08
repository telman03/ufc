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

		// Check if event already exists in the database
		var existingEvent models.Event
		if err := db.DB.Where("url = ?", eventURL).First(&existingEvent).Error; err == nil {
			fmt.Println("Event already exists, skipping:", eventName)
			return
		}

		// Save only if it doesn't exist
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

		// Check if fight already exists
		var existingFight models.Fight
		if err := db.DB.Where("event_id = ? AND fighter1 = ? AND fighter2 = ?", eventID, fighter1, fighter2).First(&existingFight).Error; err == nil {
			fmt.Println("Fight already exists, skipping:", fighter1, "vs", fighter2)
			return
		}

		// Save only if fight doesn't exist
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
