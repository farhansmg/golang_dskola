package main

import(
	"fmt"	
	// "reflect"
	"strings"
) 

func hitungGaji(jamKerja float64, upahPerJam float64)float64{
	var gaji float64
	if jamKerja > 40{
		gaji = 40 * upahPerJam + ((jamKerja-40) * 1.5 * upahPerJam)
	}else {
		gaji = jamKerja * upahPerJam
	}
	return gaji
}

func main(){
	// Buatlah fungsi hitungGaji yang menerima jumlah jam kerja dan upah per jam sebagai parameter, 
	//dan mengembalikan gaji karyawan. Jika jumlah jam kerja lebih dari 40 jam, maka karyawan mendapatkan upah overtime 
	//(1.5 kali upah per jam untuk setiap jam overtime). Gunakan fungsi tersebut untuk 
	//menghitung gaji karyawan yang bekerja 45 jam dengan upah per jam 20.

	gaji := hitungGaji(45, 20)
	fmt.Println("Gaji karyawan : ", gaji)

	kata := "Digital Skola"
	fmt.Println("Panjang kata : ", len(kata))

	kata1 := "Di Jakarta"
	gabung := kata + " " + kata1
	fmt.Println("Panjang kata : ", gabung)
	fmt.Println("Lowercase kata : ", strings.ToLower(gabung))
	fmt.Println("Uppercase kata : ", strings.ToUpper(gabung))
	ganti := strings.Replace(gabung, " ", "_", -1)
	fmt.Println("Replacing kata : ", ganti)

}