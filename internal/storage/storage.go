package storage

import "fmt"

type Contact struct {
	ID    int    `json:"contact_id"`
	Nom   string `json:"nom"`
	Email string `json:"email"`
}

type Storer interface {
	Ajouter(c Contact) Contact
	Lister() []Contact
	Supprimer(ID int) bool
	MettreAJour(c Contact) (Contact, bool)
	Recuperer(ID int) (Contact, bool)
	NextID() int
}

var ErrContactNotFound = func(id int) error {
	return fmt.Errorf("Contact avec l'ID non trouvé")
}
