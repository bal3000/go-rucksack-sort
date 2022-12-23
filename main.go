package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sacks, err := loadInventory("./rucksacks.txt")
	if err != nil {
		panic(err)
	}

	lm := createLetterMap()

	total := 0
	for _, s := range sacks {
		c := findDups(s)
		if c == "" {
			continue
		}

		total += lm.getPriority(c)
	}

	fmt.Println(total)
}

type letterMap map[string]int

func createLetterMap() letterMap {
	lm := make(map[string]int)

	lower := 97
	for i := 0; i < 26; i++ {
		c := string(rune(i + lower))
		lm[c] = i + 1
	}

	upper := 65
	for i := 0; i < 26; i++ {
		c := string(rune(i + upper))
		lm[c] = i + 27
	}

	return lm
}

func (lm letterMap) getPriority(letter string) int {
	return lm[letter]
}

func findDups(sacks []string) string {
	for _, fr := range sacks[0] {
		for _, sf := range sacks[1] {
			if fr == sf {
				return string(fr)
			}
		}
	}

	return ""
}

func loadInventory(filename string) ([][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	str := strings.Split(
		strings.Replace(
			string(data),
			"\r\n",
			"\n",
			-1,
		),
		"\n",
	)

	sacks := make([][]string, 0)

	for _, line := range str {
		half := len(line) / 2
		first := line[:half]
		second := line[half:]

		sacks = append(sacks, []string{first, second})
	}

	return sacks, nil
}
