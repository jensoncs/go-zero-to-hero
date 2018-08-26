package main

import "fmt"

func main() {
	grp1 := make(map[string]int)
	grp1["tom"] = 22
	grp1["sam"] = 25

	grp2 := make(map[string]int)
	grp2["tom"] = 22
	grp2["sam"] = 25

	compareStatus := compare(grp1, grp2)
	if compareStatus {
		fmt.Println("The given maps are same")
	} else {
		fmt.Println("The given maps are not equal")
	}
}

func compare(grp1, grp2 map[string]int) bool {
	flag := true
	for name, _ := range grp1 {
		if grp1[name] != grp2[name] {
			flag = false
		} else {
			for name, _ := range grp2 {
				if grp2[name] != grp1[name] {
					flag = false
				}
			}
		}
	}
	return flag
}
