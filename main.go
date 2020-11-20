package main

import (
	"fmt"
	"gitlab.com/skywalker_etw/DaimlerTssMerge/merge"
)

//Deno Function
func main() {
	l1 := merge.List{{25, 30}, {2, 19}, {14, 23}, {4, 8}}
	nl := merge.Merge(l1)
	fmt.Println(nl)
}
