package dfname

import (
	"testing"
)

func TestSeededInitialization(t *testing.T) {
	nameOne := NewName(371993, DwarvenLanguage)
	nameTwo := NewName(371993, EnglishLanguage)

	nameOneNouns := []*Noun{
		nameOne.FirstName, nameOne.LastName[0], nameOne.LastName[1],
	}

	nameTwoNouns := []*Noun{
		nameTwo.FirstName, nameTwo.LastName[0], nameTwo.LastName[1],
	}

	for i, nameOneNoun := range nameOneNouns {
		if nameOneNoun != nameTwoNouns[i] {
			t.Errorf("same seeded names should return the same name")
		}
	}
}

func TestString(t *testing.T) {
	name := NewName(371993, DwarvenLanguage)

	actualName := name.String()
	expectedName := "Arceth Gumùrngegdol"

	if actualName != expectedName {
		t.Errorf(
			"expected name to be '%s' but it is '%s'",
			expectedName,
			actualName,
		)
	}
}

func TestTranslate(t *testing.T) {
	name := NewName(371993, DwarvenLanguage)

	actualName := name.Translate(EnglishLanguage)
	expectedName := "Arceth Idlenessdepression"

	if actualName != expectedName {
		t.Errorf(
			"expected name to be '%s' but it is '%s'",
			expectedName,
			actualName,
		)
	}
}
