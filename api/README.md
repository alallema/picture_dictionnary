Create simple api in go

## API Resources

- GET /pulse — heartbeat check if our API is online
- GET /pictures(#get-pictures) — fetch all pictures from the database (Not implement yet)
- GET /picture/[tag] — fetch pictures by tag from the database
- GET /tags - fetch all labels and objects from database
- GET /labels - fetch all labels from database
- GET /objects - fetch all objects from database

## Setup

Don't need it if you launch it with docker

* ```export GOPATH="${HOME}/.go"```
* ```export GOROOT="$(brew --prefix golang)/libexec"```
* ```cd GOPATH```
* ```go get github.com/go-redis/redis```
* ```go get github.com/gorilla/mux```
* ```go get github.com/rs/zerolog```
* ```go get github.com/caarlos0/env```
* ```go run main.go```

## Dependencies

* fmt
* time
* zerolog
* net/http
* go-redis
* gorilla
* caarlos0/env

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