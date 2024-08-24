package main

import "fmt"

func PrintDeployOptions2() {
	deployOptions := [...]string{"R-pi", "AWS", "GCP", "Azure", "IBM"}

	for index, option := range deployOptions {
		fmt.Println(index, option)
	}

}

func main() {
	PrintDeployOptions2()
}
