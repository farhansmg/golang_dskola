package main

import(
	"fmt"	
	// "reflect"
) 

func main(){
	// while
	// i:=0
	// for i <= 10{
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// for pada umum nya
	for i:=0;i<=10;i+=1{
		fmt.Println(i)
	}

	// looping inside array
	// arrSlice := []interface{}{1, 2, 3, 4, 5}
	// for _,palue := range arrSlice{
	// 	fmt.Println(" value : ", palue)
	// }

	// looping inside map
	// mapSekolah := make(map[string]interface{})
	// mapSekolah["devops"] = "Jakarta"
	// mapSekolah["data"] = "Bogor"
	// mapSekolah["ip"] = 317
	// for _,value := range mapSekolah{
	// 	fmt.Println(" Value : ", value)
	// }

	// Buatlah program yang mencetak tabel perkalian 5 hingga 50 dengan menggunakan loop.
	// for i:=1;i<=10;i+=1{
	// 	result := i * 5
	// 	fmt.Println("hasil ", i, " dikali dengan 5 = ", result)
	// }

	//Buatlah program yang meminta pengguna untuk memasukkan sebuah kata, kemudian mencetak setiap huruf dari kata tersebut secara terpisah.
	var input string
	fmt.Print("Masukkan kata2 : ")
	fmt.Scan(&input)
	jumlah_a := 0
	for index,char := range input{
		if string(char) == "a"{
			jumlah_a += 1
		}
		fmt.Println("Huruf ke ", index, " adalah ", string(char))
	}
	fmt.Println("jumlah huruf a ada ", jumlah_a)

	// if
	// nilai := 65
	// if nilai >= 70{
	// 	fmt.Println("Lulus")
	// } else if nilai >= 60{
	// 	fmt.Println("Remidi")
	// } else{
	// 	fmt.Println("Tidak Lulus")
	// }

	// Switch
	// tirai := 3

	// switch tirai{
	// case 1:
	// 	fmt.Println("Selamat kamu mendapatkan Ferrari")
	// case 2:
	// 	fmt.Println("Selamat kamu mendapatkan rumah di senayan")
	// case 3:
	// 	fmt.Println("Selamat kamu dipecat")
	// }
}