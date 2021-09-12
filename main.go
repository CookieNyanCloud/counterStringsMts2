package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main()  {
	fmt.Println(Counter("aaabbbccccc"))
	fmt.Println(Counter("aaabbbcccccaaaaa"))
	fmt.Println(Counter("zzzzcccUUUuu"))
	fmt.Println(Counter("ЯЯЯБББддд"))
}

func Counter(s string) string {
	strmap:=make(map[string]int)
	keys:= make([]string,0, len(s))
	var out string
	for _, c:= range s{
		strmap[string(c)]++
	}
	for k := range strmap {
		keys = append(keys,k)
	}
	sort.Strings(keys)
	for _,v:= range keys{
		out += v
		out += strconv.Itoa(strmap[v])
	}
	return out
}