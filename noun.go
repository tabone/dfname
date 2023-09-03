package dfname

/*
Noun is the building blocks of a name.
*/
type Noun struct {
	/*
		English translation for the noun.
	*/
	English string

	/*
		Dwarven translation for the noun.
	*/
	Dwarven string

	/*
		Elvish translation for the noun.
	*/
	Elvish string

	/*
		Goblin translation for the noun.
	*/
	Goblin string

	/*
		Human translation for the noun.
	*/
	Human string
}

/*
Translate noun to a particular language.
*/
func (noun *Noun) Translate(language Language) string {
	switch language {
	case DwarvenLanguage:
		return noun.Dwarven
	case ElvishLanguage:
		return noun.Elvish
	case GoblinLanguage:
		return noun.Goblin
	case HumanLanguage:
		return noun.Human
	default:
		return noun.English
	}
}
