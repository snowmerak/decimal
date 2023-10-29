# decimal

`decimal` is a simple Fraction type library for go.

## Install

```bash
go get -u github.com/snowmerak/decimal
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/snowmerak/decimal"
)

func main() {
	a := decimal.New64(decimal.Positive, 300, 7)

	fmt.Println(a) // 300/7

	b := decimal.New64(decimal.Negative, 15, 2)

	fmt.Println(b) // -15/2

	fmt.Println(a.Add(b)) // 300/7 + -15/2 = 495/14

	fmt.Println(a.Sub(b)) // 300/7 - -15/2 = -705/14

	fmt.Println(a.Mul(b)) // 300/7 * -15/2 = -2250/7

	fmt.Println(a.Div(b)) // 300/7 / -15/2 = -40/7

	f32 := a.Float32()
	f64 := a.Float64()

	fmt.Println(f32) // 42.857143
	fmt.Println(f64) // 42.857142857142854

	fmt.Println(a.Float64() + b.Float64()) // 42.857142857142854 + -7.5 = 35.357142857142854
	fmt.Println(a.Add(b).Float64())        // 495/14 = 35.357142857142854
}
```