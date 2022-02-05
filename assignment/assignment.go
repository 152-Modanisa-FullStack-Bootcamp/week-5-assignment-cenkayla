package assignment

import (
	"bytes"
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	sumUint64 := uint64(x) + uint64(y)
	maxUint32 := 1<<32 - 1
	if sumUint64 > uint64(maxUint32) {
		return x + y, true
	}
	return x + y, false
}

func CeilNumber(f float64) float64 {
	return math.Ceil(4*f) / 4
}

func AlphabetSoup(s string) string {
	var r []rune

	for _, runeValue := range s {
		r = append(r, runeValue)
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}

func StringMaskConcat(s string, n uint) string {
	var count uint
	var mask string
	if uint(len(s)) == 0 {
		return "*"
	}

	if uint(len(s)) == n {
		mask = strings.Repeat("*", int(n))
		return mask
	}

	if n > uint(len(s)) {
		mask = strings.Repeat("*", len(s))
		return mask
	}

	for i := range s {
		if n <= count {
			mask += "*"
		} else {
			mask += string(s[i])
		}
		count++
	}
	return mask
}

func StringMaskWithBuffer(s string, n uint) string {
	var count uint
	var buf bytes.Buffer
	if uint(len(s)) == 0 {
		return "*"
	}

	if uint(len(s)) == n {
		buf.WriteString(strings.Repeat("*", int(n)))
		return buf.String()
	}

	if n > uint(len(s)) {
		buf.WriteString(strings.Repeat("*", len(s)))
		return buf.String()
	}

	for _, v := range s {
		if n <= count {
			buf.WriteString("*")
		} else {
			buf.WriteString(string(v))
		}
		count++
	}
	return buf.String()
}

func WordSplit(arr [2]string) string {
	var (
		count   int
		words   []string
		numbers []int
		buf     bytes.Buffer
	)

	sliceWords := strings.Split(arr[1], ",")

	for _, word := range sliceWords {
		if strings.Contains(arr[0], word) {
			indexNumber := strings.Index(arr[0], word)
			numbers = append(numbers, indexNumber)

			if indexNumber == 0 {
				words = append(words, arr[0][:len(word)])
			} else {
				words = append(words, arr[0][indexNumber:])
			}
			count++
		}
	}

	if count != 2 {
		return "not possible"
	}

	//Verilen kelimelerin sırasının farklı olduğu durum için eklendi.
	if numbers[0] > numbers[1] {
		buf.WriteString(words[1])
		buf.WriteString(",")
		buf.WriteString(words[0])
	} else {
		buf.WriteString(words[0])
		buf.WriteString(",")
		buf.WriteString(words[1])
	}

	return buf.String()
}

func VariadicSet(i ...interface{}) []interface{} {
	keys := make(map[interface{}]struct{})
	list := []interface{}{}

	for _, v := range i {
		if _, value := keys[v]; !value {
			keys[v] = struct{}{}
			list = append(list, v)
		}
	}
	return list
}
