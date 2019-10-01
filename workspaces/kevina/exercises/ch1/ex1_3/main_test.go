package main

import "testing"

var argsList = []string{
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
	"1",
}

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1(argsList)
	}
}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2([]string{"dsfsdf", "sdfsdfsd", "dsfdsf", "sdfsdfd"})
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3([]string{"dsfsdf", "sdfsdfsd", "dsfdsf", "sdfsdfd"})
	}
}
