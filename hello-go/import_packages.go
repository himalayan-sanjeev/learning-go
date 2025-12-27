package main

// import "fmt"
// import "rsc.io/quote"

// Multiple imports
import (
	"fmt"

	"rsc.io/quote"
)

func printQuote() {
	fmt.Println(quote.Go())
}
