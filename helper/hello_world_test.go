package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

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

func TestSubTest(t *testing.T) {
	t.Run("Muhamad", func(t *testing.T) {
		result := HelloWorld("Muhamad")
		require.Equal(t, "Hello Muhamad", result, "Result must be 'Hello Muhamad'")
	})
	t.Run("Izz", func(t *testing.T) {
		result := HelloWorld("Izz")
		require.Equal(t, "Hello Izz", result, "Result must be 'Hello Izz'")
	})
}
