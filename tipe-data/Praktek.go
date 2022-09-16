package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x int
	var y float64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ = strconv.ParseFloat(scanner.Text(), 32)

	//Tulis kode disini
	Result1 := x + int(y)
	Result2 := float64(x) - y
	Result3 := x * int(y)

	fmt.Printf("%d\n", Result1)
	fmt.Printf("%.2f \n", Result2)
	fmt.Printf("%d", Result3)
}
