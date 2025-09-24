package app

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/armanceau/mini-crm/internal/storage"
)

func Run(store storage.Storer) {
	nomFlag := flag.String("nom", "", "Nom du contact à ajouter")
	emailFlag := flag.String("email", "", "Email du contact à ajouter")

	flag.Parse()
	if *nomFlag != "" && *emailFlag != "" {
		contact := storage.Contact{
			Nom:   strings.TrimSpace(*nomFlag),
			Email: strings.TrimSpace(*emailFlag),
		}
		c := store.Ajouter(contact)
		fmt.Println("Contact ajouté via flag :", c.ID, c.Nom, c.Email)
		return
	}
	menu(store)
}

func menu(store storage.Storer) {
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
			ajouterContact(reader, store)
		case "2":
			listeContacts(store)
		case "3":
			supprimerContact(reader, store)
		case "4":
			mettreAJourContact(reader, store)
		case "5":
			fmt.Println("À bientôt 👋")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}

}

func New(ID int, Nom string, Email string) storage.Contact {
	contact := storage.Contact{ID: ID, Nom: Nom, Email: Email}
	return contact
}

func ajouterContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Nom : ")
	nom, _ := reader.ReadString('\n')
	fmt.Print("Email : ")
	email, _ := reader.ReadString('\n')

	contact := storage.Contact{
		Nom:   strings.TrimSpace(nom),
		Email: strings.TrimSpace(email),
	}

	c := store.Ajouter(contact)
	fmt.Println("Contact ajouté ✅ :", c.ID, c.Nom, c.Email)
}

func listeContacts(store storage.Storer) {
	contacts := store.Lister()
	if len(contacts) == 0 {
		fmt.Println("Aucun contact trouvé ❌")
		return
	}
	for _, c := range contacts {
		fmt.Printf("[%d] %s - %s\n", c.ID, c.Nom, c.Email)
	}
}

func supprimerContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("ID du contact à supprimer : ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("ID invalide")
		return
	}
	if store.Supprimer(id) {
		fmt.Println("Contact supprimé ✅")
	} else {
		fmt.Println("Contact introuvable ❌")
	}
}

func mettreAJourContact(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("ID du contact à mettre à jour : ")
	input, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("ID invalide")
		return
	}

	contact, ok := store.Recuperer(id)
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

	if c, ok := store.MettreAJour(contact); ok {
		fmt.Println("Contact mis à jour ✅ :", c.ID, c.Nom, c.Email)
	} else {
		fmt.Println("Erreur lors de la mise à jour ❌")
	}
}
