package leetcode

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	record, res := map[string][]string{}, [][]string{}

	for _, str := range strs {
		runes := []rune(str)
		sort.Sort(sortRunes(runes))
		sstr := record[string(runes)]
		sstr = append(sstr, str)
		record[string(runes)] = sstr
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}

//func groupAnagrams(strs []string) [][]string {
//	record, res := map[string][]string{}, [][]string{}
//
//	for _, str := range strs {
//		b := []byte(str)
//		sort.Slice(b, func(i, j int) bool {
//			return b[i] < b[j]
//		})
//		key := string(b)
//		record[key] = append(record[key], str)
//	}
//	for _, v := range record {
//		res = append(res, v)
//	}
//	return res
//}
//
