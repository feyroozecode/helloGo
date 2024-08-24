package main

import "fmt"

func main() {
	languages := [9]string{"Python", "Java", "Go", "C", "C++", "C#", "JavaScript", "Ruby", "Rust"}

	classes := languages[3:6]
	modern := make([]string, 4)
	modern = languages[2:6]
	new := languages[7:9]

	fmt.Println("Classes: \n", classes)
	fmt.Println("Modern: \n", modern)
	fmt.Println("New: \n", new)

}
