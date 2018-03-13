package main

import (
  "fmt"
  "net"
  "bufio"
  "os"
  "strconv"
  "strings"
)

/* Fonction main du client, celle ci va permettre au client de discuter avec le serveur
*  C'est aussi cette méthode qui va vérifier si les entrées du client sont valides
*/
func main(){

connexion,err := net.Dial("tcp","127.0.0.1:10000") // Attend de se connecter à l'adresse mis en paramètre (Adresse à changer en fonction du serveur voulu)
if err == nil {
  writer := bufio.NewWriter(connexion) // Associe un writer à la connexion
  myReader := bufio.NewReader(os.Stdin) // Associe un reader à Stdin (cela permet de lire les entrées que le client fait sur son terminal)

  for {
    fmt.Println("Veuillez entrer un nombre (ou 'q' pour quitter)")
    message, _  := myReader.ReadString('\n') // Lit la première entrée du client se terminant par un retour chariot
    _ , err := strconv.Atoi(strings.TrimRight(message, "\n")) // Si l'entrée est un entier, err == nil
    if err == nil { // Si l'entrée est valide
      _, err = writer.WriteString(message) // envoie le message au serveur
      writer.Flush()
      if err !=nil {
          fmt.Println("Erreur d'écriture")
      }
    }else{
      if strings.TrimRight(message, "\n") == "q" { // Si l'entrée est "q", déconnecte du serveur via le break; (qort de la boucle infinie)
        fmt.Println("Deconnexion")
        break
      }else{
        fmt.Println("Erreur de saisie\n")
      }
    }
    }
    fmt.Println("Fin de connexion")
  }else{
    fmt.Println("Erreur de connexion")
  }
}
