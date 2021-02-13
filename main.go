package main

import "fmt"

const x string = "fixed!"

func main() {

	var a = "x-team"
	fmt.Println(a)

	var b, c, d int = 1, 2, 3
	fmt.Println(b, c, d)

	var e = true
	fmt.Println(e)

	var f int
	fmt.Println(f)
	fmt.Printf("value: %v type: %T\n", f, f)

	var g string
	fmt.Println(g)
	fmt.Printf("value: %v type: %T\n", g, g)

	h := "go bootcamp"
	fmt.Println(h)

	fmt.Println(x)

	const w = "yeah"
	fmt.Println(w)

	var y float64 = 666.70
	fmt.Println(y)

	z := int(y)
	fmt.Println(z)

	msg := fmt.Sprintf("%d %d %d %v %v %v", b, c, d, a, h, w)
	fmt.Println(msg)
}
