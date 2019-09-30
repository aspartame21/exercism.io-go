package proverb

import "fmt"

func Proverb(rhyme []string) []string {

	var rhymeLength = len(rhyme)
	var res []string

	if rhymeLength == 0 {
		return []string{}
	}

	for i := 0; i < rhymeLength-1; i++ {
		res = append(res, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}

	return append(res, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
}
