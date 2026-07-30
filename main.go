package main

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// learnVariable()
	// learnTypeData()
	// constAndOperator()
	// learnArray()
	learnCondition()
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
	// number := [3]int{22, 44, 66}
	// num := []int{100, 300, 500, 700}
	arry := [3]int{1, 2, 3}
	arry[2] = 50
	// spew.Dump(arry)
	for k, v := range arry {
		fmt.Println(k, v)
	}
	// spew.Dump(num)
	// spew.Dump(number)

	skl := [2][3]string{{"SD 15", "SD 14", "SD YPK"}, {"SMP 2", "SMA 1", "UNSIA"}}
	for _, v := range skl {
		fmt.Println("\n", strings.Repeat("=", 55), "\n")
		for _, hasil := range v {
			fmt.Println(hasil)
		}
	}

	// slice

	fmt.Println("\n", strings.Repeat("=", 55), "\n")

	animal := []string{"Dog", "Pig", "Bird"}
	// menambahkan
	animal = append(animal, "Kanguru")

	for _, v := range animal {
		fmt.Println(v)
	}

	// copy
	fruit1 := []string{"Mangga", "Pepaya", "Rujak"}
	fruit2 := []string{"Nasi", "ikan", "sayur"}
	fmt.Println("insdex :", fruit1[1:])
	fruit1 = append(fruit1, fruit2...)
	fmt.Println("Hasil Fruit1 :", fruit1)
	hasl := copy(fruit1, fruit2)
	fmt.Println(hasl)
}

func learnCondition() {
	a := true
	b := false

	if a == b {
		fmt.Println("Welcome")
	} else {
		fmt.Println("Kembali Lain waktu")
	}

	currentYear := 2021
	age := currentYear - 2000

	if age <= 17 {
		fmt.Println("Anda Belum bisa cetak KTP")
	} else {
		fmt.Println("Silahkan Cetak KTP")
	}
}

// 2:00 jam
