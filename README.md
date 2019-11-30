# Picture_dictionnary

This project is search engine of picture and video for mood board
The goal is to labelise pictures from choosen source download and reference them to be able to retrieve interesting data by word or mood

## Architecture

Architecture schemas - Todo

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
- Check Api on http://localhost:8080/
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
        7) "Source"
        8) "instagram:Lala"
        9) "format"
        10) "jpg"
```

ZSET    key field value
- key is PictureId Field is mid zscore

```
KEY     19461332
FIELD   1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

SET     key member [member ...]
 - key is mid return all mid

```
KEY     Mid
MEMBER  1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

SET     key member [member ...]
 - key is mid, value is image id

```
KEY     Mid:/m/0bt9lr
MEMBER  1) "19461156"
        2) "19461153"
        3) "19461154"
        4) "19461155"
        5) "19461156"
```

HASH    key field value

```
KEY     midlang:/m/0bt9lr
FIELD   1) "fr"
        2) "chien"
        3) "en"
        4) "dog"
        5) "de"
        6) "hund"
        7) "es"
        8) "perro"
```

## API

- GET /pulse — heartbeat check if our API is online
- GET /pictures(#get-pictures) — fetch all pictures from the database
- GET /picture/[tag] — fetch pictures by tag from the database
- GET /tags - fetch all tags from database

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