package main

import (
	"fmt"
	"os"
	"strconv"
)

// struct untuk menyimpan data teman
type Teman struct {
    Absen   int
    Nama    string
    Alamat  string
    Pekerjaan string
    Alasan  string
}

// function untuk menampilkan data teman berdasarkan nomor absen
func tampilkanDataTeman(absen int, daftarTeman []Teman) {
    for _, teman := range daftarTeman {
        if teman.Absen == absen {
            fmt.Printf("Nama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan memilih kelas Golang: %s\n", teman.Nama, teman.Alamat, teman.Pekerjaan, teman.Alasan)
            return
        }
    }
    fmt.Println("Data teman dengan nomor absen tersebut tidak ditemukan.")
}

func main() {
    // data teman
    daftarTeman := []Teman{
        {1, "John", "Jakarta", "Software Engineer", "Ingin mempelajari bahasa pemrograman Golang"},
        {2, "Jane", "Bogor", "Data Analyst", "Ditugaskan untuk mempelajari Golang untuk proyek perusahaan"},
        {3, "Bob", "Tangerang", "IT Support", "Ingin meningkatkan skill programming dan memilih Golang sebagai bahasa baru yang dipelajari"},
    }

    // memeriksa argumen yang diberikan pada command line
    if len(os.Args) < 2 {
        fmt.Println("Penggunaan: go run biodata.go <nomor_absen>")
        return
    }
    absenStr := os.Args[1]
    absen, err := strconv.Atoi(absenStr)
    if err != nil {
        fmt.Println("Nomor absen harus berupa bilangan bulat.")
        return
    }

    // memanggil function untuk menampilkan data teman berdasarkan nomor absen yang diberikan
    tampilkanDataTeman(absen, daftarTeman)
}
