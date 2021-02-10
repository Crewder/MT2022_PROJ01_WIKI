# gowiki-api

*Camille Arsac, Rémi Coufourier, Florian Leroy et Steven Nativel*

## Description du projet

Une API qui gère des wikis. On peut se connecter, créer un compte, créer un article, le modifier et mettre des commentaires.
## Pré requis

**GORM**
```
go get -u gorm.io/gorm  
go get -u gorm.io/driver/mysql
```

**Dépendance Router**
```
go get -u github.com/gorilla/mux
```

**Variables d'environnement**
```
go get github.com/joho/godotenv
```

## Lancement de l'application 
```
go run
```

## Table de contenu

# Requetes

| Méthodes |    Endpoint |Action|
|--|--|--|
|POST | api/article  | Création d'un article |
|PUT | api/article/{id} | Update d'un article  |
|GET | api/article/{id} | Récupération d'un article |
|GET | api/allarticle | Récupération de tout les articles |


| Méthodes |    Endpoint |Action|
|--|--|--|
|POST| api/comment/ | Création d'un commentaire |

| Méthodes |    Endpoint |Action|
|--|--|--|
|POST| api/user/ | Création d'un user |
|GET| api/user/ | Récupération des users |
|POST| api/auth/ | Connexion utilisateur |

# Article
## Création d'un article

## Update d'un article

## Récupération d'un article

## Récupération de tout les articles

# Commentaire
## Création d'un commentaire

# Utilisateur
## Création d'un user
## Récupération des users
## Connexion utilisateur


# Modèle de données



