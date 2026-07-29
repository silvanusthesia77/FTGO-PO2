package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// learnVariable()
	// learnTypeData()
	// constAndOperator()
	learnArray()
}

func learnVariable() {
	name := "arthur"
	age := 22
	address := "kilo 17"
	address = "malibela"

	a, b, c := "junior", 12, "dream"
	dataA, dataB, dataC := "dataA", "dataB", "dataC"
	_, _, _ = dataA, dataB, dataC
	_ = "wanus thesia"
	fmt.Println(a, b, c)
	fmt.Printf("nama saya %s , saya berusia %d dan saya tinggal di %s", name, age, address)
	spew.Dump("Nama", dataA)
}
func learnTypeData() {
	// type data number
	number := 0.12
	spew.Dump(number)
	//  type data string
	thisIsString := "thby"
	spew.Dump(thisIsString)

	thisIsBool := true
	spew.Dump(thisIsBool)
}
func constAndOperator() {
	const name string = "Aplikasi Baju"
	spew.Dump(name)

	opratorA := 2 + 2
	opratorB := 2 * 2
	opratorC := 2 / 8
	opratorD := 2 % 5

	spew.Dump("A : %d , B : %d , C : %d , D : %d ,\n", opratorA, opratorB, opratorC, opratorD)
	//  Perbandaingan
	kl := 30
	kc := 22

	hasil := kl == kc

	fmt.Println(hasil)
	fmt.Println(kl != kc)
	h := true
	j := false

	fmt.Println(h && j)
	fmt.Println(h || j)
}
func learnArray() {
	number := [3]int{22, 44, 66}
	num := []int{100, 300, 500, 700}
	arry := [5]string{}
	arry[0] = "wanus"
	arry[1] = "thby"
	arry[2] = "reza"
	arry[3] = "lukas"
	for _, v := range arry {
		spew.Dump(v)
	}
	spew.Dump(num)
	spew.Dump(number)
}

// 1:15 jam
