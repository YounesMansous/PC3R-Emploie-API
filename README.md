Binôme:

- Touré-Ydaou TEOURI
- Younes MANSOUS

# Traffik : une application web informant sur les pertubations le réseau IDFM (Ile de France Mobilité)

Traffik est une application permettant de connaitre les pertubations sur le réseau de transport de l'ile de france.

# Choix de l'API web

Nous avons choisi d'utiliser l'API [d'ile de france mobilité](https://prim.iledefrance-mobilites.fr/fr/apis/idfm-navitia-line_reports-v2) qui donne les messages d'informations sur le trafic en temps réel sur l'ensemble de ses transports. Cette APi est gratuite, elle permet de faire 4000 requêtes par jour. Pour une requête elle donne la possibilité de faire une recherche d'informations selon :

- le mode de transport
- l'identifiant de la ligne

Notre application permet de rechercher les messages de pertubations en fonction de la ligne de transport choisie par l'utilisateur.

Pour chaque pertubation l'application donne : - le nom de la ligne - le mode de transport affecté - le message d'avertissement - la date de publication

# Fonctionnalités de l'application

- L'application récupère tous les messages de pertubations chaque jour

- Les informations sont affichés selon la ligne choisie dans un flux d'information similaire à Facebook

# Cas d'utilisation

- Mocktar s'inscrit et se connecte à l'application, une page d'accueil vide lui propose le mode de transport, puis la ligne dont il veut consulter les informations. Après avoir réalisé sont choix il voit les informations de la ligne correspondante.

- Julie se connecte à l'application après avoir affiché les informations de sa ligne quotidienne, elle voit une information qui attire son interêt et décide d'y ajouter un commentaire.

- Marc n'est pas connecté il peut choisir de lire les informations d'une ligne particulière mais ne peut pas ajouter de commentaires.

# Base de données

## SQL

Pour la base de données nous avons utilisé PostgreSQL qui est une base de donnée relationelle. Celle-ci est hébergée sur Supabase qui nous fourni un api afin de communiquer avec notre base de données et d'effectuer des opérations basiques (ajout, mise à jour, lecture et suppresion) sur les données.

#### Lines

Il s'agit de la table contenant les informations sur les lignes de transport d'ile de france mobilité.

| id          | name | type   |
| ----------- | ---- | ------ |
| IDFM:C01380 | 4    | Subway |

Les identifiants des lignes doivent respecter un format précis fourni par IDFM, pour cela ceux-ci on été récupérés depuis leur jeux de données et insérés directement en base.

#### Events

Cette table contient l'ensemble des évenements concernant l'actualité du trafic d'IDFM, chaque évenements est lié à une ligne de transport donnée.

| id  | title             | content           | line_id (clé étrangère) | published_date  | created_date    |
| --- | ----------------- | ----------------- | ----------------------- | --------------- | --------------- |
| 1   | Metro 5 : Travaux | <p>Lorem Ipsum<p> | IDFM:C01380             | 20250318T220307 | 20250318T220307 |

#### Comments

Cette table contient les commentaires des utilisateurs de la plateforme. Chaque commentaire est rattaché à un message de pertubation et à un utilisateur donné.

| id  | content     | event_id (clé étrangère) | user_id (clé étrangère) | created_date    | created_date    |
| --- | ----------- | ------------------------ | ----------------------- | --------------- | --------------- |
| 1   | Lorem Ipsum | 85                       | 525                     | 20250318T220307 | 20250318T220307 |

#### Users

Cette table contient les informations des utilisateurs, chaque email contenu dans cette table est unique.

| id  | name     | email             | password                                                       | created_at      |
| --- | -------- | ----------------- | -------------------------------------------------------------- | --------------- |
| 1   | John Doe | johndoe@email.com | \$2a\$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK | 20250318T220307 |

Les données sont liées entre elles avec par les clés étrangères. Nous avons établi des règles sur celles-ci :

- Sur la règle `ON DELETE`, nous avons choisi `CASCADE`, ce qui permet lors de la suppression d'une donnée de supprimer automatique toutes les données qui y font référence.
  - Par exemple : lors de la suppression d'un évenement tous les commentaires associés sont automatiquement supprimés
- Sur la règle `ON UPDATE`, nous avons choisi `CASCADE`, ce qui permet lorque l'identifiant d'une ressource est mise à jour, alors cet identifiant est mis à jour pour toutes les données y faisant référence.
  - Par exemple : si l'identifiant de la ligne de métro 8 change alors tous les évenements y faisant référence veront leur attribut `line_id` mis à jour

# Mise à jour des données

Pour la mise à jour des données nous faisons appel à l'API externe **tous les jours à 8h00** et sauvegardons les nouveaux évenements concernants chaque lignes. Avant de réaliser cette opérations les évenements précedent sont supprimés.

# Serveur

Nous avons implémentés une **API REST** et avons optés pour une approche **ressource**, notre API étant principalement centrée sur des entitées. Voici les composants permettant de manipuler nos ressources :

- Authentification
  - Connexion :
- Users
  - Inscription : création d'un nouvel utilisateur
- Lines:
  - Modes : récuperer les modes de transports
  - Identifiants : récuperer les identifiants des lignes en fonction du mode de transport
- Evenements
  - Flux informations : récupère les évenements de la ligne choisie par l'utilisateur
  - Détails evenement : récupère un évenement particulier
- Comments
  - Commenter : l'utilisateur commente un évenement
  - Flux commentaires : récupère les commentaires d'un évenement

## Endpoints de l'API

Voici une description des différents endpoints de notre API.

- `/auth`: composant _authentification_
  - `POST /login` : connexion d'un utilisateur
  - `POST /register` : création d'un nouvel utilisateur
- `/lines` : composant _lines_
  - `GET /lines/modes` : récupère les modes de transport d'IDFM
  - `GET /lines/modes/id?mode=bus` : rècupère les identifiants de toutes les lignes d'un mode de transport donné
- `/events` : composant _events_
  - `GET /events/line?id_line=IDFM::CO1352` : récupère l'ensemble des évenements concernant une ligne donnée
  - `GET /events?id=52` : récupère un évenement donné en fonction de son identifiant
- `/comments` : composant _comments_
  - `POST /comments/add?event_id=20` : permet à l'utilisateur de rajouter un commentaire **ici l'authentification est obligatoire**
  - `GET /comments?event_id=20` : permet de récupérer les commentaires d'un évenement
