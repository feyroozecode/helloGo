package main

import "testing"

func TestSayHello(t *testing.T){
  got  := SayHello()
  want := "Salam" 
  
  if got != want {
    t.Error("got %q want %q, got, other thing ")
  } 
}
