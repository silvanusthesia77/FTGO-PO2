package main

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

// dari funct latihan
type person struct {
	name  string
	age   int
	hobby string
}

func main() {
	// learnVariable()
	// learnTypeData()
	// constAndOperator()
	// learnArray()
	// learnCondition()
	// learnLooping()
	// msg, repons := greet("Junior", 20)
	// fmt.Println(msg, repons)
	// fmt.Println("\n", strings.Repeat("=", 50), "\n")
	// radical([]int{20, 22, 88}, "luiz", "arthur", "kandamy")
	// learnMap()
	// latihan()
	// ponter()
	struc()

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
	// a := true
	// b := false

	// if a == b {
	// 	fmt.Println("Welcome")
	// } else {
	// 	fmt.Println("Kembali Lain waktu")
	// }

	currentYear := 2026
	age := currentYear - 2008

	if age < 17 {
		fmt.Println("not Allow to get driver license")
	} else if age < 20 {
		fmt.Println("your age is 17 more and < 20")
	} else {
		fmt.Println("Allow to get driver license")
	}

	nilai := 6

	switch nilai {
	case 9:
		fmt.Println("Good Job")
	case 7:
		fmt.Println("Not Bad")
	case 6:
		fmt.Println("Try Again")
	default:
		fmt.Println("Semangat Coba lagi Nanti")

	}

	score := 8

	switch {
	case score == 8:
		fmt.Println("Good Job")
	case (score < 7) && (score > 5):
		fmt.Println("Not Bad")
	default:
		{
			fmt.Println("Belum Sesuai Target")
			fmt.Println("Jangan berkecil Hati, Kembali Lagi Nanti")
		}

	}

}
func learnLooping() {
	for angka := 1; angka <= 5; angka++ {
		fmt.Println(angka)
	}
	fmt.Println("\n", strings.Repeat("=", 55), "\n")
	for _, names := range []string{"thoby", "reza", "Luis", "Junior"} {
		fmt.Println(names)
	}

	num := 0

	for num < 5 {
		num++
		fmt.Println("nilai :", num)
	}

	fmt.Println("\n", strings.Repeat("=", 55), "\n")

	i := 0

	for {
		i++
		if i%2 != 0 {
			continue
		}
		fmt.Println("Nilai :", i)
		if i == 8 {
			break
		}
	}
}
func greet(name string, age int) (msg string, reponse string) {
	msg = fmt.Sprintf("halo saya %s berusia %d", name, age)
	return msg, ", Selamat Bergabung Bersama Barito Putra"
}

func radical(grade []int, names ...string) {
	for _, v := range grade {
		fmt.Println(v)
	}
	fmt.Println("\n", strings.Repeat("#", 30), "\n")
	for _, name := range names {
		fmt.Println(name)
	}
}

func learnMap() {
	people := map[string]int{}
	people["thby"] = 20
	people["luiz"] = 22
	people["reza"] = 25

	for person, v := range people {
		fmt.Println(person, v)
	}
}
func latihan() {
	people := map[string]person{}
	people["junior"] = person{age: 22, hobby: "memcaa"}
	people["arthur"] = person{age: 20, hobby: "lari"}

	for name, age := range people {
		fmt.Println(name, age)

	}
}
func ponter() {
	first := 4
	var second *int = &first
	fmt.Println(first)
	fmt.Println(&first)
	fmt.Println(*second)
	fmt.Println(second)

	var name string = "thby"
	var rename *string = &name
	fmt.Println(name)
	fmt.Println(&name)
	fmt.Println(*rename)
	fmt.Println(rename)

	fmt.Println("\n", strings.Repeat("=", 50), "\n")

	*rename = "Luiz"
	fmt.Println(*rename)
	fmt.Println(rename)
}
func struc() {
	var s1 person
	s1.name = "wanus"
	s1.age = 22
	s1.hobby = "sport"

	fmt.Println("Nama :", s1.name)
	fmt.Println("Age :", s1.age)
	fmt.Println("Hobby :", s1.hobby)

	emp2 := person{
		name:  "Luiz",
		age:   23,
		hobby: "Nurse",
	}
	fmt.Println(emp2)

}

// 40 jam (struct)
