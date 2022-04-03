package helper

import (
	"fmt"
	"sync"
	"time"
)

var Wg = sync.WaitGroup{}

// import by providing the first letter capitalize
func GetDocs(docs []string) {
	// For loops printing
	// the index can be replaced by the _, meaning blank identifier, if we don't need the index.
	for index, doc := range docs {
		if index == 1 {
			fmt.Println("Skipped")
			continue
		} else {
			fmt.Printf("%v) %v\n", index+1, doc)
		}
	}
}

func SendEmail(fullName string) {
	time.Sleep(10 * time.Second) // 10 second
	fmt.Printf("Email has been sent to %v", fullName)
	Wg.Done()
}
