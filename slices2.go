package main

import (
	"fmt"
	"reflect"
)


/// docs
// A slice is a reference to a contiguous segment of an array. 
// It is a variable-length sequence which stores elements of a similar type.
func main() {
	languages := [9]string{"Python", "Java", "Go", "C", "C++", "C#", "JavaScript", "Ruby", "Rust"}

	allLangs := languages[:]                   // copy the entire array
	fmt.Print(reflect.TypeOf(allLangs).Kind()) // slice

	// create slices contianing web frameworks
	webFrameworks := []string{
		"React", "Vue", "Angular", "Svelt",
		"Slim", "Laravel", "Django", "Flask",
	}

	jsFrameworks := webFrameworks[0:4:4]                                            // length 4, capacity 4
	webFrameworks = append(webFrameworks, "Express", "Meteor", "Ember", "Backbone") // not possible with arrays

	fmt.Println("\nWeb Frameworks: ", webFrameworks)
	fmt.Println("JS Frameworks: ", jsFrameworks)
}
