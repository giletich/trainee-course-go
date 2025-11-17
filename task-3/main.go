package main

import "fmt"

func main() {
	strings := []string{"a", "aa", "aa", "a", "b", "c", "c", "d"}
	strings = ReturnUniqueValaues(strings)
	fmt.Println(strings)
}

func ReturnUniqueValaues(s []string) []string {

	mapa := make(map[string]int)
	for _, key := range s {
		mapa[key]++
	}

	newSlice := []string{}

	for _, value := range s {
		if _, ok := mapa[value]; ok {
			newSlice = append(newSlice, value)
			delete(mapa, value)
		}
	}

	return newSlice
}