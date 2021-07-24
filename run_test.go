package workgroup

import (
	"strings"
	"testing"
)

func TestEmptySlice(t *testing.T) {
	var slice []func() interface{}
	output := Run(slice)
	if output != nil {
		t.Fatalf("expected: %v\ngot: %v", nil, output)
	}
}

func TestSearchSubstringSuccess(t *testing.T) {
	phrase := "Neko-chan the cat goes meow."
	keywords := []string{"dog", "camel", "horse", "cat", "wolf", "fox", "tiger"}

	var fns []func() interface{}
	for _, keyword := range keywords {
		kw := keyword
		fns = append(
			fns,
			func() interface{} {
				var k interface{}
				if strings.Contains(phrase, kw) {
					k = kw
				}
				return k
			},
		)
	}

	output := Run(fns)
	if output != "cat" {
		t.Fatalf("expected: cat\ngot: %v", output)
	}
}

func TestSearchSubstringFailure(t *testing.T) {
	phrase := "Neko-chan the cat goes meow."
	keywords := []string{"dog", "camel", "horse", "katto", "wolf", "fox", "tiger"}

	var fns []func() interface{}
	for _, keyword := range keywords {
		kw := keyword
		fns = append(
			fns,
			func() interface{} {
				var k interface{}
				if strings.Contains(phrase, kw) {
					k = kw
				}
				return k
			},
		)
	}

	output := Run(fns)
	if output != nil {
		t.Fatalf("expected: %v\ngot: %v", nil, output)
	}
}

func TestSearchSubstringSuccessMany(t *testing.T) {
	phrase := "Neko-chan the cat goes meow."
	keywords := []string{"dog", "camel", "horse", "cat", "wolf", "fox", "tiger", "meow"}

	var fns []func() interface{}
	for _, keyword := range keywords {
		kw := keyword
		fns = append(
			fns,
			func() interface{} {
				var k interface{}
				if strings.Contains(phrase, kw) {
					k = kw
				}
				return k
			},
		)
	}

	output := Run(fns)
	if output != "cat" && output != "meow" {
		t.Fatalf("expected: 'cat' or 'meow'\ngot: %v", output)
	}
}
