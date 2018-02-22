package main

import (
  "fmt"
  "net"
  "bufio"
  "os"
  "strconv"
  "strings"
)

func main(){

connexion,err := net.Dial("tcp","127.0.0.1:10000")
if err == nil {

  //reader := bufio.NewReader(connexion)
  writer := bufio.NewWriter(connexion)
  myReader := bufio.NewReader(os.Stdin)

  for {
    fmt.Println("Veuillez entrer un nombre (ou 'q' pour quitter)")
    message, _  := myReader.ReadString('\n')
    _ , err := strconv.Atoi(strings.TrimRight(message, "\n"))
    if err == nil {
      _, err = writer.WriteString(message)
      writer.Flush()
      if err !=nil {
          fmt.Println("Erreur d'écriture")
      }
    }else{
      if strings.TrimRight(message, "\n") == "q" {
        fmt.Println("Deconnexion")
        break
      }else{
        fmt.Println("Erreur de saisie\n")
      }
    }
  }
fmt.Println("Fin de connexion")
} else {
  fmt.Println("Erreur de connexion")
}

}
