package main

import (
  "fmt"
  "net"
  "bufio"
  "os"
)

func main(){

connexion,err := net.Dial("tcp","172.21.66.106:10000")
if err == nil {

  fmt.Println("Connexion effectué")

  reader := bufio.NewReader(connexion)
  writer := bufio.NewWriter(connexion)

  myReader := bufio.NewReader(os.Stdin)

  message,err:=reader.ReadString('\n')

  if err!=nil {
    fmt.Println("Erreur de réception de message")
  }
  fmt.Print(message)

  pseudoValide := true

  for pseudoValide {

    pseudo, _  := myReader.ReadString('\n')
    _, err = writer.WriteString(pseudo)
    writer.Flush()

    if err !=nil {
        fmt.Println("Erreur d'écriture")
    }

    message, err := reader.ReadString('\n')

    if err!=nil {
      fmt.Println("Erreur de réception de message 2")
    }

    fmt.Print(message)
    if message == "ok!\n" {
      pseudoValide = false
    }
  }

  deconnexion := true
  for deconnexion {

    message, _  := myReader.ReadString('\n')
    _, err = writer.WriteString(message)
    writer.Flush()

    if err !=nil {
        fmt.Println("Erreur d'écriture")
    }

    message, err := reader.ReadString('\n')

    if err!=nil {
      fmt.Println("Erreur de réception de message")
    }

    fmt.Print(message)
    if message != "ok!\n" {
      deconnexion = false
    }

  }


fmt.Println("Fin de connexion")
} else {
  fmt.Println("Erreur de connexion")
}

}
