package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "lai",
		"course":  "golang",
		"site":    "52go",
		"quality": "notbad",
	}
	m2 := make(map[string]int) // m2== empty map
	var m3 map[string]int      // m3 == nil
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting vlaues")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if coursName, ok := m["cours"]; ok {
		fmt.Println(coursName, ok)
	} else {
		fmt.Println("keys does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	fmt.Println(m)
	name1, ok := m["name"]
	fmt.Println(name1, ok)
}
