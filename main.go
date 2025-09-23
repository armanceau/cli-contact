package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nomFlag := flag.String("nom", "", "Nom du contact à ajouter")
	emailFlag := flag.String("email", "", "Email du contact à ajouter")

	flag.Parse()
	if *nomFlag != "" && *emailFlag != "" {
		contact := Contact{
			ID:    nextID,
			Nom:   strings.TrimSpace(*nomFlag),
			Email: strings.TrimSpace(*emailFlag),
		}
		contacts[nextID] = contact
		fmt.Println("Contact ajouté via flag :", contact.ID, contact.Nom, contact.Email)
		nextID++
		return
	}
	menu()
}

type Contact struct {
	ID    int
	Nom   string
	Email string
}

var contacts = make(map[int]Contact)
var nextID = 1

func menu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- cli contact ---")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Liste des contacts")
		fmt.Println("3. Supprimer un contact")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter")
		input, _ := reader.ReadString('\n')
		choix := strings.TrimSpace(input)

		switch choix {
		case "1":
			ajouterContact(reader)
		case "2":
			listeContacts()
		case "3":
			supprimerContact(reader)
		case "4":
			mettreAJourContact(reader)
		case "5":
			fmt.Println("À bientôt 👋")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}

}

func New(ID int, Nom string, Email string) Contact {
	contact := Contact{ID, Nom, Email}
	return contact
}

func ajouterContact(reader *bufio.Reader) {
	fmt.Print("Nom : ")
	nom, _ := reader.ReadString('\n')
	fmt.Print("Email : ")
	email, _ := reader.ReadString('\n')

	c := New(nextID, nom, email)

	c.ajouterContact()
}

func (c Contact) ajouterContact() {
	contacts[c.ID] = c
	fmt.Println("Contact ajouté ✅ :", c.ID, c.Nom, c.Email)
	nextID++
}

func listeContacts() {
	if len(contacts) == 0 {
		fmt.Println("Aucun contact trouvé ❌")
	}
	for _, c := range contacts {
		c.afficherContact()
	}
}

func (c Contact) afficherContact() {
	fmt.Printf("[%d] %s - %s\n", c.ID, c.Nom, c.Email)
}

func supprimerContact(reader *bufio.Reader) {
	fmt.Print("ID du contact à supprimer : ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("ID invalide")
		return
	}
	if _, ok := contacts[id]; ok {
		delete(contacts, id)
		fmt.Println("Contact supprimé ✅")
	} else {
		fmt.Println("Contact introuvable ❌")
	}
}

func mettreAJourContact(reader *bufio.Reader) {
	fmt.Print("ID du contact à mettre à jour : ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	contact, ok := contacts[id]
	if !ok {
		fmt.Println("Contact introuvable ❌")
		return
	}

	fmt.Print("Nouveau nom (laisser vide pour garder actuel) : ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)
	if nom != "" {
		contact.Nom = nom
	}

	fmt.Print("Nouvel email (laisser vide pour garder actuel) : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		contact.Email = email
	}

	contacts[id] = contact
	fmt.Println("Contact mis à jour ✅ :", contact.ID, contact.Nom, contact.Email)
}
