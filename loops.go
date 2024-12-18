package main

import "fmt"

func main(){
  // simple for 
  for i := 0; i< 5 ; i++ {
     fmt.Println(i)
  }

  // for loops as while
  sum := 2
  for sum < 15 {
   sum += sum
   fmt.Println("new some = ", sum)
  }
  
  // range based loops
  fruits := []string{"Apple", "Fraise", "Orange"}
  for index, value := range fruits { 
   fmt.Printf("The index %d, content=  %s\n ", index, value)
  }
}
