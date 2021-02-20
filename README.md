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
|POST| comment/create | Création d'un commentaire |

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

#### Request Body
```
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

#### Request Url
```
http://localhost:8080/article/1/view
```
#### Request Response
```
{
    "ID": 1,
    "CreatedAt": "2021-02-15T22:25:34+01:00",
    "UpdatedAt": "2021-02-15T22:25:37+01:00",
    "DeletedAt": null,
    "UserId": 0,
    "User": {
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "ID": 0,
        "Name": "",
        "Email": "",
        "Password": ""
    },
    "Title": "J'aime les chats",
    "Content": "La civilisation aztèque est une civilisation d’Amérique Centrale basée dans la vallée de Mexico dès le début du XIVème siècle. Les Aztèques seront délogés par les conquistadors aux alentours de 1519. Entre-temps, ils avaient eu le temps d’atteindre un des niveaux de civilisation les plus avancés d’Amérique. Tout comme chez les Mayas, le système de croyances aztèque imposait des sacrifices humains pour les dieux. Apprenez-en plus grâce à cet article.\nPourquoi les Aztèques pratiquaient les sacrifices humains ?\n\nLes Aztèques pensaient que le sang humain était la principale nourriture des dieux, ainsi il était pour eux, tout à fait normal de pratiquer le sacrifice humain afin de s’attirer les bonnes grâces des divinités. Si les sacrifices humains avaient une fonction religieuse dans la civilisation aztèque, ils avaient également une fonction politique.\nLes sacrifices humains et la religion aztèques\n\nD’après les croyances aztèques, les sacrifices humains étaient des éléments indispensables au bon fonctionnement et à l’équilibre de l’univers. La première référence à ces pratiques se trouve dans le mythe de la création du monde. Dans celui-ci, la déesse-terre, Tlaltecuhtli, exige des sacrifices humains et refuse même d’apporter ses bienfaits à moins d’être arrosée de sang. Par la suite, deux dieux, Nanahuatzin et Tecciztecatl sont sacrifiés pour pouvoir renaître sous la forme du Soleil et de la Lune. D’autres sacrifices sont indispensables pour que le Soleil entame sa course autour de la Terre.\n\nDans la Légende des soleils, on raconte que la déesse-Terre avait donné naissance à 400 Mimixcoas, un type de dieu, et à 5 Mecitin, c’est-à-dire des humains. Tandis que les dieux s’adonnaient régulièrement à la luxure et à la fête, ils ne permettaient pas de nourrir la Terre et le Soleil. Les 5 humains furent donc chargés de les tuer afin d’utiliser leur sang pour nourrir les dieux supérieurs. On raconte également que tous les mondes dans lesquels les humains ne pratiquaient pas de sacrifices avaient été détruits par les dieux. On sacrifiait donc régulièrement des humains afin d’apaiser la colère des divinités."
}
```

**GET** - recuperation de tous les articles

#### Request Url
```
http://localhost:8080/articles/view
```
#### Request Response
```
[
   
    {
        "ID": 1,
        "CreatedAt": "2021-02-15T22:25:34+01:00",
        "UpdatedAt": "2021-02-15T22:25:37+01:00",
        "DeletedAt": null,
        "UserId": 0,
        "User": {
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "ID": 0,
            "Name": "",
            "Email": "",
            "Password": ""
        },
        "Title": "J'aime les chats",
        "Content": "La civilisation aztèque est une civilisation d’Amérique Centrale basée dans la vallée de Mexico dès le début du XIVème siècle. Les Aztèques seront délogés par les conquistadors aux alentours de 1519. Entre-temps, ils avaient eu le temps d’atteindre un des niveaux de civilisation les plus avancés d’Amérique. Tout comme chez les Mayas, le système de croyances aztèque imposait des sacrifices humains pour les dieux. Apprenez-en plus grâce à cet article.\nPourquoi les Aztèques pratiquaient les sacrifices humains ?\n\nLes Aztèques pensaient que le sang humain était la principale nourriture des dieux, ainsi il était pour eux, tout à fait normal de pratiquer le sacrifice humain afin de s’attirer les bonnes grâces des divinités. Si les sacrifices humains avaient une fonction religieuse dans la civilisation aztèque, ils avaient également une fonction politique.\nLes sacrifices humains et la religion aztèques\n\nD’après les croyances aztèques, les sacrifices humains étaient des éléments indispensables au bon fonctionnement et à l’équilibre de l’univers. La première référence à ces pratiques se trouve dans le mythe de la création du monde. Dans celui-ci, la déesse-terre, Tlaltecuhtli, exige des sacrifices humains et refuse même d’apporter ses bienfaits à moins d’être arrosée de sang. Par la suite, deux dieux, Nanahuatzin et Tecciztecatl sont sacrifiés pour pouvoir renaître sous la forme du Soleil et de la Lune. D’autres sacrifices sont indispensables pour que le Soleil entame sa course autour de la Terre.\n\nDans la Légende des soleils, on raconte que la déesse-Terre avait donné naissance à 400 Mimixcoas, un type de dieu, et à 5 Mecitin, c’est-à-dire des humains. Tandis que les dieux s’adonnaient régulièrement à la luxure et à la fête, ils ne permettaient pas de nourrir la Terre et le Soleil. Les 5 humains furent donc chargés de les tuer afin d’utiliser leur sang pour nourrir les dieux supérieurs. On raconte également que tous les mondes dans lesquels les humains ne pratiquaient pas de sacrifices avaient été détruits par les dieux. On sacrifiait donc régulièrement des humains afin d’apaiser la colère des divinités."
    },
    {
        "ID": 2,
        "CreatedAt": "2021-02-15T23:05:52+01:00",
        "UpdatedAt": "2021-02-15T23:05:55+01:00",
        "DeletedAt": null,
        "UserId": 0,
        "User": {
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "ID": 0,
            "Name": "",
            "Email": "",
            "Password": ""
        },
        "Title": "J'aime les lapins",
        "Content": "orem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
    }
]
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
#### Request Url
```
http://localhost:8080/comment/create
```
#### Request Body
```
{
    "UserId": 1,
    "ArticleId": 2,
    "comment": "J'aime les pistacles"
}
```
#### Request Response
```
null
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



