package main

import "fmt"

func nonempty(strings []string) []string {
    i := 0
    for _, data := range strings {
        if data != "" {
            strings[i] = data
            i++
        }
    }
    return strings[:i]
}
func nonempty2(strings []string) []string {
    out := strings[:0]
    for _, data := range strings {
        if data != "" {
            out = append(out, data)
        }
    }
    return out
}
func main() {
    s := []string{"hello", "", "world"}
    s = nonempty(s)
    //s = nonempty2(s)
    fmt.Println(s)

}
