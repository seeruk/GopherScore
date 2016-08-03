package wow

// -- Character Models

// Character represents an in-game character's basic information.
type Character struct {
	Name                string      `json:"name"`
	Realm               string      `json:"realm"`
	Battlegroup         string      `json:"realm"`
	Class               int         `json:"class"`
	Race                int         `json:"race"`
	Gender              int         `json:"gender"`
	Level               int         `json:"level"`
	AchievementPoints   int         `json:"achievementPoints"`
	Faction             int         `json:"faction"`
	TotalHonorableKills int         `json:"totalHonorableKills"`
	Items               Items       `json:"items"`
	Professions         Professions `json:"professions"`
	Progression         Progression `json:"progression"`
}

// ResolveClassName attempts to resolve a class ID into a class name.
func (c *Character) ClassName() string {
	switch c.Class {
	case 1:
		return "Warrior"
	case 2:
		return "Paladin"
	case 3:
		return "Hunter"
	case 4:
		return "Rogue"
	case 5:
		return "Priest"
	case 6:
		return "Death Knight"
	case 7:
		return "Shaman"
	case 8:
		return "Mage"
	case 9:
		return "Warlock"
	case 10:
		return "Monk"
	case 11:
		return "Druid"
	case 12: // Guessing
		return "Demon Hunter"
	}

	return "Unknown"
}

// ResolveFaction attempts to resolve a faction ID into a faction name.
func (c *Character) FactionName() string {
	switch c.Faction {
	case 0:
		return "Alliance"
	case 1:
		return "Horde"
	}

	return "Unknown"
}

// ResolveRace attempts to resolve a race ID into a race name.
func (c *Character) RaceName() string {
	switch c.Race {
	case 1:
		return "Human"
	case 2:
		return "Orc"
	case 3:
		return "Dwarf"
	case 4:
		return "Night Elf"
	case 5:
		return "Undead"
	case 6:
		return "Tauren"
	case 7:
		return "Gnome"
	case 8:
		return "Troll"
	case 9:
		return "Goblin"
	case 10:
		return "Blood Elf"
	case 11:
		return "Draenei"
	case 22:
		return "Worgen"
	case 25, 26:
		return "Pandaren"
	}

	return "Unknown"
}

// -- Item Models

// Items represents an overview of a character's items.
type Items struct {
	AverageItemLevel         int  `json:"averageItemLevel"`
	AverageItemLevelEquipped int  `json:"averageItemLevelEquipped"`
	Head                     Item `json:"head"`
	Neck                     Item `json:"neck"`
	Shoulder                 Item `json:"shoulder"`
	Back                     Item `json:"back"`
	Chest                    Item `json:"chest"`
	Wrist                    Item `json:"wrist"`
	Hands                    Item `json:"hands"`
	Waist                    Item `json:"waist"`
	Legs                     Item `json:"legs"`
	Feet                     Item `json:"feet"`
	Finger1                  Item `json:"finger1"`
	Finger2                  Item `json:"finger2"`
	Trinket1                 Item `json:"trinket1"`
	Trinket2                 Item `json:"trinket2"`
	MainHand                 Item `json:"mainHand"`
	OffHand                  Item `json:"offHand"`
}

// Item represents a piece of gear a character can wear.
type Item struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Quality   int    `json:"quality"`
	ItemLevel int    `json:"itemLevel"`
}

// -- Professions Models

// Professions represents a character's professions (primary, and secondary).
type Professions struct {
	Primary   []Profession `json:"primary"`
	Secondary []Profession `json:"secondary"`
}

// Profession represents a character's profession, and progress through it.
type Profession struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Rank    int    `json:"rank"`
	MaxRank int    `json:"max"`
}

// -- Progression Models

// Progression represents a character's progression.
type Progression struct {
	Raids               []Raid `json:"raids"`
	TotalHonorableKills int    `json:"totalHonorableKills,omitempty"`
}

// Raid represents a character's progression within a given raid.
type Raid struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	LFRClears    int    `json:"lfr,omitempty"`
	NormalClears int    `json:"normal,omitempty"`
	HeroicClears int    `json:"heroic,omitempty"`
	MythicClears int    `json:"mythic,omitempty"`
	Bosses       []Boss `json:"bosses"`
}

// Boss represents character's progression on a boss within a raid.
type Boss struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	LFRKills    int    `json:"lfrKills,omitempty"`
	NormalKills int    `json:"normalKills,omitempty"`
	HeroicKills int    `json:"heroicKills,omitempty"`
	MythicKills int    `json:"mythicKills,omitempty"`
}
