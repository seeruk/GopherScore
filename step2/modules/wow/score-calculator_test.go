package wow

import "testing"

func TestItemsScoreCalculator(t *testing.T) {
	var calculator ItemsScoreCalculator
	var character Character
	var items Items

	item := Item{
		ID:        1,
		Name:      "Test Item",
		Quality:   2,
		ItemLevel: 600,
	}

	items.Head = item
	items.Shoulder = item

	character.Items = items

	t.Run("it should calculate a score based item level, and quality, with a weighting", func(t *testing.T) {
		expected := 2400
		actual := calculator.Calculate(character)

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("it should double main-hand score if there's no off-hand", func(t *testing.T) {
		items.MainHand = item
		character.Items = items

		expected := 4800
		actual := calculator.Calculate(character)

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})

	t.Run("it should not double main-hand score if there is an off-hand", func(t *testing.T) {
		items := Items{}
		items.MainHand = item
		items.OffHand = item
		character.Items = items

		expected := 2400
		actual := calculator.Calculate(character)

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	})
}

func TestItemsScoreCalculatorTable(t *testing.T) {
	var calculator ItemsScoreCalculator
	var character Character

	type Test struct {
		Name     string
		Expected int
		Items    Items
	}

	item := Item{
		ID:        1,
		Name:      "Test Item",
		Quality:   2,
		ItemLevel: 600,
	}

	tests := []Test{
		{
			Name:     "it should calculate a score based item level, and quality, with a weighting",
			Expected: 2400,
			Items: Items{
				Head:     item,
				Shoulder: item,
			},
		},
		{
			Name:     "it should double main-hand score if there's no off-hand",
			Expected: 4800,
			Items: Items{
				Head:     item,
				Shoulder: item,
				MainHand: item,
			},
		},
		{
			Name:     "it should not double main-hand score if there is an off-hand",
			Expected: 2400,
			Items: Items{
				MainHand: item,
				OffHand:  item,
			},
		},
	}

	for _, test := range tests {
		character.Items = test.Items

		t.Run(test.Name, func(t *testing.T) {
			actual := calculator.Calculate(character)

			if actual != test.Expected {
				t.Errorf("Expected %d, got %d", test.Expected, actual)
			}
		})
	}
}
