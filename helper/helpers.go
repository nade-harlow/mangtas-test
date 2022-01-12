package helper

import (
	"mangtas-test/models"
	"regexp"
	"sort"
)

func Counter(txt string) []models.Order {
	var sorter []models.Order
	texts := regexp.MustCompile("[^0-9a-zA-Z]+").Split(txt, -1)

	m := map[string]int{}
	for _, word := range texts {
		m[word]++
	}

	for key, value := range m {
		sorter = append(sorter, models.Order{Key: key, Value: value})
	}

	sort.Slice(sorter, func(i, j int) bool {
		return sorter[i].Value > sorter[j].Value
	})

	return sorter[:10]
}
