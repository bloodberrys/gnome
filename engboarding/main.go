package main

import "fmt"

func main() {

	fmt.Println("Engboarding, the engineering on-boarding solution for your business.")
	fmt.Println("Get your free trial, and see how happy your new hires")

	// String var using walrus operator (:=)
	fullName := "Alfian Firmansyah"
	fmt.Printf("Hello %v, welcome\n", fullName)

	// Integer stuffs
	const newHiresMaxPoint int = 100000
	var newHiresInitialPoint uint = 1000

	fmt.Printf("Your points: %v\n", newHiresInitialPoint)
	fmt.Printf("Max points: %v\n", newHiresMaxPoint)

	// Dynamic array
	docs := []string{}

	docs = append(docs, "Nice")

	fmt.Printf("%v", docs)

}
