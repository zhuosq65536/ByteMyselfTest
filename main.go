package main

import (
	"fmt"
  "math/rand"

)

func main() {
  maxNum := 100
  secretNumber := rand.Intn(maxNum);
  fmt.Println("The secretNumber is ",secretNumber)

	fmt.Println("Hello, World!")
}

