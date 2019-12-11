# Picture_dictionnary

This project is search engine of picture and video for mood board
The goal is to labelise pictures from choosen source download and reference them to be able to retrieve interesting data by word or mood

## Architecture

### Code Architecture

This projects try to create module that can be changed easily for another solution
It's means that every modules has to use or translate his data by core(lib) data structure
At the end the goal is that every module could run like micro services
For each module:

*  Domain part is placed inside \<subproject\>/service
*  Dependency injection is done inside \<subproject\>/cmd/main.go
*  Mostly, the other directories inside subprojects are adapters

Picture Dictionnary follows the principle of **Monorepo** strategy, this means that all (micro)services running in production should be inside this repo Except Api and Redis Datastore. Let's take a look of each folder:

* In **api**, we have the service that allows us to make request to the picture-dictionnary

* In **vision-client**, we have an dev client that retreive all images informations request to google-cloud-vision

* In **redis-client**, we have an dev client that post and get all informations from redis datastore

* In **core**, there is useful structs or constant useful and shared between subproject

**Docker** Have to be remade for now it's just launch redis datastore and api

### General architecture overview

![Architecture](docs/Architecture.jpg)

## Getting Started

### Instructions to launch Redis and Api

In order to bring up:
- Use `docker-compose build` to build api
- Use `docker-compose up` to see the logs of all the containers
- Use `docker-compose up -d` if you want it to run in the foreground
- Checking the installation with `redis-cli ping`
- Help `redis-cli --help`
- Launch redis-shell `docker exec -it redis redis-cli`
- See database `CONFIG GET databases`
- See keyspace `INFO keyspace`
- Visit http://localhost:8080/ to check api
- In order to clean up the cluster, use `docker-compose down`

## Redis Structure

HASH    key field value
- Picture data

```
KEY     picture:"98asdjw2"
FIELD   1) "id"
        2) "19461154"
        3) "title"
        4) "img_657.jpg"
        5) "path"
        6) "/path/img_657.jpg"
        7) "source"
        8) "instagram:Lala"
        9) "format"
        10) "jpg"
```

ZSET    key field value
- key is PictureId Field is mid zscore

```
KEY     pictureIdLabel:19461332
FIELD   1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

ZSET    key field value
- key is PictureId Field is mid zscore

```
KEY     pictureIdObject:19461332
FIELD   1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

SET     key member [member ...]
 - key is mid return all mid

```
KEY     labelId
MEMBER  1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

SET     key member [member ...]
 - key is mid return all mid

```
KEY     objectId
MEMBER  1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

SET     key member [member ...]
 - key is mid, value is image id

```
KEY     Id:/m/0bt9lr
MEMBER  1) "19461156"
        2) "19461153"
        3) "19461154"
        4) "19461155"
        5) "19461156"
```

HASH    key field value

```
KEY     labelDescr:/m/0bt9lr
FIELD   1) "en"
        2) "dog"
        3) "fr"
        4) "chien"
        5) "de"
        6) "hund"
        7) "es"
        8) "perro"
```

```
KEY     objectDescr:/m/0bt9lr
FIELD   1) "en"
        2) "dog"
        3) "fr"
        4) "chien"
        5) "de"
        6) "hund"
        7) "es"
        8) "perro"
```

## API

- GET /pulse — heartbeat check if our API is online
- GET /pictures(#get-pictures) — fetch all pictures from the database (Not implement yet)
- GET /picture/[tag] — fetch pictures by tag from the database
- GET /tags - fetch all labels and objects from database
- GET /labels - fetch all labels from database
- GET /objects - fetch all objects from database

## To do

- Dockerise environment
- create logger error - logger history ?

* Check
    - check picture size
    - encodage base64 ?

* Client api vision
    - send picture request
    - get resp from api
    - stock response
    - see how it's for video

* Api Go
    - create clean function to get data from db
    - create function to get request

* Database
    - launch db
    - organize it

* Web app
    - ???

Labeling pictures
Labeling Video
Download instagram Vimeo
Check instagram api

* Create and Dockerise analyzer
    - Detection
    - Indexation
    - Search indexation