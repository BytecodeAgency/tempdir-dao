# Temporary Directory Data-Access Object [![Build Status](https://travis-ci.com/BytecodeAgency/tempdir-dao.svg?branch=master)](https://travis-ci.com/BytecodeAgency/tempdir-dao)

_Warning: this library has been developed for R&D purposes for an internal project at Bytecode. This package is not tested for production usage and it is not (yet) feature complete. For example, using sub-directories is not supported (yet)_

## Why don't you use `ioutil.TempFile` and/or `os.TempDir`

This package was created specifically to work with multiple files in a grouped together. This rules out using `iouril.TempFile` as it only works with single files.

Using `os.TempDir` was ruled out as it [says in the documentation](https://golang.org/pkg/os/#TempDir) that `The directory is neither guaranteed to exist nor have accessible permissions`, while this directory does not delete the files/directory until that method has been called explicitly.



## Installation

Use `go get` or your Go package manager choice:

```
go get -u github.com/BytecodeAgency/tempdir-dao
```

## Usage

To create a new DAO, use 

```go
tempDirDAO, err := tempdirdao.NewTempFileAccess() // returns `(*TempFileAccess, error)`
```

The following methods are available:

| Method | Functionality | Arguments | Returns |
| ------ | ------------- | --------- | ------- |
| `LoadDirContents` | Lists contents of temp directory | `()` | `([]os.FileInfo, error`
| `GetFullFilePath` | Gets the full file path of a file in the temp dir | `(filename string)`  | `(string, error)`
| `LoadFileContents` | Loads contents of a file in the temp dir | `()` | `([]byte, error)`
| `SaveFile` | Created or overwrites a file in the temp dir | `(filename string, fileContents []byte, permissions os.FileMode)` | `error`
| `RemoveTempFileAccess` | Removes the temp dir | `()` | `error`
 

_Note: you never need to include the temporary directory name in the arguments for methods_

## Example

```go
package main

import (
    "github.com/BytecodeAgency/tempdir-dao"
    "log"
)

const filename = "somefile.ext"

func main () {
    fileContents := funcThatReturnsString()
    
    tfa, err := tempdirdao.NewTempFileAccess()
    defer tfa.RemoveTempFileAccess()
    if err != nil {
    	log.Fatal(err)
    }
    err = tfa.SaveFile(filename, fileContents, 0644)
    if err != nil {
        log.Fatal(err)
    }
    
    fullFilePath, err := tfa.GetFullFilePath(filename)
    if err != nil {
        log.Fatal(err)
    }
    funcThatUsesTheTempFile(fullFilePath)
}
```
