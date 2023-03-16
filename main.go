package main

import (
	"fmt"
	"strings"
)

func main() {

   
    fmt.Println("Challange 3|| Yusuf Syahputra || 1955617840-478")
   // Input kalimat
	kalimat := "selamat malam"

	// Konversi kalimat menjadi lowercase agar tidak terpengaruh huruf besar atau kecil
	kalimat = strings.ToLower(kalimat)


	// Deklarasi map untuk menyimpan jumlah kemunculan huruf
	hurufCount := make(map[rune]int)

	// Looping untuk setiap karakter dalam kalimat
	for _, karakter := range kalimat {
      // mencetak karakter 
      fmt.Printf("%c \n", karakter)
		// Tambahkan jumlah kemunculan huruf dalam map
		hurufCount[karakter]++
	}

	// Tampilkan hasil perhitungan kemunculan huruf
	fmt.Println("Hasil perhitungan kemunculan huruf:")
	for huruf, count := range hurufCount {
		fmt.Printf("%c: %d ,", huruf, count)
	}
}
