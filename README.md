# Workgroup

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/ksdfg/workgroup)
[![Test](https://github.com/ksdfg/workgroup/actions/workflows/test.yml/badge.svg)](https://github.com/ksdfg/workgroup/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/ksdfg/workgroup.svg)](https://pkg.go.dev/github.com/ksdfg/workgroup)

A small utility to manage the lifetime of a set of related goroutines.

## Installation

```shell
$ go get -v github.com/ksdfg/workgroup/v2
```

## Functions

### Run

```go
func Run(fns []func () interface{}, maxParallelGoroutines int) interface{}
```

- Run will execute all functions in the slice passed to it in individual goroutines.
- Goroutines are created in batches, never exceeding max parallel goroutines value passed as input.
- Run blocks until all the goroutines spawned have ended execution.
- The first function to return a non-nil value will trigger the end of execution of all other goroutines spawned.
- The return value from the first function will be returned to the caller of Run.
- If all none of the functions return a non-nil value, all the spawned goroutines will naturally end execution and Run
  will return nil.

### RunTemplate

```go
func RunTemplate(n int, template func (int) interface{}, maxParallelGoroutines int) interface{}
```

- RunTemplate will execute a given template function n no. of times in individual goroutines.
- Goroutines are created in batches, never exceeding max parallel goroutines value passed as input.
- RunTemplate blocks until all the goroutines spawned have ended execution.
- Each goroutine will pass an index to the template function which it can use to execute accordingly.
- The first function to return a non-nil value will trigger the end of execution of all other goroutines spawned.
- The return value from the first function will be returned to the caller of RunTemplate.
- If all none of the functions return a non-nil value, all the spawned goroutines will naturally end execution and
  RunTemplate will return nil.

## Examples

### Run

```go
package main

import (
	"fmt"
	"strings"

	"github.com/ksdfg/workgroup/v2"
)

// Search for substrings in parallel
func main() {
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

	// Call Run to search for keywords in parallel, with a maximum of 3 goroutines running at a time
	output := workgroup.Run(fns, 3)
	fmt.Println(output, "found")
}
```

### RunTemplate

```go
package main

import (
	"fmt"
	"strings"
	
	"github.com/ksdfg/workgroup/v2"
)

// Search for substrings in parallel
func main() {
	// Declare phrase and keywords to search in phrase
	phrase := "A small utility to manage the lifetime of a set of related goroutines."
	keywords := []string{"function", "variable", "slice", "goroutine", "package"}

	// Call RunTemplate to search for keywords in parallel, with a maximum of 3 goroutines running at a time
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
}
```

## Credits

Built as a replacement with revamped internal working and extended functionality
for [heptio/workgroup](https://pkg.go.dev/github.com/heptio/workgroup)

~~because that became an internal library
of [contour](https://github.com/projectcontour/contour/tree/main/internal/workgroup), and most of the other libraries I
found were nowhere as simple as this, leaving me to make an alternative myself **(〒︿〒)**~~