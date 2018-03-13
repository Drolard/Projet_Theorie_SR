Ce fichier décrit en premier lieu comment installer ce projet puis comment l'utiliser :


>>>>>>>>>>>>>>>>>>>>     INSTALLATION     <<<<<<<<<<<<<<<<<<<<

Tout d'abord décompresser le fichier dans le dossier voulu

Ensuite il faut mettre à jour la variable d'environnement GOPATH, pour ce faire entrer dans un terminal :

	export GOPATH=[AdresseFichierGo]

Où [AdresseFichierGo] correspond au chemin absolu du fichier contenant le projet (ATTENTION : ne pas mettre les crochet [] ) 
Ce fichier doit avoir une architecture similaire à celle-ci :

[AdresseFichierGo]
	│
	├── bin
	├── src
	│   ├── projet_Client
	│   │		└── client.go
	│   └── projet_Serveur
	│		├── collecteur.go
	│		├── repartiteur.go
	│		├── serveur.go
	│		└── travailleur.go
	├── bin
	└── README.txt

EXEMPLE : export GOPATH=/mnt/Théorie_SR/MonProjetGo

Pour voir si $GOPATH à bien été configuré faîtes 
	
	echo $GOPATH


	>>>>>>>>>> INSTALLATION/CONFIGUATION DU SERVEUR <<<<<<<<<<


		>>>>> CONFIGURATION DU SERVEUR <<<<<

		De base le serveur utilise le protocole tcp et écoute sur le port 10000,
		si cela ne vous convient pas vous pouvez le changer, pour ce faire suivez les instruction ci-dessous:
	
			Pour configurer sur quel port le serveur écoute et via quel protocole il faut changer le fichier collecteur.go à la ligne 18.
			Cette ligne doit être de la forme :
		
				listener,err:=net.Listen("[Protocole]",":[Port]")

			Où [Protocole] correspond au protocole utilisé (tcp de base) 
			et [Port] correspond au port utilisé par le serveur (10000 de base) 

			(ATTENTION : ne pas mettre les crochet [] ) 

			EXEMPLE : De base la ligne est : listener,err:=net.Listen("tcp",":10000")


		>>>>> INSTALLATION DU SERVEUR <<<<<

		Le serveur doit ensuite être proprement installé, pour ce faire dans un terminal entrer la commande 
		
			go install projet_Serveur

		Si la variable d'environnement GOPATH à bien été initialisé, un fichier projet_Serveur sera créé dans le dossier bin.




	>>>>>>>>>> INSTALLATION/CONFIGUATION DU CLIENT <<<<<<<<<<


		>>>>> CONFIGURATION DU CLIENT <<<<<

		De base le client utilise essaie de se connecter à l'adresse 127.0.0.1 (machine locale) et au port 10000 via le protocole tcp,
		si cela ne vous convient pas vous pouvez le changer, pour ce faire suivez les instruction ci-dessous:
	
			Pour configurer à quelle adresse et à quel port le client va vouloir se connecter et aussi via quel protocole, il faut changer le fichier client.go à la ligne 17.
			Cette ligne doit être de la forme :
		
				connexion,err := net.Dial("[Protocole]","[AdresseIP]:[Port]")

			Où [Protocole] correspond au protocole utilisé (tcp de base), 
			   [AdresseIP] correspond à l'adresse ip où le client se connecte (127.0.0.1 de base)
			et [Port] correspond au port que le client tente d'accéder (10000 de base) 

			(ATTENTION : ne pas mettre les crochet [] ) 

			EXEMPLE : De base la ligne est : connexion,err := net.Dial("tcp","127.0.0.1:10000")


		>>>>> INSTALLATION DU CLIENT <<<<<

		Le client doit ensuite être proprement installé, pour ce faire dans un terminal entrer la commande 
		
			go install projet_Client

		Si la variable d'environnement GOPATH à bien été initialisé, un fichier projet_Client sera créé dans le dossier bin.



>>>>>>>>>>>>>>>>>>>>     UTILISATION     <<<<<<<<<<<<<<<<<<<<

Pour utiliser ce projet Serveur/Client il faut tout d'abord lancer un serveur puis s'y connecter avec des clients pour ce faire voici le démarche à suivre :



	>>>>>>>>>> LANCER UN SERVEUR <<<<<<<<<<

	Après avoir configuré le serveur à votre guise (voir partie configuration serveur) ou l'avoir laissé avec sa configuration de base, puis l'avoir installé, il faut le lancer.
	Pour ce faire il faut via un terminal lancer le fichier projet_Serveur contenu dans le fichier bin, cela peut être fait de deux façons :

		$GOPATH/bin/projet_Serveur

	(La variable d'environnement GOPATH doit être bien initialisée)
		
	Ou bien il faut se placer dans le dossier bin et lancer :
	
		./projet_Serveur

	(utiliser la commande cd pour se déplacer de dossier en dossier)

	Le serveur est lancé et attend des connexions.
	
	ATTENTION cela bloque votre terminal, celui-ci est devenu le serveur, vous ne pouvez donc plus entrer de commande, si avez besoin d'en faire, ouvrez un autre terminal.



	>>>>>>>>>> LANCER UN CLIENT <<<<<<<<<<
	
	Après avoir configuré le client à votre guise (voir partie configuration client) ou l'avoir laissé avec sa configuration de base, puis l'avoir installé, il faut le lancer.
	Pour ce faire il faut via un terminal lancer le fichier projet_Client contenu dans le fichier bin, cela peut être fait de deux façons :

		$GOPATH/bin/projet_Client

	(La variable d'environnement GOPATH doit être bien initialisée)
		
	Ou bien il faut se placer dans le dossier bin et lancer :
	
		./projet_Client

	(utiliser la commande cd pour se déplacer de dossier en dossier)

	Le Client est lancé, si il arrive à se connecter, il peut ensuite envoyer des messages au serveur.



>>>>>>>>>>>>>>>>>>>>     FONCTIONNEMENT DU SERVEUR     <<<<<<<<<<<<<<<<<<<<

Le serveur peut recevoir plusieurs connexions et dispose d'un nombre de travailleur limité.
Si un client entre un entier x (entier = seul type de message valide avec "q" qui permet au client de se déconnecter) alors le serveur va utiliser un de ses travailleur et faire une boucle de x tours.






























