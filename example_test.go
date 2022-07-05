package workgroup_test

import (
	"fmt"
	"strings"

	"github.com/ksdfg/workgroup/v2"
)

func ExampleRun() {
	// Declare phrase and keywords to search in phrase
	phrase := "A small utility to manage the lifetime of a set of related goroutines."
	keywords := []string{"function", "variable", "slice", "goroutine", "package"}

	// Create slice of functions
	// Each function checks if a keyword is a substring present in the phrase
	// Return the keyword if it is found
	// Return nil if it is not found
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

	// Call Run to search for keywords in parallel
	output := workgroup.Run(fns, 3)
	fmt.Println(output, "found")
	// Output:
	// goroutine found
}

func ExampleRunTemplate() {
	// Declare phrase and keywords to search in phrase
	phrase := "A small utility to manage the lifetime of a set of related goroutines."
	keywords := []string{"function", "variable", "slice", "goroutine", "package"}

	// Call RunTemplate to search for keywords in parallel
	// Pass number of keywords as n
	// In the template function, check if the ith keyword is a substring in the phrase
	// Return the keyword if it is found
	// Return nil if it is not found
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
	fmt.Println(output, "found")
	// Output:
	// goroutine found
}
