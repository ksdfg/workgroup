# Workgroup

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/ksdfg/workgroup/Test?style=flat)
[![Go Reference](https://pkg.go.dev/badge/github.com/ksdfg/workgroup.svg)](https://pkg.go.dev/github.com/ksdfg/workgroup)

A small utility to manage the lifetime of a set of related goroutines.

## Installation

```shell
$ go get -v github.com/ksdfg/workgroup
```

## Functions

### Run

```go
func Run(fns []func () interface{}) interface{}
```

- Run will execute all functions in the slice passed to it in individual goroutines.
- Run blocks until all the goroutines spawned have ended execution.
- The first function to return a non-nil value will trigger the end of execution of all other goroutines spawned.
- The return value from the first function will be returned to the caller of Run.
- If all none of the functions return a non-nil value, all the spawned goroutines will naturally end execution and Run
  will return nil.

### RunTemplate

```go
func RunTemplate(n int, template func (int) interface{}) interface{}
```

- RunTemplate will execute a given template function n no. of times in individual goroutines.
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

	"github.com/ksdfg/workgroup"
)

// Search for substrings parallelly
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

	// Call Run to search for keywords parallelly
	output := workgroup.Run(fns)
	fmt.Println(output, "found")
}
```

### RunTemplate

```go
package main

import (
	"fmt"
	"strings"

	"github.com/ksdfg/workgroup"
)

// Search for substrings parallelly
func main() {
	// Declare phrase and keywords to search in phrase
	phrase := "A small utility to manage the lifetime of a set of related goroutines."
	keywords := []string{"function", "variable", "slice", "goroutine", "package"}

	// Call RunTemplate to search for keywords parallelly
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