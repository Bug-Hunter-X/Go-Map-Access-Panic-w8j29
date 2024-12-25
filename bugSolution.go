```go
package main

import "fmt"

func main() {
    var m map[string]int
    // Check if the map is nil before accessing an element
    if m == nil {
        m = make(map[string]int)
    }
    m["key"] = 10
    fmt.Println(m["key"])
}
```