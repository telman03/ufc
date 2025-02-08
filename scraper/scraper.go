package scraper

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"

	// "github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/telman03/ufc/db"
	"github.com/telman03/ufc/models"
	// "gorm.io/gorm"
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



func ScrapeAndStoreRankings() {
	url := "https://www.tapology.com/rankings/current-top-ten-heavyweight-mma-fighters-265-pounds"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch page: ", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Error: Status code %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Failed to parse page", err)
	}

	// Fetch all fighters from the database
	var fighters []models.Fighter
	db.DB.Find(&fighters)

	doc.Find(".rankingItemsItem").Each(func(i int, s *goquery.Selection) {
		if i >= 15 { // Limit to the first 15 rankings
			return // Exit the loop after 15 iterations
		}
		rank := i + 1
		name := strings.TrimSpace(s.Find(".rankingItemsItemRow .name").Text()) // Corrected selector

		// Debug: Print the scraped name
		fmt.Printf("Scraped Name: %s\n", name)

		// Remove the nickname from the scraped name
		cleanedName := RemoveNickname(name)
		fmt.Printf("Cleaned Name: %s\n", cleanedName)

		// Find the fighter in the database using the cleaned name
		var fighter models.Fighter
		result := db.DB.Where("name = ?", cleanedName).First(&fighter)
		if result.Error != nil {
			log.Printf("Fighter %s not found in database. Skipping...\n", cleanedName)
			return
		}

		// Create a new ranking entry
		ranking := models.Ranking{
			FighterID: fighter.ID,
			Rank:      rank,
			Division:  "Heavyweight",
		}

		// Save the ranking to the database
		db.DB.Create(&ranking)
	})
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


// RemoveNickname removes the nickname (text within quotes) from the scraped name
func RemoveNickname(name string) string {
	// Regex to remove text within quotes
	re := regexp.MustCompile(`\s*".*?"\s*`)
	return re.ReplaceAllString(name, " ")
}