package arr

import (
	"bytes"
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/integer-to-english-words/description/
func TestIntToEnglish(t *testing.T) {
	fmt.Println(numberToWords(123456789))
}

func numberToWords(num int) string {
	if num == 0 {
		return "zero"
	}
	hundred := 100
	thousand := 10 * hundred
	million := 1000 * thousand
	billion := 1000 * million
	dict := getMap()
	buf := bytes.Buffer{}
	b := num / billion
	buf.WriteString(getDesc(dict, b, "Billion"))
	m := num / million % 1000
	buf.WriteString(getDesc(dict, m, "Million"))
	t := num / thousand % 1000
	buf.WriteString(getDesc(dict, t, "Thousand"))
	h := num / hundred % 10
	buf.WriteString(getDesc(dict, h, "Hundred"))
	ten := num % 100
	if _, ok := dict[ten]; ok {
		buf.WriteString(getDesc(dict, ten, ""))
	} else {
		ten = num / 10 % 10
		buf.WriteString(getDesc(dict, ten*10, ""))
		one := num % 10
		buf.WriteString(getDesc(dict, one, ""))
	}
	return buf.String()
}

func getDesc(dict map[int]string, n int, unit string) string {

	v, ok := dict[n]
	if !ok {
		return ""
	}
	v += " "
	if unit == "" {
		return v
	}
	return v + unit + " "
}

func getMap() map[int]string {
	dict := make(map[int]string)
	dict[1] = "One"
	dict[2] = "Two"
	dict[3] = "Three"
	dict[4] = "Four"
	dict[5] = "Five"
	dict[6] = "Six"
	dict[7] = "Seven"
	dict[8] = "Eight"
	dict[9] = "Nine"
	dict[10] = "Ten"
	dict[11] = "Eleven"
	dict[12] = "Twelve"
	dict[13] = "Thirteen"
	dict[14] = "Fourteen"
	dict[15] = "Fifteen"
	dict[16] = "Sixteen"
	dict[17] = "Seventeen"
	dict[18] = "Eighteen"
	dict[19] = "Nineteen"
	dict[20] = "Twenty"
	dict[30] = "Thirty"
	dict[40] = "Forty"
	dict[50] = "Fifty"
	dict[60] = "Sixty"
	dict[70] = "Seventy"
	dict[80] = "Eighty"
	dict[90] = "Ninety"
	return dict
}
