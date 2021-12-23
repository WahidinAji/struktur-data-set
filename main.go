package main

import (
	"fmt"
	"struktur-data-set/numeric"
	string2 "struktur-data-set/string"
)

func main() {
	numeric.Start()
	fmt.Println("string version : \n")
	string2.Start()
}


//func add(value string) bool {
//	if contains(value) {
//		return false
//	} else {
//		return true
//	}
//}
//
//func contains(value string) bool {
//	var items []string
//	for _, item := range items {
//		if value == item {
//			return true
//		}
//	}
//	return false;
//}
//
//func size() int {
//	return 0;
//}
//func remove(value string) bool {
//	return true;
//}