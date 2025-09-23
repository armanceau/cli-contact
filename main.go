package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Storer interface {
	Ajouter(c Contact) Contact
	Lister() []Contact
	Supprimer(ID int) bool
	MettreAJour(c Contact) (Contact, bool)
	Recuperer(ID int) (Contact, bool)
	NextID() int
}

type Contact struct {
	ID    int
	Nom   string
	Email string
}

type MemoryStore struct {
	contacts map[int]Contact
	nextID   int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]Contact),
		nextID:   1,
	}
}

func (m *MemoryStore) Ajouter(c Contact) Contact {
	c.ID = m.nextID
	m.contacts[m.nextID] = c
	m.nextID++
	return c
}

func (m *MemoryStore) Lister() []Contact {
	var list []Contact
	for _, c := range m.contacts {
		list = append(list, c)
	}
	return list
}

func (m *MemoryStore) Supprimer(ID int) bool {
	if _, ok := m.contacts[ID]; ok {
		delete(m.contacts, ID)
		return true
	}
	return false
}

func (m *MemoryStore) Recuperer(ID int) (Contact, bool) {
	c, ok := m.contacts[ID]
	return c, ok
}

func (m *MemoryStore) MettreAJour(c Contact) (Contact, bool) {
	if _, ok := m.contacts[c.ID]; ok {
		m.contacts[c.ID] = c
		return c, true
	}
	return Contact{}, false
}

func (m *MemoryStore) NextID() int {
	return m.nextID
}

func main() {
	store := NewMemoryStore()

	nomFlag := flag.String("nom", "", "Nom du contact à ajouter")
	emailFlag := flag.String("email", "", "Email du contact à ajouter")

	flag.Parse()
	if *nomFlag != "" && *emailFlag != "" {
		contact := Contact{
			Nom:   strings.TrimSpace(*nomFlag),
			Email: strings.TrimSpace(*emailFlag),
		}
		c := store.Ajouter(contact)
		fmt.Println("Contact ajouté via flag :", c.ID, c.Nom, c.Email)
		return
	}
	menu(store)
}

func menu(store Storer) {
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

func New(ID int, Nom string, Email string) Contact {
	contact := Contact{ID, Nom, Email}
	return contact
}

func ajouterContact(reader *bufio.Reader, store Storer) {
	fmt.Print("Nom : ")
	nom, _ := reader.ReadString('\n')
	fmt.Print("Email : ")
	email, _ := reader.ReadString('\n')

	contact := Contact{
		Nom:   strings.TrimSpace(nom),
		Email: strings.TrimSpace(email),
	}

	c := store.Ajouter(contact)
	fmt.Println("Contact ajouté ✅ :", c.ID, c.Nom, c.Email)
}

func listeContacts(store Storer) {
	contacts := store.Lister()
	if len(contacts) == 0 {
		fmt.Println("Aucun contact trouvé ❌")
		return
	}
	for _, c := range contacts {
		fmt.Printf("[%d] %s - %s\n", c.ID, c.Nom, c.Email)
	}
}

func supprimerContact(reader *bufio.Reader, store Storer) {
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

func mettreAJourContact(reader *bufio.Reader, store Storer) {
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
