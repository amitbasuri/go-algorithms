package hashMaps

import "fmt"

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {
	magMap := make(map[string]int)
	noteMap := make(map[string]int)
	noteMapMultipleWords := make(map[string]struct{})

	for _, c := range magazine {
		magMap[c]++
	}

	for _, c := range note {
		if _, ok := magMap[c]; !ok {
			fmt.Println("No")
			return
		}
		noteMap[c]++
		if noteMap[c] > 1 {
			noteMapMultipleWords[c] = struct{}{}
		}
	}

	for c, _ := range noteMapMultipleWords {
		if noteMap[c] > magMap[c] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
	return

}
