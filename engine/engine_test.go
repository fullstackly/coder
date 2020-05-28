package engine

import "testing"

func TestEngine(t *testing.T) {

	engineTests := []struct {
		initial string
		encoded string
	}{
		{initial: "hello", encoded: "h2ll4"},
		{initial: "hi there", encoded: "h3 th2r2"},
		{initial: "YOU", encoded: "Y45"},
		{initial: "Ola", encoded: "4l1"},
	}

	assertCoding := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("encode", func(t *testing.T) {
		for _, testCase := range engineTests {
			got := Encode(testCase.initial)
			assertCoding(t, got, testCase.encoded)
		}
	})

	t.Run("decode", func(t *testing.T) {
		for _, testCase := range engineTests {
			got := Decode(testCase.encoded)
			assertCoding(t, got, testCase.initial)
		}
	})
}
