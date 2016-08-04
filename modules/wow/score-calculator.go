package wow

// ScoreCalculator contains methods for calculating player's
type ScoreCalculator interface {
	Calculate(Character) int
}

// AggregateScoreCalculator
type AggregateScoreCalculator struct {
	calculators []ScoreCalculator
}

// AddCalculator adds a score calculator to get a score from.
func (c *AggregateScoreCalculator) AddCalculator(calculator ScoreCalculator) {
	c.calculators = append(c.calculators, calculator)
}

// Calculate a score based on a combination of the results of other score calculators.
func (c *AggregateScoreCalculator) Calculate(character Character) int {
	score := 0

	for _, calculator := range c.calculators {
		score += calculator.Calculate(character)
	}

	return score
}

// AchievementScoreCalculator calculates a character's achievement score, based on the total number
// of achievement points.
type AchievementScoreCalculator struct{}

// Calculate a score for the character's achievement points.
func (c AchievementScoreCalculator) Calculate(character Character) int {
	return character.AchievementPoints / 2
}

// ItemsScoreCalculator calculates a character's item's score, based on item level and quality.
type ItemsScoreCalculator struct{}

// Calculate a score for the character's items.
func (c ItemsScoreCalculator) Calculate(character Character) int {
	var score int

	items := character.Items
	weighting := 1

	score += calculateItemScore(items.Head, weighting)
	score += calculateItemScore(items.Neck, weighting)
	score += calculateItemScore(items.Shoulder, weighting)
	score += calculateItemScore(items.Back, weighting)
	score += calculateItemScore(items.Chest, weighting)
	score += calculateItemScore(items.Wrist, weighting)
	score += calculateItemScore(items.Hands, weighting)
	score += calculateItemScore(items.Waist, weighting)
	score += calculateItemScore(items.Legs, weighting)
	score += calculateItemScore(items.Feet, weighting)
	score += calculateItemScore(items.Finger1, weighting)
	score += calculateItemScore(items.Finger2, weighting)
	score += calculateItemScore(items.Trinket1, weighting)
	score += calculateItemScore(items.Trinket2, weighting)

	// 2-handed weapons should have double the score.
	mainHandWeighting := 2

	if items.OffHand.ID != 0 {
		mainHandWeighting = weighting

		score += calculateItemScore(items.OffHand, weighting)
	}

	score += calculateItemScore(items.MainHand, mainHandWeighting)

	return score
}

func calculateItemScore(item Item, weighting int) int {
	return (item.ItemLevel * weighting) * item.Quality
}

// ProfessionsScoreCalculator calculates a character's title score.
type ProfessionsScoreCalculator struct{}

// Calculate a score for the character's professions.
func (c ProfessionsScoreCalculator) Calculate(character Character) int {
	return calculateProfessionsScore(character.Professions.Primary, 30) +
		calculateProfessionsScore(character.Professions.Secondary, 10)
}

// calculateProfessionsScores takes a collection of professions and calculates a score, with the
// given weighting.
func calculateProfessionsScore(professions []Profession, weighting int) int {
	var score int

	for _, profession := range professions {
		if profession.MaxRank == 0 {
			continue
		}

		score += ((100 / profession.MaxRank) * profession.Rank) * weighting
	}

	return score
}

// TitleScoreCalculator calculates a character's title score.
type ProgressionScoreCalculator struct{}

// Calculate a score for the character's progression.
func (c ProgressionScoreCalculator) Calculate(character Character) int {
	var score int

	for _, raid := range character.Progression.Raids {
		score += raid.LFRClears * 100
		score += raid.NormalClears * 200
		score += raid.HeroicClears * 300
		score += raid.MythicClears * 400
	}

	return score
}
