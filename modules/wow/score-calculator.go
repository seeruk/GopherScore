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
	return 11
}

// ProfessionsScoreCalculator calculates a character's title score.
type ProfessionsScoreCalculator struct{}

// Calculate a score for the character's professions.
func (c ProfessionsScoreCalculator) Calculate(character Character) int {
	return 111
}

// TitleScoreCalculator calculates a character's title score.
type ProgressionScoreCalculator struct{}

// Calculate a score for the character's progression.
func (c ProgressionScoreCalculator) Calculate(character Character) int {
	return 1111
}
