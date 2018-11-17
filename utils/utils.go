package utils

import (
	"math/rand"
	"strings"
)

func ContainsOne(str string, test []string) bool {
	var contains bool
	contains = false

	for _, t := range test {
		if strings.Contains(str, t) {
			contains = true
		}
	}

	return contains
}

func PickOne(list []string) string {
	count := int32(len(list))

	i := rand.Int31n(count)

	return list[i]
}
