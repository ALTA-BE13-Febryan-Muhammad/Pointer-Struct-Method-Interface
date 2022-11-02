package main

import "fmt"

func Swap(a, b *int) (int, int) {
	//your code here

	//pendek
	// var wadahA, wadahB int = *a, *b
	// *a = wadahB
	// *b = wadahA
	// return wadahB, wadahA

	//panjang
	// wadahA := *a
	// wadahB := *b
	// *a = wadahB
	// *b = wadahA
	// return wadahA, wadahB

	//very short
	*a, *b = *b, *a
	return *a, *b

}

func main() {
	a := 10
	b := 20

	Swap(&a, &b)
	fmt.Println(a, b)
}

//status * itu di negate dengan &
