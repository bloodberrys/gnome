package main

import (
	"engboarding/helper" // we are going to import the custom package called helper
	"fmt"
)

// Dynamic array, if we create this array outside the main function, we can't use the walrus operator, instead we can use var.
var docs = []string{}

// Integer stuffs
const newHiresMaxPoint int = 100000

var newHiresInitialPoint uint = 1000

func main() {
	fmt.Println("Engboarding, the engineering on-boarding solution for your business.")
	fmt.Println("Get your free trial, and see how happy your new hires")

	// String var using walrus operator (:=)
	fullName := "Alfian Firmansyah"
	fmt.Printf("Hello %v, welcome\n", fullName)

	helper.Wg.Add(1)
	go helper.SendEmail(fullName)

	fmt.Printf("Your points: %v\n", newHiresInitialPoint)
	fmt.Printf("Max points: %v\n\n\n\n", newHiresMaxPoint)

	docs = append(docs, "How to maximize your pruductivity")
	docs = append(docs, "Engineering Values")
	docs = append(docs, "Infra and Site Reliability Engineer")
	docs = append(docs, "Server Migrations")

	// All the custom package, started by the capitalize letter, it means the function will be imported to all package.
	helper.GetDocs(docs)

	// Conditional Statement
	if len(docs) >= 3 {
		fmt.Println("You exceed a max daily docs")
	} else {
		fmt.Print("You need to add daily docs")
	}
	helper.Wg.Wait()
}
