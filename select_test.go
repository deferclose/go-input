package input

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestSelect(t *testing.T) {
	cases := []struct {
		list      []string
		opts      *Options
		userInput io.Reader
		expect    string
	}{
		{
			list:      []string{"A", "B", "C"},
			opts:      &Options{},
			userInput: bytes.NewBufferString("1\n"),
			expect:    "A",
		},

		{
			list: []string{"A", "B", "C"},
			opts: &Options{
				Default: "A",
			},
			userInput: bytes.NewBufferString("\n"),
			expect:    "A",
		},

		{
			list: []string{"A", "B", "C"},
			opts: &Options{
				Default: "A",
			},
			userInput: bytes.NewBufferString("3\n"),
			expect:    "C",
		},

		// Loop
		{
			list: []string{"A", "B", "C"},
			opts: &Options{
				Loop: true,
			},
			userInput: bytes.NewBufferString("\n3\n"),
			expect:    "C",
		},

		// Loop
		{
			list: []string{"A", "B", "C"},
			opts: &Options{
				Loop: true,
			},
			userInput: bytes.NewBufferString("\n\n\n\n\n2\n"),
			expect:    "B",
		},

		// Loop
		{
			list: []string{"A", "B", "C"},
			opts: &Options{
				Loop: true,
			},
			userInput: bytes.NewBufferString("4\n3\n"),
			expect:    "C",
		},

		// Loop
		{
			list: []string{"A", "B", "C"},
			opts: &Options{
				Loop: true,
			},
			userInput: bytes.NewBufferString("A\n3\n"),
			expect:    "C",
		},
	}

	for i, c := range cases {
		ui := &UI{
			Writer: io.Discard,
			Reader: c.userInput,
		}

		ans, _, err := ui.Select("", c.list, c.opts)
		if err != nil {
			t.Fatalf("#%d expect not to occurr error: %s", i, err)
		}

		if ans != c.expect {
			t.Fatalf("#%d expect %q to be eq %q", i, ans, c.expect)
		}
	}
}

func TestSelect_invalidDefault(t *testing.T) {
	ui := &UI{
		Writer: io.Discard,
	}
	_, _, err := ui.Select("Which?", []string{"A", "B", "C"}, &Options{
		// "D" is not in select target list
		Default: "D",
	})

	if err == nil {
		t.Fatal("expect err to be occurr")
	}
}

func TestSelect_SelectDefault(t *testing.T) {
	ui := &UI{
		Reader: bytes.NewBufferString("\r"),
		Writer: io.Discard,
	}
	rslt, n, err := ui.Select("Which?", []string{"A", "B", "C"}, &Options{
		// "D" is not in select target list
		Default: "A",
	})

	if err != nil {
		t.Fatal("expect err to be nil, but got", err)
	}

	if rslt != "A" {
		t.Fatal("expect rslt to be A")
	}
	if n != 0 {
		t.Fatal("expect n to be 0")
	}
}

func TestSelect_SelectDefault2(t *testing.T) {
	ui := &UI{
		Reader: bytes.NewBufferString("\r"),
		Writer: io.Discard,
	}
	rslt, n, err := ui.Select("Which?", []string{"A", "B", "C"}, &Options{
		// "D" is not in select target list
		DefaultSelected: 2,
	})

	if err != nil {
		t.Fatal("expect err to be nil, but got", err)
	}

	if rslt != "B" {
		t.Fatal("expect rslt to be B")
	}
	if n != 1 {
		t.Fatal("expect n to be 1")
	}
}

func ExampleUI_Select() {
	ui := &UI{
		// In real world, Reader is os.Stdin and input comes
		// from user actual input.
		Reader: bytes.NewBufferString("3\n"),
		Writer: io.Discard,
	}

	query := "Which language do you prefer to use?"
	lang, n, _ := ui.Select(query, []string{"go", "Go", "golang"}, &Options{
		Default: "Go",
	})

	fmt.Println(lang, n)
	// Output: golang 2
}

func ExampleUI_Select_2() {
	ui := &UI{
		// In real world, Reader is os.Stdin and input comes
		// from user actual input.
		Reader: bytes.NewBufferString("3\r"),
		Writer: io.Discard,
	}

	query := "Which language do you prefer to use?"
	lang, n, _ := ui.Select(query, []string{"go", "Go", "golang"}, &Options{
		Default: "Go",
	})

	fmt.Println(lang, n)
	// Output: golang 2
}
