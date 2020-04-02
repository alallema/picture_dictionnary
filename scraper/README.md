Simple vimeo scraper from following account

## Setup

```export GOOGLE_APPLICATION_CREDENTIALS="[PATH]"```

```npm install```
```node scraper-vimeo.js```
```node storage.js```

## Redis Structure

HASH    key field value

```
KEY     vimeo-userInfo:17331649
FIELD   1) "date"
        2) "20-01-2020"
        3) "total"
        4) "12"
```

```
KEY     vimeo-videoInfo:291048155
FIELD   1) "download-date"
        2) "20-01-2020"
        3) "directory"
        4) "vimeo/17331649"
```

SET     key member [member ...]
 - key is id return all id

```
KEY     vimeo-userIdList
MEMBER  1) "17331649"
        2) "7272496"
        3) "4939407"
        4) "1456592"
        5) "516907"
```

```
KEY     vimeo-userVideoList:17331649
MEMBER  1) "291048155"
        2) "206567344"
        3) "104819423"
        4) "toprocess-183497557"
        5) "toprocess-108611393"
```