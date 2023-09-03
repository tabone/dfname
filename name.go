package dfname

import (
	"fmt"
	"math/rand"
	"strings"
)

/*
Name that is generated.
*/
type Name struct {
	// FirstName represents the first name of the generated name.
	FirstName *Noun

	// LastName represents the last name of the generated name.
	LastName [2]*Noun

	// language that the name was created for.
	language Language
}

/*
NewName is used to create a new name.
*/
func NewName(seed int64, language Language) *Name {
	random := rand.New(rand.NewSource(seed))

	return &Name{
		language:  language,
		FirstName: nouns[random.Intn(len(nouns))],
		LastName: [2]*Noun{
			nouns[random.Intn(len(nouns))],
			nouns[random.Intn(len(nouns))],
		},
	}
}

func capitalize(str string) string {
	return fmt.Sprintf(
		"%s%s",
		strings.ToUpper(fmt.Sprintf("%c", str[0])),
		strings.ToLower(str[1:]),
	)
}

/*
String is used to stringify the name (first and last name) in the language
specified during initialization.
*/
func (name *Name) String() string {
	return fmt.Sprintf(
		"%s %s",
		capitalize(name.FirstName.Translate(name.language)),
		capitalize(
			fmt.Sprintf(
				"%s%s",
				name.LastName[0].Translate(name.language),
				name.LastName[1].Translate(name.language),
			),
		),
	)
}

/*
Translate is used to stringify the name (first and last name) while translating
the last name to a different language.
*/
func (name *Name) Translate(language Language) string {
	return fmt.Sprintf(
		"%s %s",
		capitalize(name.FirstName.Translate(name.language)),
		capitalize(
			fmt.Sprintf(
				"%s%s",
				name.LastName[0].Translate(language),
				name.LastName[1].Translate(language),
			),
		),
	)
}
