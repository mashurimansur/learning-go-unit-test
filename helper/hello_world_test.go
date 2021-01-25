package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Hury",
			request: "Hury",
		},
		{
			name:    "Mashuri",
			request: "Mashuri",
		},
		{
			name:    "Jumardi",
			request: "Jumardi",
		},
		{
			name:    "Budi",
			request: "Budi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Hury", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hury")
		}
	})
	b.Run("Mashuri", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Mashuri")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hury")
	}
}

func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Mashuri")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hury",
			request:  "Hury",
			expected: "Hello Hury",
		},
		{
			name:     "Mashuri",
			request:  "Mashuri",
			expected: "Hello Mashuri",
		},
		{
			name:     "Mansur",
			request:  "Mansur",
			expected: "Hello Mansur",
		},
		{
			name:     "Jumardi",
			request:  "Jumardi",
			expected: "Hello Jumardi",
		},
		{
			name:     "Muna",
			request:  "Muna",
			expected: "Hello Muna",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Hury", func(t *testing.T) {
		result := HelloWorld("Hury")
		require.Equal(t, "Hello Hury", result, "Result must be 'Hello Hury'")
	})
	t.Run("Mashuri", func(t *testing.T) {
		result := HelloWorld("Mashuri")
		require.Equal(t, "Hello Mashuri", result, "Result must be 'Hello Mashuri'")
	})
	t.Run("Mansur", func(t *testing.T) {
		result := HelloWorld("Mansur")
		require.Equal(t, "Hello Mansur", result, "Result must be 'Hello Mansur'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Hury")
	require.Equal(t, "Hello Hury", result, "Result must be 'Hello Hury'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Hury")
	require.Equal(t, "Hello Hury", result, "Result must be 'Hello Hury'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Hury")
	assert.Equal(t, "Hello Hury", result, "Result must be 'Hello Hury'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldHury(t *testing.T) {
	result := HelloWorld("Hury")

	if result != "Hello Hury" {
		// error
		t.Error("Result must be 'Hello Hury'")
	}

	fmt.Println("TestHelloWorldHury Done")
}

func TestHelloWorldMansur(t *testing.T) {
	result := HelloWorld("Mansur")

	if result != "Hello Mansur" {
		// error
		t.Fatal("Result must be 'Hello Mansur'")
	}

	fmt.Println("TestHelloWorldMansur Done")
}

func TestHelloWorldJumardi(t *testing.T) {
	result := HelloWorld("Jumardi")

	if result != "Hello Jumardi" {
		// error
		t.Fatal("Result must be 'Hello Jumardi'")
	}

	fmt.Println("TestHelloWorldJumardi Done")
}
