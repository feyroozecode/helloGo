package main

import "fmt"

func PrintDeployOptions() {

	deployOptions := [4]string{"R-pi", "AWS", "GCP", "Azure"}

	for i := 0; i <= len(deployOptions); i++ {
		option := deployOptions[i]
		fmt.Println(i, option)
	}

}

func main() {
	PrintDeployOptions()
}
