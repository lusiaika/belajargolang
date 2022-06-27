package main

import (
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	dataBio()

}

func dataBio() {
	datas := []*Data{
		{
			nama:      "David",
			alamat:    "Cakung",
			pekerjaan: "karyawan",
			alasan:    "belajar golang",
		}, {
			nama:      "Lusia",
			alamat:    "Cakung",
			pekerjaan: "karyawan",
			alasan:    "belajar golang",
		}, {
			nama:      "Calvin",
			alamat:    "Cakung",
			pekerjaan: "karyawan",
			alasan:    "belajar golang",
		}, {
			nama:      "Nathanael Gilbert",
			alamat:    "Cakung",
			pekerjaan: "karyawan",
			alasan:    "belajar golang",
		}, {
			nama:      "Bakri",
			alamat:    "Cakung",
			pekerjaan: "karyawan",
			alasan:    "belajar golang",
		}, {
			nama:      "Chandra",
			alamat:    "Cakung",
			pekerjaan: "karyawan",
			alasan:    "belajar golang",
		}, {
			nama:      "Kevin",
			alamat:    "Cakung",
			pekerjaan: "karyawan",
			alasan:    "belajar golang",
		}, {
			nama:      "Michael",
			alamat:    "Cakung",
			pekerjaan: "karyawan",
			alasan:    "belajar golang",
		}, {
			nama:      "Kristian",
			alamat:    "Cakung",
			pekerjaan: "karyawan",
			alasan:    "belajar golang",
		}, {
			nama:      "Nathanael Orvin",
			alamat:    "Cakung",
			pekerjaan: "karyawan",
			alasan:    "belajar golang",
		},
	}
	newArgs := os.Args
	if len(newArgs) == 2 {
		i, err := strconv.Atoi(newArgs[1])
		if err != nil {
			fmt.Println("masukin angka!")
			return
		}
		v := datas[i-1]
		fmt.Printf(" Nama : %s\n Alamat : %s\n Pekerjaan : %s\n Alasan mengikuti kelas golang : %s\n", v.nama, v.alamat, v.pekerjaan, v.alasan)

	}

}
