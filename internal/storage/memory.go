package storage

import (
	"encoding/json"
	"os"
)

type JsonStore struct {
	filePath string
	contacts map[int]Contact
	nextID   int
}

func NewJsonStore(filePath string) *JsonStore {
	return &JsonStore{
		filePath: filePath,
		contacts: make(map[int]Contact),
		nextID:   1,
	}
}

func (j *JsonStore) save() {
	data, _ := json.MarshalIndent(j.contacts, "", " ")
	os.WriteFile("contacts.json", data, 0644)
}

func (j *JsonStore) Ajouter(c Contact) Contact {
	c.ID = j.nextID
	j.contacts[j.nextID] = c
	j.nextID++
	j.save()
	return c
}

func (j *JsonStore) Lister() []Contact {
	var list []Contact
	for _, c := range j.contacts {
		list = append(list, c)
	}
	return list
}

func (j *JsonStore) Supprimer(ID int) bool {
	if _, ok := j.contacts[ID]; ok {
		delete(j.contacts, ID)
		j.save()
		return true
	}
	return false
}

func (j *JsonStore) Recuperer(ID int) (Contact, bool) {
	c, ok := j.contacts[ID]
	return c, ok
}

func (j *JsonStore) MettreAJour(c Contact) (Contact, bool) {
	if _, ok := j.contacts[c.ID]; ok {
		j.contacts[c.ID] = c
		j.save()
		return c, true
	}
	return Contact{}, false
}

func (j *JsonStore) NextID() int {
	return j.nextID
}
