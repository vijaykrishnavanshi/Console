# Console

Coming from node I was missing console and its simplicity.
This is just a fun package and implements most basic functionality of console.


```closure
// Usage of console package
package main

import (
        "console"
)

func main() {
        console.Log("Hello ", "World !!")
        console.Log(1, "World !!")
        m := make(map[string]int)
        m["Hello"] = 1
        m["Helloa"] = 1
        m["Hellob"] = 1
        m["Helloc"] = 1
        console.Log(m, "World !!")
}

```

PS: This is my first go package

