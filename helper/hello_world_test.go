package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorldSiroj(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Siroj")
	}
}

func BenchmarkHelloWorldIzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Izz Luthfi")
	}
}

func BenchmarkHelloWorldSubBenchmark(b *testing.B) {
	b.Run("IzzBenchmark", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Izz")
		}
	})
	b.Run("LuthfiBenchmark", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Luthfi")
		}
	})
}

func BenchmarkHelloWorldTableBenchmark(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{name: "SitiBenchmarkTable", request: "Siti"},
		{name: "RohimatulBenchmarkTable", request: "Rohimatul"},
		{name: "LatifahBenchmarkTable", request: "Latifah"},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
}

func TestHelloSirojAssert(t *testing.T) {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "darwin" {
		t.Skip("Skip unit test on Mac OS")
	}

	result := HelloWorld("Siroj")
	require.Equal(t, "Hello Siroj", result, "Result must be 'Hello Siroj'")
}

func TestHelloSiroj(t *testing.T) {
	result := HelloWorld("Siroj")
	if result != "Hello Siroj" {
		// error
		t.Fatal("Result is not 'Hello Siroj'")
	}
	fmt.Println("TestHelloSiroj done")
}

func TestSubTestHelloWorld(t *testing.T) {
	t.Run("Muhamad", func(t *testing.T) {
		result := HelloWorld("Muhamad")
		require.Equal(t, "Hello Muhamad", result, "Result must be 'Hello Muhamad'")
	})
	t.Run("Izz", func(t *testing.T) {
		result := HelloWorld("Izz")
		require.Equal(t, "Hello Izz", result, "Result must be 'Hello Izz'")
	})
}

func TestTableTestHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Luthfi)",
			request:  "Luthfi",
			expected: "Hello Luthfi",
		},
		{
			name:     "HelloWorld(El Shirazy)",
			request:  "El Shirazy",
			expected: "Hello El Shirazy",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
