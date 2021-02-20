# gowiki-api

*Camille Arsac, Rémi Coufourier, Florian Leroy et Steven Nativel*

<a name="description"/>

## 1 - Description du projet

Une API qui gère des wikis. On peut se connecter, créer un compte, créer un article, le modifier et mettre des commentaires.

### Table de contenu

* [1. Description](#description)
   * [1.1 Pré requis](#required)
   * [1.2 Lancement du projet](#launch)
* [2. Requete](#request)
* [3. Article](#article)
   * [3.1 Création d'un article](#createarticle)
   * [3.2 Mise a jour d'un article](#updatearticle)
   * [3.3 Récupération d'un article](#fetcharticle)
   * [3.4 Récupération de tout les articles](#fetchallarticle)
* [4. Commentaire](#comment)
    * [4.1 Création d'un commentaire](#createcomment)
* [5. Utilisateur](#user)
    * [5.1 Creation d'un utilisateur](#createuser)
    * [5.2 Récuperation des utilisateurs](#fetchuser)
    * [5.3 Connexion utilisateur](#Auth)
* [5. Modèle de données ](#models)
    

<a name="required"/>

### Pré requis

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
<a name="launch"/>

## Lancement de l'application 
```
go run
```



<a name="request"/>

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

<a name="article"/>

# Article

<a name="createarticle"/>

## Création d'un article

**POST** - Création d'un article

#### Parameters

| Name |    Type |Description|
|--|--|--|
|id| int| Id de l'article |
|created_at| date| Date de création |
|updated_at| date| date de mise à jour |
|deleted_at| date| date de suppretion |
|user_id| int| Id de l'utilisateur actif. Required |
|title| string| titre de l'article. Required |
|content| string| contenu de l'article. Required |

#### Request URL

```
http://localhost/article
```

#### Request Body
```
{
    "UserId": 1,
    "title": "Vache",
    "content":"Vache est le nom vernaculaire donné à la femelle du mammifère domestique de l'espèce Bos taurus, un ruminant appartenant à la famille des bovidés, généralement porteur de deux cornes sur le front. Les individus mâles sont appelés taureaux et les jeunes, veaux. Une génisse ou vachette, appelée aussi taure au Québec ou dans le Poitou, est une vache qui n'a pas vêlé. Descendant de plusieurs sous-espèces d'aurochs, les bovins actuels (zébus compris) sont élevés pour produire du lait et de la viande, ou comme animaux de trait. En Inde, la vache est sacrée. Le mot vache vient du latin vacca, de même sens."
}
```
#### Request Response
```
```

<a name="updatearticle"/>

## Update d'un article

**PUT** - Update de l'article

#### Parameters

| Name |    Type |Description|
|--|--|--|
|Article_id| int| Id de l'article. Required |


#### Request Body
```
```
#### Request Response
```
```
<a name="fetcharticle"/>

## Récupération d'un article

**GET** - recuperation d'un article

#### Parameters

| Name |    Type |Description|
|--|--|--|
|Article_id| int| Id de l'article. Required |

#### Request Body
```
```
#### Request Response
```
```

<a name="fetchallarticle"/>

## Récupération de tout les articles

**GET** - Récuperation de tout les articles

#### Request Body
```
```
#### Request Response
```
```

<a name="comment"/>

# Commentaire

<a name="createcomment"/>

## Création d'un commentaire

**POST** - Création d'un commentaire

#### Request Body
```
```
#### Request Response
```
```

<a name="user"/>

# Utilisateur

<a name="createuser"/>

## Création d'un user

**POST** - Création d'un User

#### Request Body
```json
{
    "Email":"string",
    "Username":"string"
}
```
#### Request Response
```json
{
    "Email":"string",
    "Username":"string"
}
```

<a name="fetchuser"/>

## Récupération des users
#### Request Body
```
```
#### Request Response
```
```

<a name="Auth"/>

## Connexion utilisateur

**POST** - Connexion d'un utilisateur

#### Request Body
```json
{
    "Email":"string",
    "Username":"string"
}
```
#### Request Response
```
```

<a name="models"/>


# Modèle de données



