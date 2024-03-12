package main

import(
	"fmt"	
	"reflect"
) 

func main(){
	var nama string
	var angka int
	var koma float32
	var bener bool
	var code rune
	var bite byte
	// ------- ARRAY ---------
	var arrAngka [5]int
	arrAngka = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Ini angka array --> ",  arrAngka[2])
	fmt.Println(reflect.TypeOf(arrAngka))
	// ------- SLICE ---------

	arrSlice := []interface{}{1, 2, 3, 4, 5}
	fmt.Println("Ini angka seluruh slice sebelum append --> ",  arrSlice)
	arrSlice = append(arrSlice, "sekolah")
	fmt.Println("Ini angka seluruh slice setelah append --> ",  arrSlice)
	// ------- MAP ---------
	mapSekolah := make(map[string]interface{})
	mapSekolah["devops"] = "Jakarta"
	mapSekolah["data"] = "Bogor"
	
	fmt.Println("Ini lokasi sekolah devops --> ",  mapSekolah["devops"])
	mapSekolah["backend"] = "Surabaya"
	fmt.Println("Ini lokasi sekolah backend --> ",  mapSekolah["backend"])
	mapSekolah["ip"] = 317
	fmt.Println("Ini lokasi sekolah ip --> ",  mapSekolah["ip"])
	sekolah := "Digital Skola"
	nama = "Farhan"
	angka = 18
	koma = 3.14
	bener = true
	fmt.Println(nama)
	fmt.Println(angka)
	fmt.Println(koma)
	fmt.Println(bener)
	fmt.Println(code)
	fmt.Println(bite)
	fmt.Println(reflect.TypeOf(sekolah))
}