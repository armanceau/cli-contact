package main

import (
	"strings"
	"testing"
)

func AjouterContactDirect(nom, email string) Contact {
	contact := Contact{
		ID:    nextID,
		Nom:   strings.TrimSpace(nom),
		Email: strings.TrimSpace(email),
	}
	contacts[nextID] = contact
	nextID++
	return contact
}

func TestAjouterContact(t *testing.T) {
	contacts = make(map[int]Contact)
	nextID = 1

	c := AjouterContactDirect("Arthur", "arthur@mail.com")

	if len(contacts) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(contacts))
	}
	if contacts[c.ID].Nom != "Arthur" {
		t.Errorf("Expected Nom 'Arthur', got '%s'", contacts[c.ID].Nom)
	}
	if contacts[c.ID].Email != "arthur@mail.com" {
		t.Errorf("Expected Email 'arthur@mail.com', got '%s'", contacts[c.ID].Email)
	}
}

func TestSupprimerContact(t *testing.T) {
	contacts = make(map[int]Contact)
	nextID = 1

	c := AjouterContactDirect("Arthur", "arthur@mail.com")
	delete(contacts, c.ID)

	if len(contacts) != 0 {
		t.Errorf("Expected 0 contacts, got %d", len(contacts))
	}
}

func TestMettreAJourContact(t *testing.T) {
	contacts = make(map[int]Contact)
	nextID = 1

	c := AjouterContactDirect("Arthur", "arthur@mail.com")
	c.Nom = "Arthur Modifié"
	c.Email = "arthurmod@mail.com"
	contacts[c.ID] = c

	if contacts[c.ID].Nom != "Arthur Modifié" {
		t.Errorf("Expected Nom 'Arthur Modifié', got '%s'", contacts[c.ID].Nom)
	}
	if contacts[c.ID].Email != "arthurmod@mail.com" {
		t.Errorf("Expected Email 'arthurmod@mail.com', got '%s'", contacts[c.ID].Email)
	}
}

func TestListeContacts(t *testing.T) {
	contacts = make(map[int]Contact)
	nextID = 1

	AjouterContactDirect("Arthur", "arthur@mail.com")
	AjouterContactDirect("Bob", "bob@mail.com")

	if len(contacts) != 2 {
		t.Errorf("Expected 2 contacts, got %d", len(contacts))
	}
}
