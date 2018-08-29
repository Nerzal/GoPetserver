# GoPetserver
Repository to learn golang using a petserver

# GO Basics

## General
 - Composition over inheritance
 	- There is no inheritance
 - *There are no Classes!*
 - No Method Overloading
 	- Each Method is unique
 - Go erzwingt Egypt-Style Klammern
 - https://golang.org
 - IDE VSCode / Goland

### Keywords
**const**
```go
const {
	PI = 3 //Yes i did!
	Language = "Go"
}
```
 
## Commands
 - go run <path>
 	- runs the program
 - go build <path>
 	- builds the program
 - go install <path>
 	- copies an executable to the bin path

## Visibility
   - lowercase = Only visible inside same package
   - uppercase = Visible outside package

## Packages
 - Similar to namespaces
 - Import
 	- Imports Packages
 	 - Local
 		- Remote (github repo)
 		- SourceCode
   
## Types

### Basic
 - int
 - bool
 - string
 - uint, uinptr
 - float32, float64
 - complex64, complex128
 - byte (uint8)
 - rune (char)

### Advanced (complex)
 - Array
 	- numbered collection
 - Slice
 	- array that can grow in size (like a list)
 - Struct
 - Pointer
 - Function
 - Interface
 	- Not implemented explicitly
 - Map
 	- Like a dictionary
 - Channel
 	- Used for communication between goroutines

# Go Functions
 - Go functions can have multiple return values.
 - Can be used like a type.
 - Support function literals
 	- You can write functions inside another function.
 	- Runs in context of parent function

```go
func MakeNoise(noise string){
	fmt.PrintLn(noise);
}
```

# Useful links
## Go Language Spec 
https://golang.org/ref/spec