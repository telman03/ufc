package models

// FighterStats swagger model definition
// @Description Fighter statistics
// @Type FighterStats
type FighterStats struct {
    Name        string `json:"name"`
    Age         string `json:"age"`
    Division    string `json:"divsion"`
    Height      string `json:"physique.height"`
    Weight      string `json:"physique.weight"`
    Reach       string `json:"physique.reach"`
    LegReach    string `json:"physique.leg reach"`
    Knockouts   string `json:"knockouts"`
    Submissions string `json:"submissions"`
    Record      struct {
        Wins  int `json:"wins"`
        Losses int `json:"losses"`
        Draws int `json:"draws"`
    } `json:"record"`
    Strikes struct {
        Accuracy     int    `json:"accuracy"`
        Attempted    string `json:"attempted"`
        Landed       string `json:"landed"`
    } `json:"Stikes"`
    Takedowns struct {
        Accuracy     int    `json:"accuracy"`
        Attempted    string `json:"attempted"`
        Landed       string `json:"landed"`
    } `json:"Takedowns"`
}
