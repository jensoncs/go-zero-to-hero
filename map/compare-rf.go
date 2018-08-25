package main

import "fmt"
import "reflect"

func main() {
	grp1 := make(map[string]int)
	grp1["tom"] = 22
	grp1["sam"] = 25

	grp2 := make(map[string]int)
	grp2["tom"] = 22
	grp2["sam"] = 25

	fmt.Println(reflect.DeepEqual(grp1, grp2))
}
