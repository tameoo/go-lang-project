package main

import "fmt"

func main() {
	testMap := map[string]int{
		"test0": 1,
		"test1": 2,
		"test2": 3,
	}

	fmt.Println(Anagram("test", "test"))
	// fmt.Println(FindEvens([]int{1, 2, 3, 4, 5}))
	fmt.Println(SliceProduct([]int{1, 2, 3}))
	fmt.Println(Unique([]int{1, 2, 3, 3, 4, 4, 5, 5, 6, 6, 6, 6}))
	fmt.Println(InvertMap(testMap))
	fmt.Println(TopCharacters("tameooo", 2))
}

func Anagram(s1, s2 string) bool {

	if len(s1) != len(s2) {
		return false
	}

	list := make(map[string]int)

	for i := 0; i < len(s1); i++ {
		list[string(s1[i])]++
	}

	for i := 0; i < len(s2); i++ {
		list[string(s2[i])]--
	}

	for i := 0; i < len(s1); i++ {
		if list[string(s1[i])] != 0 {
			return false
		}
	}

	return true
}

func FindEvens(e []int) []int {
	for index, value := range e {
		if value%2 != 0 {
			e = append(e[:index], e[index+1:]...)
		}
	}
	return e
}

func SliceProduct(e []int) int {
	total := 1

	for _, value := range e {
		total *= value
	}

	return total
}

func Unique(e []int) []int {
	list := make(map[int]bool)
	store := []int{}

	for _, entry := range e {
		if _, value := list[entry]; !value {
			list[entry] = true
			store = append(store, entry)
		}
	}

	return store
}

func InvertMap(kv map[string]int) map[int]string {
	list := map[int]string{}

	for key, value := range kv {
		list[value] = key
	}

	return list
}

func TopCharacters(s string, k int) map[rune]int {
	list := map[rune]int{}
	result := map[rune]int{}

	for _, char := range s {
		if value, ok := list[char]; ok {
			list[char] = value + 1
		} else {
			list[char] = 1
		}
	}

	for key, value := range list {
		if value > k {
			result[key] = value
		}
	}

	return result
}
