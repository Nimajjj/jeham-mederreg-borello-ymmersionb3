# Documentation de l'API Go Pebbles
Cette documentation décrit les routes, modèles de données et fonctionnalités de l'API Go développée pour la gestion d'un stock de cailloux (Pebbles).

# Sommaire
- [[#Routes]]
- [[#Modèles de données]]
- [[#Base de données]]
---

# Routes
`/pebble/:id`
Methode: `GET`
Paramètre: id (int) 
Récupère les détails d'un caillou par son id
> e.g.  `/pebble/31`

`/pebbles/:categories/:order/:keywords`
Methode: `GET`
Paramètres:
  - categories (string) : catégories filtrées
	  - Valeur par default: `[]`
  - order (string) : tri appliqué
	  - Valeur par default: `nil`
	  - Valeurs possible: `price_asc`, `price_desc`
  - keywords (string) : mots-clés de recherche
	  - Valeur par default: `nil`

Recherche des cailloux selon les filtres
> e.g. `/pebbles/["historique", "precieux"]/price_asc/"diamant"`
> e.g. `/pebbles/["historique", "precieux"]/nil/"nil"`
> e.g. `/pebbles/[]/price_desc/nil`

`/categories/`
Methode: `GET`  
Récupère la liste de toutes les catégories

`/cart/:user`
Methode: `GET`
Paramètre: user (int)
Récupère le panier d'un utilisateur
> e.g.  `/cart/2`

`/cart/add/:user/:pebble/:quantity`
Methode: `GET`
Paramètres:
  - user (int)
  - pebble (int) 
  - quantity (int)
Ajoute un nombre de caillou `quantity` au panier d'un utilisateur
Si la quantité est négative cela les retire du panier.
Si après modification la quantité d'un article est égale ou inférieure à 0 alors l'item est automatiquement retiré du panier. 
> e.g. `/cart/add/1/31/2`
> e.g. `/cart/add/1/31/-1`

# Modèles de données
## Pebble
```go
type Pebble struct {
  ID          int     
  Title       string   
  Description string  
  Price       float64  
  Breed       string   
  Quantity    int      
  Weight      float64
  Categories  []string 
  Photos      []string

  Creation    string
}
```

## Cart
```go
type Cart struct {
  ID int
  UserID int
  Content map[*Pebble]int
}
```

## Categorie
```go
type Categorie struct {
  ID int
  Title string
}
```

# Base de données
**Veuillez fournir les informations d'acces a la base de donnés dans le fichier `.env` se trouvant a la racine du dossier `backend`.**

Les interactions avec la base de données sont gérées dans le package repo.
PebbleRepo

Permet d'insérer, updater, supprimer des Pebbles et effectuer des recherches.
CartRepo

Gère les opérations sur les paniers : récupération, ajout d'articles, mise à jour.
CategorieRepo

Gestion des catégories : insertion, lien avec les Pebbles, récupération.
Fichiers

Ces fichiers contiennent les détails de l'implémentation :

    pebbles_repo.go
    cart_repo.go
    categorie_repo.go
    photo_repo.go
    routes.go
    pebble_route.go
    cart_route.go
    cat_route.go
