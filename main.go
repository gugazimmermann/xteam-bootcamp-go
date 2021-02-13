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

	var (
		i = 12
		j = 3
		k = false
		l = "foo"
		m float64
		n float64 = 2.4
	)
	fmt.Printf("value: %v type: %T\n", i, i)
	fmt.Printf("value: %v type: %T\n", j, j)
	fmt.Printf("value: %v type: %T\n", k, k)
	fmt.Printf("value: %v type: %T\n", l, l)
	fmt.Printf("value: %v type: %T\n", m, m)
	fmt.Printf("value: %v type: %T\n", n, n)

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
