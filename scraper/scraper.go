package scraper

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/telman03/ufc/db"
	"github.com/telman03/ufc/models"
)

func ScrapeAndStoreFighters() {
	baseURL := "http://ufcstats.com/statistics/fighters?char=z&page=all"
	c := colly.NewCollector()


	c.OnHTML("table.b-statistics__table tbody tr", func(e *colly.HTMLElement) {
		firstName := e.ChildText("td:nth-child(1)")
		lastName := e.ChildText("td:nth-child(2)")
		nickname := e.ChildText("td:nth-child(3)")
		height := e.ChildText("td:nth-child(4)")
		weight := e.ChildText("td:nth-child(5)")
		reach := e.ChildText("td:nth-child(6)")
		stance := e.ChildText("td:nth-child(7)")
		wins := parseInt(e.ChildText("td:nth-child(8)"))
		losses := parseInt(e.ChildText("td:nth-child(9)"))
		draws := parseInt(e.ChildText("td:nth-child(10)"))


		fighter := models.Fighter{
			Name:      fmt.Sprintf("%s %s", firstName, lastName),
			FirstName: firstName,
			LastName:  lastName,
			Nickname:  nickname,
			Height:    height,
			Weight:    weight,
			Reach:     reach,
			Stance:    stance,
			Wins:      wins,
			Losses:    losses,
			Draws:     draws,
		}

		err := db.DB.Create(&fighter).Error
		if err != nil {
			log.Printf("Error saving fighter %s: %v", fighter.Name, err)
		} else {
			fmt.Printf("Saved fighter: %s\n", fighter.Name)
		}
	})

	err := c.Visit(baseURL)
	if err != nil {
		log.Fatal(err)
	}
}

func parseInt(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}
