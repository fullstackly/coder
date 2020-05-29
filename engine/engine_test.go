package engine

import "testing"

func TestEngine(t *testing.T) {
	t.Run("encode", func(t *testing.T) {
		for _, testCase := range engineTests {
			got := Encode(testCase.initial)
			assertCoding(t, got, testCase.encoded)
		}
	})

	t.Run("decode", func(t *testing.T) {
		for _, testCase := range engineTests {
			got := Decode(testCase.encoded)
			assertCoding(t, got, testCase.decoded)
		}
	})
}

var engineTests = []struct {
	initial string
	encoded string
	decoded string
}{
	{initial: "hello", encoded: "h2ll4", decoded: "hello"},
	{initial: "hi there", encoded: "h3 th2r2", decoded: "hi there"},
	{initial: "YOU", encoded: "Y45", decoded: "You"},
	{initial: "Ola", encoded: "4l1", decoded: "ola"},
}

func assertCoding(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
