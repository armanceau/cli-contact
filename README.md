# cli-contact

Un petit CRM en ligne de commande développé en Go pour gérer des contacts (ID, Nom, Email).  
Permet d’ajouter, lister, supprimer et mettre à jour des contacts directement depuis le terminal.  

## Fonctionnalités

- Afficher un menu interactif en boucle
- Ajouter un contact
- Lister tous les contacts
- Supprimer un contact par ID
- Mettre à jour un contact
- Quitter l’application
- Ajouter un contact via des flags sans passer par le menu

## Installation

1. Cloner le repository :

```bash
git clone https://github.com/armanceau/cli-contact.git
cd cli-contact
```

2. Lancer le projet :
```bash
go run main.go
```

## Utilisation
### Menu interactif 
_Lancer le programme. Puis suivre les instructions pour :_
1. Ajouter un contact
2. Lister les contacts
3. Supprimer un contact
4. Mettre à jour un contact
5. Quitter

### Ajouter un contact direct
_Il est aussi possible d'ajouter un contact à sa liste directment sans passer par le menu interactif via les flags._

Lancer le projet à l'aide des flags : 
```bash
go run main.go -nom "Arthur" -email "arthur@mail.com"
```

| Nom du flag | Type | Description |
|----------------|--------|------|
| `nom` | String | Nom du contact |
| `email` | String | Email du contact |

## Tests
_Des tests unitaires sont inclus pour vérifier l’ajout, la suppression et la mise à jour des contacts._

```bash
go test
```

## Auteur
Arthur Manceau 🙉


