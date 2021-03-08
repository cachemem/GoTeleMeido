package reverse

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

const wordSize int = 100
const maxSymbol int = 16384

type ReverseTest struct {
	v  string
	rv string
}

func (r *ReverseTest) Generate(rand *rand.Rand, size int) reflect.Value {
	value := make([]rune, wordSize)
	reversedValue := make([]rune, wordSize)

	for x := 0; x < wordSize; x++ {
		// Symbols from 0 to 16384
		new_char := rune(rand.Intn(maxSymbol + 1))
		value[x] = new_char
		reversedValue[wordSize-x-1] = new_char
	}
	return reflect.ValueOf(&ReverseTest{v: string(value), rv: string(reversedValue)})
}

func TestReverse(t *testing.T) {
	f := func(test string) bool {
		first := Reverse(test)
		return Reverse(first) == test
	}
	if err := quick.Check(f, &quick.Config{MaxCount: 100000}); err != nil {
		t.Error(err)
	}
	g := func(test *ReverseTest) bool {
		return Reverse(test.v) == test.rv
	}
	if err := quick.Check(g, &quick.Config{MaxCount: 100}); err != nil {
		t.Error(err)
	}
	tableTests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"abc", "cba"},
		{"русское поле экспериментов", "вотнемирепскэ елоп еокссур"},
		{"doença", "açneod"},
		{"a", "a"},
	}
	for _, test := range tableTests {
		if got := Reverse(test.input); got != test.want {
			t.Errorf("Reverse incorrectly reversed %q, got %q", test.input, got)
		}
	}
}
