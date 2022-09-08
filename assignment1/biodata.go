package main

import (
	"fmt"
	"os"
	"strconv"
)

type Participant struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	var participant = []Participant{
		{nama: "Try", alamat: "Jakarta Selatan", pekerjaan: "Devops", alasan: "pengembangan karir"},
		{nama: "Wiguna", alamat: "Jakarta Barat", pekerjaan: "Programmer", alasan: "pengembangan karir"},
		{nama: "Adhitya", alamat: "Jakarta Timur", pekerjaan: "UI/UX", alasan: "pengembangan karir"},
		{nama: "Primantara", alamat: "Jakarta Timur", pekerjaan: "UI/UX", alasan: "pengembangan karir"},
	}

	name := os.Args[1]
	var dataSlice []string
	for i, v := range participant {
		if v.nama == name {
			dataSlice = append(dataSlice, strconv.Itoa(i), v.nama, v.alamat, v.pekerjaan, v.alasan)

			// fmt.Printf("%#v", dataSlice)
			detailData(dataSlice...)
		}

	}
	if dataSlice == nil {
		fmt.Println("Whoops data not found")
	}

}

func detailData(biodata ...string) {
	fmt.Println("ID : ", biodata[0])
	fmt.Println("Nama : ", biodata[1])
	fmt.Println("Alamat : ", biodata[2])
	fmt.Println("Pekerjaan : ", biodata[3])
	fmt.Println("Alasan : ", biodata[4])
}
