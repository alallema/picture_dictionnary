# Picture_dictionnary

This project is a search engine of picture and video for a mood board
The goal is to labelise pictures from choosen source download and reference them to be able to retrieve interesting data by word or mood

## Architecture

### Code Architecture

This project try to create a module that may  be changed easily for another solution
It's means that every modules has to use or translate his data by core(lib) data structure
At the end the goal is that every module could run like micro services
For each module:

*  Domain part is placed inside \<subproject\>/service
*  Dependency injection is done inside \<subproject\>/cmd/main.go
*  Mostly, the other directories inside subprojects are adapters

Picture Dictionnary follows the principle of **Monorepo** strategy, this means that all (micro)services running in production should be inside this repo Except Api and Redis Datastore. Let's take a look of each folder:

* In **api**, we have the service that allows us to make request to the picture-dictionnary
This is for now the only service to be independent

* In **vision-client**, we have an dev client that retreive all images informations request to google-cloud-vision

* In **redis-client**, we have an dev client that post and get all informations from redis datastore

* In **core**, there is useful structs or constant useful and shared between subproject

* In **app-client** we have an dev client for search with a basic simple UI

* In **collector** we have an dev client who goint to request all files from cloud-storage

* In **scraper**, scrap picture and video from differents sources. So they will be surely multiple scraper

**Docker** Have to be remade for now it's just launch redis datastore and api

### General architecture overview

![Architecture](docs/Architecture.jpg)

## Getting Started

### Instructions to launch Redis, Api and App

In order to bring up:

- Use `docker-compose build` to build api
- Use `docker-compose up` to see the logs of all the containers
- Use `docker-compose up -d` if you want it to run in the foreground
- Checking the installation with `redis-cli ping`
- Help `redis-cli --help`
- Launch redis-shell `docker exec -it redis redis-cli`
- See database `CONFIG GET databases`
- See keyspace `INFO keyspace`
- Visit http://localhost:4000/ to check api
- Visit http://localhost:3000/ to check the app-client
- In order to clean up the cluster, use `docker-compose down`

### Instructions to Launch Core and Collector

 - Setup env variable like `export BUCKETNAME="picture-dictionnary-bucket"`
 - Setup env variable like `export BUCKETNAMETOPROCESS="picture-dictionnary-to-process"`
 - Launch it `go run core/cmd/main.go`

## API

- GET /pulse — heartbeat check if our API is online
- GET /pictures(#get-pictures) — fetch all pictures from the database (Not implement yet)
- GET /picture/[tag] — fetch pictures by tag from the database
- GET /filteredtags/?key=tag1,tag2 - fetch pictures filtered by multiple tags
- GET /tags - fetch all labels and objects from database
- GET /labels - fetch all labels from database
- GET /objects - fetch all objects from database

## Redis Structure and API

- See README.md inside Api directory

## Launch Test

- make test

## To do

- create logger error - api done - service todo
- change general architecture schema

* Client api vision
    - check picture size 20 Mo
    - check video size
    - finish test implementation

* Api Go
    - change implementation Api has to request from Core/lib
    - change tag id by name in url request
    - swagger

* Create and Dockerise analyzer
    - Detection
    - Indexation
    - Search indexation

* Web App
    - fix authentification for url (Chrome cors)
    - create config in file and env
    - implement visualization video
    - drag and drop

* Deploiement
    - docker core/cmd + cron
    - docker scraper + cron
    - launch makefile in CI