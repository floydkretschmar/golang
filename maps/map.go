package maps

import "fmt"

func PrintMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
