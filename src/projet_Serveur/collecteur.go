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

/** Fonction qui ouvre une connexion sur le port 10000 de la machine et qui va gérer tout les clients se connectant à ce serveur
* @param fromCollector, canal d'une place qui va contenir les entrées valide des clients
*/
func Collecter(fromCollector chan Task) {
  listener, err:=net.Listen("tcp",":10000") // Ouvre la connexion
  if err != nil {
    fmt.Println("Erreur de Listen")
  }
  for {
    connexionBase, err := listener.Accept() // Attend la connexion d'un client
    if err != nil{
      fmt.Println("Erreur de Connexion")
    }
    wg.Add(1)
    go EcouterClient(connexionBase, fromCollector) // Lance une goRoutine dès qu'un client s'est connecté
  }

//Si le programme sort de la boucle infinie
    fmt.Println("ShutDown du serveur")
}


/** Fontion qui gère un et un seul client en l'écoutant et en mettant ses entrées valides dans un canal
* @param connexion, correspond au client
* @param fromCollector, canal d'une place qui va contenir les entrées valide du client
*/
func EcouterClient(connexion net.Conn, fromCollector chan Task){
  defer wg.Done()
  var msg int
  reader := bufio.NewReader(connexion) // On associe un reader au client
  for {
    message, err := reader.ReadString('\n')
    if err == nil {
        msg, err = strconv.Atoi(strings.TrimRight(message, "\n"))
        if err == nil { // Dès que ce client entre un message valide (càd qu'il entre un entier)
          fromCollector <- Task{msg} // On communique le message au repartiteur via le canal fromCollector
        }
      }
    }
}
