package wc

import (
	"testing"
)

func TestWrite(t *testing.T) {
	var wc wordCounter
	var tests = []struct {
		s    []byte
		want int
	}{
		{[]byte("Hello world, wordCounter"), 3},
		{[]byte("a b c d e f g h i j k l m n o p q r s t u v w x y z."), 26},
	}
	for _, test := range tests {
		if ans, _ := wc.Write(test.s); ans != test.want {
			t.Errorf("Counting failed: get %d, want %d", ans, test.want)
		}
	}
}

// TODO: Let wc read the input from os.Stdin
/*
func TestWriteFromStdin(t *testing.T) {
	var wc wordCounter
	// new reader from os.Stdin, which is []byte
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		t.Errorf("Fail to read from Stdin\n")
	}
	wc.Write(stdin)
	t.Logf("wc: %d", wc)
}
*/
