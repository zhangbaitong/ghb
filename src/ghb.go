package main

import (
	"fmt"
	"testing"
)

func main() {
	fmt.Println("11111")
}
func BenchmarkHTTP(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello")
	}

}
