package test

import (
	"fmt"
	"testing"
)

//go test -v unit_test.go
func Test_Loop(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Sprintf("Hello")
	}
}

//go test -v -bench=".*"
//生成CPU状态图
//go test -bench=".*" -cpuprofile=cpu.prof -c
//go tool pprof test.test cpu.prof
func Benchmark_Loop(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("Hello")
	}

}
