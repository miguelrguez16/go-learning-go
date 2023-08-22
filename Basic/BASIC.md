# BASIC GO

- [BASIC GO](#basic-go)
  - [CHARACTERISTICS OF GO](#characteristics-of-go)
  - [Start](#start)
  - [Run Basic Go](#run-basic-go)
  - [DATA TYPES](#data-types)
  - [Maps](#maps)
  - [Logic](#logic)
  - [Other](#other)

## CHARACTERISTICS OF GO

- Simple syntax
- fast build
- start up
- fewer resources
- compiled

## Start

1. Create a Module
   2. Module path can correspond to a repository you can plan to publish your module to for example, Github
`$ go mod init <module_path>`

2. Initialized a go.mod file
   3. Describes the module: with name/module path

3. All code must belong to a package
4. First line of a file must be 'package ...'
5. EntryPoint

```Go
package main

func main(){
}
```

## Run Basic Go

```go
go run .\Basic.go

go run .\main.go

// RUN ALL OF A PACKAGE
go run .
```

## DATA TYPES

- Strings
- int / uint
|  GO   | Java  |
|:-----:|-------|
| int8  | byte  |
| int16 | short |
| int36 | int   |
| int64 | long  |
- float
- Maps
- Arrays <> Slice

## Maps

```go
// create a map
var myMap map[Key]Value
var myMap map[string]string

// create and empty map
// func make --> builtin function
var userData = make(map[string]string) 
```

## Logic

```go
for {}

```

## Other

```go
_ //-> to ignore a variable you do not want to use
```
