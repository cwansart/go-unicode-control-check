# Unicode control character check with Go

This is a test project to demonstrate how to check for control characters in Unicode strings. This can be handy to check input data from an untrusted source.

## Introduction

The "go-unicode-control-check" package is a tool that allows Go developers to detect control characters within Unicode strings. Control characters are characters that do not represent a printable symbol, but instead serve to control the display or processing of text. This package can be useful for preventing security vulnerabilities, as control characters can be used for malicious purposes such as injection attacks.

## Installation

To install the package, run the following command:

```
go get github.com/cwansart/go-unicode-control-check
```

## Usage

To use the package, import it into your Go code:

```go
import "github.com/cwansart/go-unicode-control-check"
```

The package provides a single function, `HasControlCharacter`, which takes a string as input and returns a boolean indicating whether the string contains any control characters. Here is an example usage:

```go
package main

import (
	"fmt"

	"github.com/cwansart/go-unicode-control-check"
)

func main() {
	hasControl := go_unicode_control_check.HasControlCharacter("Hello, world!")
	if hasControl {
		fmt.Println("String contains control characters.")
	} else {
		fmt.Println("String does not contain control characters.")
	}
}
```

## Contributing

If you find a bug or have an idea for a new feature, please open an issue on the package's GitHub repository: https://github.com/cwansart/go-unicode-control-check/issues

Pull requests are also welcome! To contribute code changes, please fork the repository, make your changes, and submit a pull request. 

## License

The "go-unicode-control-check" package is released under the MIT License. See the LICENSE file in the package's GitHub repository for more information.
