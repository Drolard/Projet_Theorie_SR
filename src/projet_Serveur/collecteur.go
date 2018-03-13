package main

import (
  "fmt"
  "bufio"
  "strings"
  "net"
  "strconv"
  "sync"
)

var wg sync.WaitGroup

func Collecter(fromCollector chan int) {
  listener, err := net.Listen("tcp",":10000")
  if err != nil {
    fmt.Println("Erreur de Listen")
  }

  for {
    connexionBase, err := listener.Accept()
    if err != nil {
      fmt.Println("Erreur de Connexion")
    }
    wg.Add(1)
    go EcouterClient(connexionBase, fromCollector)
  }

  wg.Wait()
  fmt.Println("ShutDown du serveur")
}


func EcouterClient(connexion net.Conn, fromCollector chan int) {
  defer wg.Done()
  var msg int
  reader := bufio.NewReader(connexion)
  for {
    message, err := reader.ReadString('\n')
    if err == nil {
      msg, err = strconv.Atoi(strings.TrimRight(message, "\n"))
      if err == nil {
        fromCollector <- msg
      }
    }
  }
}
