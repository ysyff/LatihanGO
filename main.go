package main

import "fmt"

func main() {
    fmt.Println("Yusuf Syahputra || 1955617840-478")
   for i :=0; i<5; i++ {
      fmt.Println("Nilai i = ",i)
   }

   var j int = 0;
   for   j<5 {
      fmt.Println("Nilai j = ",j)
      j++
   }

   for pos, char := range "САШАРВО" {
      fmt.Printf("character %#U starts at byte position %d\n", char, pos)
   }

   for {
      if j++; j <= (10) {
         fmt.Println("Nilai j = ",j)
      }else{
         break
      }
   }
}