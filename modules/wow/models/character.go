package models

// Character represents an in-game character's basic information.
type Character struct {
	Name                string `json:"name"`
	Realm               string `json:"realm"`
	Battlegroup         string `json:"realm"`
	Class               int    `json:"class"`
	Race                int    `json:"race"`
	Gender              int    `json:"gender"`
	Level               int    `json:"level"`
	AchievementPoints   int    `json:"achievementPoints"`
	Faction             int    `json:"faction"`
	TotalHonorableKills int    `json:"totalHonorableKills"`
}
