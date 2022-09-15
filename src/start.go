package start

import (
     "fmt"
     "strings"
     "strconv"
     "os"
)

var loker = make([]identity,0)
var lokerLength int

type identity struct {
     identityType  string
     identityNumber string
}

func Start() {

     Initialize()
     Instruction()
}

func Initialize(){
     fmt.Println("Selamat datang di Aplikasi Loker")

     //initialize loker length
     fmt.Println("Masukkan Jumlah Loker :")
     
     fmt.Scanln(&lokerLength)
     

     //create array of loker
     if(lokerLength <= 0 ){
          fmt.Println("WARNING.. Masukkan Jumlah loker minimal 1 !!!!!!")
          Initialize()
     }

     loker = make([]identity,lokerLength)
     fmt.Println("Jumlah Loker ", lokerLength)
     fmt.Println("----------------------------------------------")
     fmt.Println("Gunakan Instruksi Berikut :")
     fmt.Println("status => melihat status masing - masing loker")
     fmt.Println("input [Tipe Identitas] [No Identitas] =>  mengisi loker")
     fmt.Println("leave  [No Loker] => mengosongkan loker")
     fmt.Println("find [Tipe Identitas] => menampilkan data sesuai tipe identitas ")
     fmt.Println("search [No Identitas] => menampilkan data sesuai No identitas yg dicari")
     fmt.Println("exit => keluar dari aplikasi")
     fmt.Println("----------------------------------------------")

     // for i := 0; i < lokerLength ; i++ {
          //  item := identity{"SIM", "222"}
          // loker[1] = item
     // }

}


func Instruction(){
     
     var instruction string
     var value1 string
     var value2 string

     fmt.Scanln(&instruction, &value1, &value2 )

     instruction = strings.ToLower(instruction)
     fmt.Println("------------RESULT------------")
     
     switch instruction {
     case "init" :
          fmt.Println("WARNING !!! Jumlah loker sudah ditentukan")
     case "input" :
               Input(value1, value2)
     case "leave" :
               Leave(value1)
     case "exit" :
               fmt.Println("Menutup Aplikasi")
               os.Exit(0)
     case "status" :
               Status()
     case "search" :
               Search(value1)
     case "find" :
               Find(value1)
     default:    
          fmt.Println("WARNING !!! Instruksi tidak ditemukan")
     }

     Instruction()
}

func Input(val1 string, val2 string){
     
     if(val1 != "" && val2 != ""){
          var duplicate bool = false
          var item identity
          var status string = "full" 
          for i := 0; i < lokerLength; i++ {
               if(loker[i] == item){
                    status = "empty"
                    for x := 0; x < lokerLength; x++ {
                         if(loker[x] != item && loker[x].identityNumber == val2){
                              duplicate = true
                              break  
                         }
                    }
                    if(duplicate==false){
                         loker[i] = identity{strings.ToUpper(val1), val2}     
                         fmt.Println("Suskes memasukkan identitas ",strings.ToUpper(val1) ,val2, "di loker nomor ",i+1 )
                    }else{
                         fmt.Println("nomor identitas sudah ada")
                    }
                    break
               }
          }
          if(status == "full"){
               fmt.Println("Loker sudah penuh")     
          }
     }else{
          fmt.Println("Masukkan instruksi dengan benar")    
     }
}

func Leave(val1 string){
     i,err := strconv.Atoi(val1)
     err = err

     if(val1 == "" || isNumeric(val1) == false){
          fmt.Println("Masukkan instruksi dengan benar")    
     }else{
          var item identity 
          if(loker[i-1] == item){
               fmt.Println("Loker Nomor ",val1,"memang sudah kosong")
          }else{
               loker[i-1] = item     
               fmt.Println("Loker Nomor ",val1,"berhasil di kosongkan")
          }
     }
}

func Status(){
     fmt.Println("No \t Type \t Number ")
     for i := 0; i < lokerLength; i++ {
          var item identity 
          if(loker[i] == item){
               fmt.Println(i+1,"\t ","\t")
          }else{
                  fmt.Println(i+1,"\t ",loker[i].identityType,"\t ",loker[i].identityNumber)
          }
     }
}

func Find(val1 string){
     var result string = ""
     if(val1 != ""){
          for i := 0; i < lokerLength; i++ {
              if(loker[i].identityNumber == val1){
                    result = "Data Identitas "+val1+" ada di loker "+strconv.FormatInt(int64(i+1),10)
                    break
              } 
          }
          if(result==""){
               result = "Data tidak ditemukan"
          }
          fmt.Println(result)
     }else{
          fmt.Println("masukkan indtruksi yang benar")
     }
}

func Search(val1 string){
     if(val1 == ""){
          fmt.Println("Masukkan instruksi yang benar")     
     }else{
          var result string = ""
          for i := 0; i < lokerLength; i++ {
               if(loker[i].identityType == strings.ToUpper(val1)){
                    result = result+loker[i].identityNumber+", "
               }
          }
          if(result==""){
               result = "Data tidak ditemukan"
          }
          fmt.Println(result)     
     }
     
}

func isNumeric(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}
