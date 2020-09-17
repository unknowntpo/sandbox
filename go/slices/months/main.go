// Ch4.2 - Slices
// format of comments in go?

package main

import "fmt"

func main() {
	months := [...]string{0: "",
		1: "Jan", 2: "Feb", 3: "Mar",
		4: "Apr", 5: "May", 6: "Jun",
		7: "Jul", 8: "Aug", 9: "Sep",
		10: "Oct", 11: "Nov", 12: "Dec"}
	Q2 := months[4:7]
        Q1 := months[1:4]
	summer := months[6:9]

        for _, s := range summer {
            for _, q := range Q2 {
                if s == q {
                    fmt.Printf("%q appears in both\n", s) // if use %q: "Jun" appears in both
                }
            }
        }

        fmt.Println("Q2:\t", Q2, "length:\t", len(Q2), "capacity:\t", cap(Q2))
        fmt.Println("Q1:\t", Q1, "length:\t", len(Q1), "capacity:\t", cap(Q1))
        fmt.Println("")
        fmt.Println("summer:\t", summer)
}
