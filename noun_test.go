package dfname

import (
	"testing"
)

func TestNoun(t *testing.T) {
	noun := &Noun{
		English: "test-english",
		Dwarven: "test-dwarven",
		Elvish:  "test-elvish",
		Goblin:  "test-goblin",
		Human:   "test-human",
	}

	type testCase struct {
		expected    string
		actual      string
		description string
	}

	testCases := []testCase{
		{
			description: "english translation",
			expected:    "test-english",
			actual:      noun.Translate(EnglishLanguage),
		},
		{
			description: "dwarven translation",
			expected:    "test-dwarven",
			actual:      noun.Translate(DwarvenLanguage),
		},
		{
			description: "elvish translation",
			expected:    "test-elvish",
			actual:      noun.Translate(ElvishLanguage),
		},
		{
			description: "goblin translation",
			expected:    "test-goblin",
			actual:      noun.Translate(GoblinLanguage),
		},
		{
			description: "human translation",
			expected:    "test-human",
			actual:      noun.Translate(HumanLanguage),
		},
		{
			description: "english translation",
			expected:    "test-english",
			actual:      noun.Translate(10000),
		},
	}

	for _, testCase := range testCases {
		if testCase.actual != testCase.expected {
			t.Errorf(
				"%s failed. expected %s and received %s",
				testCase.description,
				testCase.expected,
				testCase.actual,
			)
		}
	}
}
