package main

import "fmt"

// membuat struct
type person struct {
	name string
	age  int
}

type std2 struct {
	person // embeded struct
	grade  int
}

type student struct {
	name  string
	grade int
}

// embeded struct dengan nama properti yang sama
type stc1 struct {
	name string
	age  int
}
type stc2 struct {
	stc1
	age   int
	grade int
}

// nested struct
type nestedStruct struct {
	person struct {
		name string
		age  int
	}
	kelas int
	hobi  []string
}

// tag properti dalam struct
type tag struct {
	name string `nama`
	umur int    `umur`
}

func main() {
	// tag properti dalam struct
	fmt.Printf("\ntag properti dalam struct :\n")
	var tagProp = tag{"fani", 18}
	fmt.Println(tagProp)

	// nested struct
	fmt.Printf("\nnested struct :\n")
	var fani nestedStruct
	fani.person.name = "fani"
	fani.person.age = 17
	fani.kelas = 14
	fani.hobi = []string{"ngoding", "membaca", "tidur"}
	fmt.Println(fani)

	// deklarasi anonymous struct menggunakan keyword var
	fmt.Printf("\ndeklarasi anonymous struct menggunakan keyword var :\n")
	var varStruct struct {
		person
		kelas int
	}

	varStruct.age = 18
	varStruct.name = "fani"
	varStruct.kelas = 14
	fmt.Println(varStruct)

	// inisialisasi slice dengan anonymous struct
	fmt.Printf("\ninisialisasi slice dengan anonymous struct :\n")
	var persons = []struct {
		person
		kelas int
	}{
		{person: person{name: "fani", age: 18}, kelas: 13},
		{person{"alfi", 18}, 1},
		{person{"fanialfi", 19}, 9},
	}
	for _, elm := range persons {
		fmt.Println(elm)
	}

	// kombinasi slice dan struct
	fmt.Printf("\nkombinasi slice dan struct :\n")
	var kombinasi = []person{
		{"fani", 17},
		{"alfi", 18},
		{"wick", 23},
		{name: "ethan", age: 20},
		{"borneo", 22},
	}
	for _, elm := range kombinasi {
		fmt.Printf("%s umurnya %d\n", elm.name, elm.age)
	}

	// anonymous struct
	fmt.Printf("\nanonymous struct :\n")
	var anonStruct = struct {
		person
		grade int
	}{}
	anonStruct.person = person{"fani", 18}
	anonStruct.grade = 2
	fmt.Println(anonStruct)

	// pengisian nilai sub-struct secara langsung
	fmt.Printf("\npengisian nilai sub-struct secara langsung :\n")
	var vStc2 = stc1{"fani", 18}
	var vStc3 = stc2{stc1: vStc2, age: 19, grade: 2}
	fmt.Println(vStc3)

	// embeded struct dengan nama properti yang sama
	fmt.Printf("\nembeded struct dengan nama properti yang sama :\n")
	var vStc1 = stc2{}
	vStc1.stc1.name = "fani alfirdaus"
	vStc1.stc1.age = 18 // pengaksesan dan pengisian element harus dilakukan secara eksplisit atau jelas
	vStc1.age = 19
	vStc1.grade = 2
	fmt.Println("name :", vStc1.stc1.name)
	fmt.Println("stc1.age :", vStc1.stc1.age)
	fmt.Println("age :", vStc1.age)
	fmt.Println("grade :", vStc1.grade)

	// embeded struct
	fmt.Printf("\nembeded struct :\n")
	var ss1 = std2{}
	ss1.name = "fani alfirdaus"
	ss1.age = 18
	ss1.grade = 2

	fmt.Println("name :", ss1.name)
	fmt.Println("age :", ss1.age)
	fmt.Println("age :", ss1.person.age)
	fmt.Println("grade :", ss1.grade)

	// penerapa struct
	fmt.Printf("\npenerapan struct :\n")
	var s1 student
	s1.name = "fani alfirdaus"
	s1.grade = 100

	fmt.Printf("name\t: %s\n", s1.name)
	fmt.Printf("nilai\t:%d\n", s1.grade)
	fmt.Println(s1)

	fmt.Println()

	// inisialisasi struct
	var s2 = student{} // nilai default => zero value sesuai tipe data-nya
	s2.name = "fani"
	s2.grade = 18
	fmt.Println(s2)

	var s3 = student{name: "alfi", grade: 18}
	fmt.Println(s3)

	var s4 = student{"fanialfi", 18}
	fmt.Println(s4)

	// membuat variabel bertipe pointer struct

	// var s5 *student
	// var nm = "sembarang"
	// var ag = 18
	// // ketika mengakses grade dan name dengan operator dereference, pointer s5 belum mengalokasikan memori dan go mencoba mengakses nilai dari alamat memori yang ditunjuk oleh s5.grade dan s5.name
	// *s5.grade = &ag
	// *s5.name = &nm
	// fmt.Println(*s5)

	var s5 *student = new(student) // menggunakan keyword new() karena new() mengembalikan pointer
	var nm = "asdf"
	var ag = 18
	s5.grade = ag
	s5.name = nm
	fmt.Println(*s5)
}
