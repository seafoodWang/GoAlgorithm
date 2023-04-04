package list

import (
	"fmt"
	"testing"
)

//https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

func TestPhoneNumber(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if "" == digits {
		return []string{}
	}
	mapping := getPhoneNumberMapping()
	init := []string{""}

	length := len(digits)
	//数字
	for i := 0; i < length; i++ {
		data := mapping[string(digits[i])]
		tmp := make([]string, 0)
		//数字对应的字母
		for j := 0; j < len(data); j++ {
			//字母和现有的init进行拼接
			cur := data[j]
			for k := 0; k < len(init); k++ {
				tmp = append(tmp, init[k]+cur)
			}
		}
		init = tmp
	}
	return init
}

func getPhoneNumberMapping() map[string][]string {
	mapping := make(map[string][]string)
	mapping["2"] = []string{"a", "b", "c"}
	mapping["3"] = []string{"d", "e", "f"}
	mapping["4"] = []string{"g", "h", "i"}
	mapping["5"] = []string{"j", "k", "l"}
	mapping["6"] = []string{"m", "n", "o"}
	mapping["7"] = []string{"p", "q", "r", "s"}
	mapping["8"] = []string{"t", "u", "v"}
	mapping["9"] = []string{"w", "x", "y", "z"}
	return mapping
}
