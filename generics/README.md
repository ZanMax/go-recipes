## Generics


```go
package main

import "fmt"

type Number interface {
	float64 | int64 | int
}

func sumOfSome[T Number](input []T) T {
	var sum T
	for _, v := range input {
		sum += v
	}
	return sum
}

func main() {
	someLst := []int{1, 2, 3, 4, 5}
	r := sumOfSome(someLst)
	fmt.Println("Sum of some:", r)
}
```