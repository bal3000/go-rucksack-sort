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
	ss := seperateSacks(sacks)
	elfGroups := splitIntoGroups(sacks)

	lm := createLetterMap()

	total := 0
	for _, s := range ss {
		c := findDups(s)
		if c == "" {
			continue
		}

		total += lm.getPriority(c)
	}

	fmt.Printf("total priority is %d\n", total)

	total = 0
	for _, g := range elfGroups {
		b := findBadge(g)
		if b == "" {
			panic("No badge found")
		}

		total += lm.getPriority(b)
	}

	fmt.Printf("total group priority is %d\n", total)
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

func findBadge(group []string) string {
	for _, fr := range group[0] {
		for _, sr := range group[1] {
			for _, tr := range group[2] {
				if fr == sr && sr == tr {
					return string(fr)
				}
			}
		}
	}
	return ""
}

func splitIntoGroups(sacks []string) [][]string {
	i := 1
	groups := make([][]string, 0)
	group := make([]string, 0)
	for _, s := range sacks {
		group = append(group, s)

		if i == 3 {
			i = 1
			groups = append(groups, group)
			group = make([]string, 0)
		} else {
			i++
		}
	}

	return groups
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

func seperateSacks(elves []string) [][]string {
	sacks := make([][]string, 0)

	for _, elf := range elves {
		half := len(elf) / 2
		first := elf[:half]
		second := elf[half:]

		sacks = append(sacks, []string{first, second})
	}

	return sacks
}

func loadInventory(filename string) ([]string, error) {
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

	return str, nil
}
