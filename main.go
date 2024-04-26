package main

import (
	"fmt"
	"math"
)

func main() {
	//deklarasi variable dengan var
	var umur int32 = 31
	var name string
	var last = "Faisal"
	var (
		ValueSkor = 13.14
		ValueNama = "Edi Santoso"
	) //deklarasi variable dengan var multiple
	const phi = 3.14 //format constanta

	name = "ahmad " + last + " " + ValueNama //concat string
	nilaiReal := 32.5                        //tanpa var

	fmt.Println("hello world ", name)
	fmt.Println(umur)
	umur = umur + int32(nilaiReal-ValueSkor)

	//Operasi perbandingan (>, <, >=, <=, ==, !=)
	//format if nya begitu sensitif
	if umur > 30 {
		fmt.Println("If Result is ", umur)
	} else {
		fmt.Println(umur)
	}

	var result bool = umur == 40
	fmt.Println("Usia hasil bersifat ", result)
	//end

	//Coversion Operasi perbandingan

	var e byte = name[4]           //convert karakter dalam variabel name menjadi byte
	var eString string = string(e) //konversi menjadi string

	fmt.Println(name)
	fmt.Println(eString)

	//end conversion

	//type declarations untuk membuat alias variabel
	type zenitsu string
	type menikah bool

	var nomorKTP zenitsu = "98384756547890665"
	var status menikah = true

	fmt.Println(nomorKTP)
	fmt.Println(status)
	//end

	//operator matematika
	var a float32 = 11.4
	var b float32 = 17
	var x = 3
	var y = 17
	var x1 float64 = 3
	var y1 float64 = 17

	var c float32 = a * b
	var d float32 = a + b
	var n float32 = b / a
	var f float32 = a - b
	var g = y % x

	fmt.Println("Result C Variable : ", c)
	fmt.Println("Result D Variable : ", d)
	fmt.Println("Result N Variable : ", n)
	fmt.Println("Result F Variable : ", f)
	fmt.Println("Result G Variable : ", g)
	fmt.Println("Result Mod Class : ", math.Mod(y1, x1))

	//end

	//array
	var barang [3]string
	barang[0] = "Sepatu"
	barang[1] = "Dasi"
	barang[2] = "Peci"

	var nomor = [3]int{
		78,
		13,
		23,
	}

	nomor[2] = 99

	//end

	//Perulangan (hanya ada for loops)
	var in = 0
	var panjangarray int = len(barang) //hitung panjang array
	for in < panjangarray {
		fmt.Println(barang[in])
		in++
	}

	in = 0
	panjangarray = len(nomor)
	for in < panjangarray {

		fmt.Println("Pangan nomor pada array ", in, "adalah", nomor[in])
		in++
	}

	for tel := 0; tel < len(nomor); tel++ {
		if tel == 1 {
			continue //penggunaan continue
		}
		fmt.Println("TELL nomor pada array ", tel, "adalah", nomor[tel])

	}

	//end

	//Tipe data slice, potongan array yang jika diassign dapat mengubah data pada array yang dislice

	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = bulan[4:]
	in = 0
	panjangarray = len(slice1)
	slice1[1] = "Wow"
	for in < panjangarray {
		fmt.Println(slice1[in])
		in++
	}

	slice1[1] = "Juni."
	slice2 := append(slice1, "Wadidaw")

	panjangarray = len(slice2)
	in = 0
	for in < panjangarray {
		fmt.Println(slice2[in])
		in++
	}

	panjangarray = len(bulan)
	in = 0
	for in < panjangarray {
		fmt.Println(bulan[in])
		in++
	}

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Faisal"
	newSlice[1] = "Muhammad"
	fmt.Println(newSlice)

	iniArray := [...]int{1, 4, 5, 2, 8, 0}
	iniSlice := []int{1, 4, 5, 2, 8, 0}

	fmt.Println("Capacity slice1", cap(slice1))
	fmt.Println("Length newSlice", len(newSlice))
	fmt.Println("Capacity newSlice", cap(newSlice))

	fmt.Println(iniArray)

	fmt.Println(iniSlice)

	//end

	//Tipe data map, seperti array asosiatif

	orang := map[string]string{
		"nama": "Faisal",
		"ktp":  "63555667725436782",
	}
	orang["gender"] = "Laki-laki"

	fmt.Println(orang)
	fmt.Println(orang["nama"])
	fmt.Println(orang["ktp"])

	hewan := make(map[string]string)
	hewan["spesies"] = "Pongo"
	hewan["kaki"] = "2"
	fmt.Println(len(hewan))
	delete(hewan, "kaki")
	fmt.Println(hewan)
	fmt.Println(len(hewan))

	for index, list := range orang {
		fmt.Println("Index", index, "=", list)
	}

	for _, list := range orang {
		fmt.Println(list)
	}

	//end

	//Kondisi Switch

	switch hewan["spesies"] {
	case "Felis":
		fmt.Println("Kucing")
	case "Pongo":
		fmt.Println("Orang Utan")
		fmt.Println("Kera")
	case "Moose":
		fmt.Println("Rusa")
		fmt.Println("Kambing")
	default:
		fmt.Println("Kuda")
	}

	switch {
	case umur < 12:
		fmt.Println("Remaja")
	case umur >= 12 && umur < 40:
		fmt.Println("Dewasa")
	case umur >= 40 && umur < 60:
		fmt.Println("Tua")
	default:
		fmt.Println("Manula")
	}

	//end

	//Function
	sayHello() //function void
	fmt.Println(kalkulasi(9, 8, "-"))
	fisrtName, middleName, lastName := getFullName()
	fmt.Println(fisrtName, middleName, lastName)

	fisrtName2, _, lastName2 := getFullName()
	fmt.Println(fisrtName2, lastName2)

	iniNama, iniSkor, iniUmur := getReturnName() //function return name
	fmt.Println(iniNama)
	fmt.Println(iniSkor * 2)
	fmt.Println(iniUmur)
	//end

	//Variadic Function
	totalVariadic := getSumAll(11.4, 55.2, 23.4, 54.1) //function return name
	fmt.Println("Summary Sum A : ", totalVariadic)

	theSlice := []float64{45.2, 11.8, 76.2, 23.1}
	totalVariadic = getSumAll(theSlice...) //function return name
	fmt.Println("Summary Sum B : ", totalVariadic)
	//end

	//Function Value
	goodbye := sayHelloName
	fmt.Println(goodbye("Ariel"))
	//end
}
func sayHello() {
	fmt.Println("Hello Function")
	//Penggunaan function
}

func sayHelloName(name string) string {
	return "Hello " + name + ", Goodbye!"
}

func kalkulasi(nilai1 float32, nilai2 float32, jenis string) float32 {
	var hasil float32
	if jenis == "+" {
		hasil = nilai1 + nilai2
	} else if jenis == "-" {
		hasil = nilai1 - nilai2
	} else if jenis == "*" {
		hasil = nilai1 * nilai2
	} else if jenis == "/" {
		hasil = nilai1 / nilai2
	}

	fmt.Println(nilai1, jenis, nilai2, " = ", hasil)

	return hasil
	//end

}

// multiple function return

func getFullName() (string, string, string) {
	return "Muhammad", "Faisal", "Dimitrescu"
}

//end

// Named return values
func getReturnName() (nama string, skor float32, umur int) {

	nama = "lasiaf"
	skor = 78.4
	umur = 26

	return
}

//end

// variadic function

func getSumAll(nomor ...float64) float64 {
	totalA := 0.00
	for _, noe := range nomor {
		totalA += noe
	}

	return totalA
}

//end
