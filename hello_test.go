package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	subTests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello world!",
		},
		{
			result: "Hello scott!",
			items:  []string{"scott"},
		},
		{
			result: "Hello scott, sarah!",
			items:  []string{"scott", "sarah"},
		},
	}

	for _, st := range subTests {
		want := st.result
		got := Say(st.items)
		if want != got {
			t.Errorf("Wanted %s, got %s", want, got)
		}

	}
}
