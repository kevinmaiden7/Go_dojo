package main

import "fmt"
import "math"
import "rsc.io/quote"

func main(){
	fmt.Println("Greetings from Go!")
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())
	fmt.Println("---")

	const base = 2
	fmt.Println("Powers of", base)
	for i := 0; i < 13; i++	{
		fmt.Printf("%.1f", math.Pow(float64(base), float64(i)))
		fmt.Println("")
	}
}
