package dfname

/*
Language that the noun is written in.
*/
type Language int

/*
Available languages.
*/
const (
	EnglishLanguage Language = iota
	DwarvenLanguage
	ElvishLanguage
	GoblinLanguage
	HumanLanguage
)
