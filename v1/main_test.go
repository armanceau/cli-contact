package main

import (
	"testing"
)

func TestAjouterContact(t *testing.T) {
	store := NewMemoryStore()

	store.Ajouter(Contact{Nom: "Arthur", Email: "arthur@mail.com"})

	contacts := store.Lister()
	if len(contacts) != 1 {
		t.Errorf("Expected 1 contact, got %d", len(contacts))
	}
	if contacts[0].Nom != "Arthur" {
		t.Errorf("Expected Nom 'Arthur', got '%s'", contacts[0].Nom)
	}
	if contacts[0].Email != "arthur@mail.com" {
		t.Errorf("Expected Email 'arthur@mail.com', got '%s'", contacts[0].Email)
	}
}

func TestSupprimerContact(t *testing.T) {
	store := NewMemoryStore()

	c := store.Ajouter(Contact{Nom: "Arthur", Email: "arthur@mail.com"})
	ok := store.Supprimer(c.ID)
	if !ok {
		t.Errorf("Expected deletion to succeed")
	}

	if len(store.Lister()) != 0 {
		t.Errorf("Expected 0 contacts, got %d", len(store.Lister()))
	}
}

func TestMettreAJourContact(t *testing.T) {
	store := NewMemoryStore()

	c := store.Ajouter(Contact{Nom: "Arthur", Email: "arthur@mail.com"})
	c.Nom = "Arthur Modifié"
	c.Email = "arthurmod@mail.com"

	updated, ok := store.MettreAJour(c)
	if !ok {
		t.Errorf("Expected update to succeed")
	}

	if updated.Nom != "Arthur Modifié" {
		t.Errorf("Expected Nom 'Arthur Modifié', got '%s'", updated.Nom)
	}
	if updated.Email != "arthurmod@mail.com" {
		t.Errorf("Expected Email 'arthurmod@mail.com', got '%s'", updated.Email)
	}
}

func TestListerContacts(t *testing.T) {
	store := NewMemoryStore()

	store.Ajouter(Contact{Nom: "Arthur", Email: "arthur@mail.com"})
	store.Ajouter(Contact{Nom: "Bob", Email: "bob@mail.com"})

	contacts := store.Lister()
	if len(contacts) != 2 {
		t.Errorf("Expected 2 contacts, got %d", len(contacts))
	}
}
