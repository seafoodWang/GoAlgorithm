package arr

import (
	"fmt"
	"testing"
)

func TestDescCity(t *testing.T) {

	paths := [][]string{{"D", "B"}}
	fmt.Println(destCity(paths))
}

func destCity(paths [][]string) string {
	from := make(map[string]struct{})
	to := make(map[string]struct{})
	desc := ""
	for i := 0; i < len(paths); i++ {
		path := paths[i]
		f := path[0]
		t := path[1]
		from[f] = struct{}{}
		to[t] = struct{}{}
	}
	for t := range to {
		_, ok := from[t]
		if !ok {
			desc = t
			break
		}
	}
	return desc
}
