package wow

type CharacterDataResolver struct{}

// ResolveClassName attempts to resolve a class ID into a class name.
func (r CharacterDataResolver) ResolveClassName(id int) string {
	switch id {
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
func (r CharacterDataResolver) ResolveFaction(id int) string {
	switch id {
	case 0:
		return "Alliance"
	case 1:
		return "Horde"
	}

	return "Unknown"
}

// ResolveRace attempts to resolve a race ID into a race name.
func (r CharacterDataResolver) ResolveRace(id int) string {
	switch id {
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
