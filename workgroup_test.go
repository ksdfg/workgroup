package workgroup_test

import (
	"strings"
	"testing"

	"github.com/ksdfg/workgroup/v2"
)

/*
Tests for Run
*/

func TestRunEmptySlice(t *testing.T) {
	var slice []func() interface{}
	output := workgroup.Run(slice, 3)
	if output != nil {
		t.Fatalf("expected: %v\ngot: %v", nil, output)
	}
}

func TestRunSearchSubstringSuccess(t *testing.T) {
	phrase := "Neko-chan the cat goes meow."
	keywords := []string{"dog", "camel", "horse", "wolf", "fox", "tiger", "cat"}

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

	output := workgroup.Run(fns, 3)
	if output != "cat" {
		t.Fatalf("expected: cat\ngot: %v", output)
	}
}

func TestRunSearchSubstringFailure(t *testing.T) {
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

	output := workgroup.Run(fns, 3)
	if output != nil {
		t.Fatalf("expected: %v\ngot: %v", nil, output)
	}
}

func TestRunSearchSubstringSuccessMany(t *testing.T) {
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

	output := workgroup.Run(fns, 3)
	if output != "cat" && output != "meow" {
		t.Fatalf("expected: 'cat' or 'meow'\ngot: %v", output)
	}
}

/*
Tests for RunTemplate
*/

func TestRunTemplateEmptySlice(t *testing.T) {
	output := workgroup.RunTemplate(
		0,
		func(_ int) interface{} {
			return nil
		},
		3,
	)
	if output != nil {
		t.Fatalf("expected: %v\ngot: %v", nil, output)
	}
}

func TestRunTemplateSearchSubstringSuccess(t *testing.T) {
	phrase := "Neko-chan the cat goes meow."
	keywords := []string{"dog", "camel", "horse", "cat", "wolf", "fox", "tiger"}

	output := workgroup.RunTemplate(
		len(keywords),
		func(index int) interface{} {
			var k interface{}
			if strings.Contains(phrase, keywords[index]) {
				k = keywords[index]
			}
			return k
		},
		3,
	)
	if output != "cat" {
		t.Fatalf("expected: cat\ngot: %v", output)
	}
}

func TestRunTemplateSearchSubstringFailure(t *testing.T) {
	phrase := "Neko-chan the cat goes meow."
	keywords := []string{"dog", "camel", "horse", "katto", "wolf", "fox", "tiger"}

	output := workgroup.RunTemplate(
		len(keywords),
		func(index int) interface{} {
			var k interface{}
			if strings.Contains(phrase, keywords[index]) {
				k = keywords[index]
			}
			return k
		},
		3,
	)
	if output != nil {
		t.Fatalf("expected: %v\ngot: %v", nil, output)
	}
}

func TestRunTemplateSearchSubstringSuccessMany(t *testing.T) {
	phrase := "Neko-chan the cat goes meow."
	keywords := []string{"dog", "camel", "horse", "cat", "wolf", "fox", "tiger", "meow"}

	output := workgroup.RunTemplate(
		len(keywords),
		func(index int) interface{} {
			var k interface{}
			if strings.Contains(phrase, keywords[index]) {
				k = keywords[index]
			}
			return k
		},
		3,
	)
	if output != "cat" && output != "meow" {
		t.Fatalf("expected: 'cat' or 'meow'\ngot: %v", output)
	}
}
