package tools_stringResolve

import (
	"fmt"
)

func RangeStringSlice(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func RangeMap(m map[int]string) {
	// tMap := make(map[int]string)

	for i := 0; i < len(m)/2; i++ {
		fmt.Println(i, m[i])
	}

}
