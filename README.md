# SanitizeFilename

A handy utility function that will sanitize a given filename string based on the current OS detected by the runtime.

### Installation

```bash
go get github.com/jacoblockett/sanitizefilename
```

You can read the godoc [here](https://pkg.go.dev/github.com/jacoblockett/sanitizefilename) for detailed documentation.

### Quickstart

```go
package main

import "github.com/jacoblockett/sanitizefilename"

func main() {
	// Assuming a Windows environment
	filename := "<>:\"/\\|?*abc.txt" // "<>:"/\|?*abc.txt" without escape chars
	sanitized := sanitizefilename.Sanitize(filename)

	fmt.Println(sanitized) // "abc.txt"

	// Assuming a Linux/Unix environment
	filename := "/.."
	sanitized := sanitizefilename.Sanitize(filename)

	fmt.Println(sanitized) // ""
}
```
