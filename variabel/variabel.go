package main

import "fmt"

//variabel adalah tempat untuk menyimpan data
//golang tidak bisa menyimpan tipe data yang berbeda jenis dalam 1 variabel

func main() {

	//variabel dengan mendeklarasi tipe data
	var name string
	name = "hansen marcelino azali"
	fmt.Println(name)

	name = "hansen azali"
	fmt.Println(name)

	//variabel menggunakan :=  untuk awal dan tipe datanya menyesuaikan
	nama_lengkap := "hansen marcelino azali"
	fmt.Println("ini nama lengkap : ", nama_lengkap)

	//variabel var tanpa deklarasi tipe data
	var nama_akhir = "hansen azali"
	fmt.Println("ini adalah nama awal dan akhir : ", nama_akhir)

	//membuat banyak variabel
	var (
		first_name  = "hansen"
		middle_name = "marcelino"
		last_name   = "azali"
	)
	fmt.Println("ini adalah banyak variabel", first_name, middle_name, last_name)

}
