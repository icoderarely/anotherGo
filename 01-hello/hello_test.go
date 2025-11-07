package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Joseph"},
			result: "Hello, Joseph!",
		},
		{
			items:  []string{"Joseph", "Matt"},
			result: "Hello, Joseph, Matt!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("want %s '%v', got %s", st.result, st.items, s)
		}
	}

	// want := "Hello, test!"
	// got := Say([]string{"test"})
	//
	// if want != got {
	// 	t.Errorf("wanted %s, got %s", want, got)
	// }
}
