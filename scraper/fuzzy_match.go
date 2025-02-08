// fuzzy_match.go
package scraper

import (
	"fmt"
	"strings"

	"github.com/paul-mannino/go-fuzzywuzzy"
	"github.com/telman03/ufc/models"
)

// FuzzyMatch finds the best match for a given name in the database
func FuzzyMatch(name string, fighters []models.Fighter) (*models.Fighter, error) {
	var bestMatch *models.Fighter
	highestScore := 0

	for _, fighter := range fighters {
		// Compare the scraped name with the fighter's name in the database
		score := fuzzy.TokenSortRatio(strings.ToLower(name), strings.ToLower(fighter.Name))

		// If the score is higher than the previous best, update the best match
		if score > highestScore {
			highestScore = score
			bestMatch = &fighter
		}
	}

	// Set a threshold for the match score (e.g., 80)
	if highestScore >= 80 {
		return bestMatch, nil
	}

	return nil, fmt.Errorf("no close match found for %s", name)
}