# Go Map Access Panic

This repository demonstrates a common error in Go: panicking when accessing a nil map.  The `bug.go` file shows the problematic code, while `bugSolution.go` provides a solution using the idiomatic Go approach for safe map access.

## Problem

Accessing a map element in Go without checking if the map is nil will result in a runtime panic if the map hasn't been initialized.  This can be hard to debug, especially in complex programs.

## Solution

Always check for `nil` before accessing elements in a map.
