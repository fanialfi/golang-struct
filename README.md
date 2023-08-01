# Struct

struct adalah sebuah tipe data yang digunakan untuk mengelompokkan data dari tipe yang berbeda, struct mirip dengan class dengan bahasa pemrograman lain, tetapi struct tidak memiliki method.
Dari sebuah struct kita bisa buat variabel baru yang memiliki atribut sama seperti skema struct tersebut.

Untuk mendeklarasikan sebuah struct digunakan keyword `type` yang diikuti dengan nama struct dan kurung kurawal yang berisi nama dan tipe field.

<details>
  <summary>example</summary>

```go
type person struct{
  name string
  age int
}
```

</details><br>

struct `person` diatas memiliki 2 field yaitu _name_ dan _age_.
Object / variabel yang dibuat dengan struct ini nantinya memiliki skema atau struktur yang sama.

struct `person` diatas akan dimanfaatkan dalam pembuatan variabel, properti variabel tersebut diisi kemudian ditampilkan.

field dari `person` dapat diakses dengan menggunakan tanda titik (`.`).

<details>
  <summary>example</summary>

```go
var fani person

fani.name = "fani"
fani.age = 17

fmt.Printf("nama :", fani.name) // fani
fmt.Printf("umur :", fani.age) // 17
```

</details><br>

selain menginisialisasi nilainya secara satu persatu, inisialisasi variabel struct bisa dilakukan secara langsung dengan cara menambahkan kurung kurawal (`{}`) dan nilai masing masing field bisa juga diisi pada saat inisialisasi.

<details>
  <summary>example</summary>

```go
type person struct{
  name string
  age int
}

var fani = person{
  name: "fani",
  age: 17,
}

var alfi = person{}
alfi.name = "alfi"
alfi.age = 17
```

</details><br>

pada contoh kode di atas, variabel `fani` menampung struct `person` yang kemudian field field nya diisi langsung pada saat pendeklarasian variabel-nya.

selain itu variabel yang dibuat dari type struct bisa diambil nilai pointer nya, dan bisa disimpan pada variabel bertipe pointer struct.
Ataupun membuat variabel bertipe pointer struct.

<details>
  <summary>example</summary>

```go
type person struct{
  name string
  age int
}
var fani = person{
  name: "fani",
  age: 17
}

var fani2 *person = &fani // variabel pointer person mengambil datanya dari variabel fani

var alfi *person = new(person) // menggunakan keyword new() untuk membuat variabel bertipe pointer

// mengisi setiap field variabel alfi
alfi.name = "alfi"
alfi.age = 17
```
</details><br>

dari contoh di atas variabel `fani2` adalah variabel bertipe pointer person, dan variabel `fani2` tersebut menampung nilai referensi dari variabel `fani`.

sebuah struct juga bisa di embed atau dimasukkan kedalam sebuah struct lain, cara ini sering disebut dengan embeded struct, agal lebih mudah di pahami berikut contohnya :

<details>
  <summary>example</summary>

```go
type person struct{
  name string
  age int
}
type student struct{
  person
  kelas int
}

func main(){
  var fani = student{} // menginisialisasi secara langsung dengan nilai default zero value dari tiap tiap field

  fani.name = "fani"
  fani.age = 17
  fani.kelas = 3
}
```
</details><br>

pada code di atas struct `student` diatas disisipkan struct `person` dengan properti yang telah tersedia adalah _name_ dan _age_, pada struct `student` juga disiapkan field dengan nama `kelas`

embeded struct seperti diatas adalah **mutable**, nilai propertinya bisa diubah.
Untuk properti yang bukan properti asli (properti turunan dari struct lain) bisa diakses dengan cara mengakses struct _parent_-nya terlebih dahulu.

Jika salah satu field sebuah struct sama dengan field struct yang di embed, maka pengaksesan field nya harus secara explisit atau jelas.

<details>
  <summary>example</summary>

```go
type person struct {
  name string
  age int
}

type student struct {
  person
  age int
  kelas int
}

func main(){
  var std1 =  student{}
  std1.name = "fani"
  std1.person.age = 17 // age of person
  std1.age = 18 // age of student
  std1.kelas = 4

  fmt.Println("nama :", std1.name)
  fmt.Println("umur person :", std1.person.age)
  fmt.Println("umur :", std1.age)
  fmt.Println("kelas :",std1.kelas)
}
```
</details><br>

pengisian nilai field sub-stract bisa dilakukan dengan langsung memasukkan variabel struct yang tercetak dari struct yang sama.

<details>
  <summary>example</summary>

```go
type person struct {
  name string
  age int
}

type student struct {
  person
  age int
  kelas int
}

func main(){
  var p1 = person{"fani", 17}
  var s1 = student{p1,17,4}
}
```
</details><br>

pada deklarasi `s1` field `person` diisi oleh variabel `p1`.

sebelumnya semua struct di atas harus di deskripsikan terlebih dahulu, namun ada juga anonymous struct, anonymous struct adalah sebuah struct yang tidak dideklarasikan di awal sebagai tipe data baru, melainkan langsung dibuat ketika pembuatan variabel struct, teknik ini cukup efisien untuk pembuatan variabel struct dimana struct tersebut hanya dipakai sekali.

<details>
  <summary>example</summary>

```go
type person struct {
  name string
  age int
}

func main(){
  var p1 = person{"fani", 17}
  var s1 = struct {
    person
    kelas int
  }{}

  s1.person = p1
  s1.kelas = 4

  fmt.Println(s1)
  fmt.Println(p1)
}
```
</details><br>

pada contoh di atas variabel `s1` berisi anonymous struct, dimana didalamnya meng embed struct _person_.

Aturan dalam pembuatan anonymous struct adalah deklarasi harus diikuti dengan inisialisasi, bisa dilihat pada variabel `s1` setelah deklarasi struct, terdapat kurung kurawal untuk menginisialisasi struct dengan zero value, meskipun nilainya tidak di sisipkan diawal, kurung kurawal tetap harus ditulis.

Selain anonymous struct, sebuah struct juga bisa dikombinasikan dengan slice, cara penggunaanya pun mirip seperti pada slice dan map, cukup tambahkan `[]` sebelum tipe data pada saat deklarasi.

<details>
  <summary>example</summary>

```go
type person struct {
  name string
  age int
}

func main(){
  var persons []person = []person{
    {"fani", 17},
    {"alfi", 17},
    {"fanialfi", 17},
  }

  fmt.Println(persons)

  for _, person := range persons {
    fmt.Printf("%s umurnya %d\n", person.name, person.age)
  }
}
```
</details><br>

selain itu juga anonymous struct juga bisa dikombinasikan bersama slice, dan nilai awalnya juga bisa langsung diinisialisasi pada saat deklarasi.

<details>
  <summary>example</summary>

```go
type person struct {
	nama string
	umur int
}

func main() {
	var students = []struct {
		person
		kelas int
	}{
		{person{"fani", 17}, 4},
		{person{"alfi", 17}, 4},
		{person{"fanialfi", 17}, 4},
	}

	for _, student := range students {
		fmt.Println(student)
	}
}
```
</details><br>

cara lain untuk deklarasi anonymous struct adalah dengan menggunakan keyword var.

<details>
  <summary>example</summary>

```go
func main(){
  var person1 = struct{
    nama string
    umur int
  }

  person1.nama = "fani"
  person1.umur = 17
  fmt.Println(person1)

  var person2 = struct {
    nama string
    umur int
  }{"fanialfi", 17}
  fmt.Println(person2)
}
```
</details><br>

pada dua variabel di atas pada dasarnya sama.