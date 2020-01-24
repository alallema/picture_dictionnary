Create simple api in go

## API Resources

- GET /pulse — heartbeat check if our API is online
- GET /filteredtags/?key=tag1,tag2 - fetch pictures and video filtered by multiple tags
- GET /picture/[tag] — fetch pictures by tag from the database
- GET /video/[tag] — fetch pictures by tag from the database
- GET /tags - fetch all labels and objects from database
- GET /labels - fetch all labels from database
- GET /objects - fetch all objects from database
- GET /categories - fetch all categories from database

## Setup

Don't need it if you launch it with docker

* ```export GOPATH="${HOME}/.go"```
* ```export GOROOT="$(brew --prefix golang)/libexec"```
* ```cd GOPATH```
* ```go get github.com/go-redis/redis```
* ```go get github.com/gorilla/mux```
* ```go get github.com/rs/zerolog```
* ```go get github.com/caarlos0/env```
* [install and launch redis](https://redis.io/topics/quickstart)
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

```
KEY     picture:"98asdjw2"
FIELD   1) "id"
        2) "19461154"
        3) "title"
        4) "img_657.jpg"
        5) "picturePath"
        6) "/path/img_657.jpg"
        7) "source"
        8) "instagram:Lala"
        9) "format"
        10) "jpg"
        11) "pictureURL"
        12) "https://url/path/img_657.jpg"
```

```
KEY     labelDescr:/m/0bt9lr
FIELD   1) "en"
        2) "song"
        3) "fr"
        4) "chanson"
        5) "de"
        6) "song"
        7) "es"
        8) "canción"
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

```
KEY     categoryDescr:/m/0bt9lr
FIELD   1) "en"
        2) "hair"
        3) "fr"
        4) "cheveux"
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

```
KEY     pictureIdObject:19461332
FIELD   1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

```
KEY     seg:19461156:/m/0bt9lr
MEMBER  1) "2.416666s-3.25s"
        2) "875ms-2.375s"
        3) "2.416666s-3.25s"
```

SET     key member [member ...]
 - key is id return all id

```
KEY     labelId
MEMBER  1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

```
KEY     objectId
MEMBER  1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

```
KEY     categoryId
MEMBER  1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

```
KEY     pictureIdCategory:19461332
MEMBER  1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

```
KEY     categoryId:/m/012yh1
MEMBER  1) "/m/0bt9lr"
        2) "/m/01pm38"
        3) "/m/04rky"
        4) "/m/02wbgd"
        5) "/m/09686"
```

```
KEY     Id:/m/0bt9lr
MEMBER  1) "19461156"
        2) "19461153"
        3) "19461154"
        4) "19461155"
        5) "19461156"
```

```
KEY     VidId:/m/0bt9lr
MEMBER  1) "19461156"
        2) "19461153"
        3) "19461154"
        4) "19461155"
        5) "19461156"
```

```
KEY     URLId:/m/0bt9lr
MEMBER  1) "https://url/path/img_657.jpg"
        2) "https://url/path/img_658.jpg"
        3) "https://url/path/img_659.jpg"
        4) "https://url/path/img_660.jpg"
        5) "https://url/path/img_661.jpg"
```
