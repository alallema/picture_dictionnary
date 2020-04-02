var redis = require('redis');

const host = 'localhost';
const port = 6379;

function groupBy(tableauObjets, propriete){
    return tableauObjets.reduce(function (acc, obj) {
        var cle = obj[propriete];
        if(!acc[cle]){
          acc[cle] = [];
        }
        acc[cle].push(obj.videoId);
        return acc;
    }, {});
}

function putHashCallBack(result, key, value) {
    const clienthset = redis.createClient(port, host);

    if (result && result.total === value.length.toString()) {
        console.log('Already in redis');
    }
    else if (result && result.total != value.length.toString()) {
        console.log('Already in redis but not same total');
        clienthset.hset(key, 'total', value.length.toString(), redis.print)
    }
    else {
        console.log('Not in redis yet ' + key);
        clienthset.hmset(key, {
            'date': Date().toString(),
            'total': value.length.toString(),
        });
    }
    clienthset.quit();
}

function addVideo(key, value, videoId) {
    const clientsadd = redis.createClient(port, host);

    clientsadd.sadd('vimeo-userVideoList:' + `${key}`, 'toprocess-' + value[videoId].substring(1))
    clientsadd.sadd('vimeo-userVideoList', `${key}` + '-' + value[videoId].substring(1))
    clientsadd.quit();
}

module.exports = function pushList(List) {
    const list = groupBy(List, 'userId');

    const client = redis.createClient(port, host);

    client.on('error', function (err) {
        console.log('Error ' + err);
    });

    for (let [key, value] of Object.entries(list)) {
        client.hgetall('vimeo-userInfo:' + `${key}`, function (err, result) {
            putHashCallBack(result, 'vimeo-userInfo:' + key, value);
        });
        client.sadd('vimeo-userIdList', `${key}`, redis.print)
        client.smembers('vimeo-userVideoList:' + `${key}`, function (err, result) {
            for (videoId in value) {
                if ((result.includes(value[videoId].substring(1))) == false
                    && (result.includes('toprocess-' + value[videoId].substring(1)) == false)) {
                        addVideo(key, value, videoId);
                }
            }
        })
    }
    client.quit();
};

function createVideoInfo(video) {

    const client = redis.createClient(port, host);

    client.on('error', function (err) {
        console.log('Error ' + err);
    });

    client.hmset('vimeo-videoInfo:' + video.videoId, {
        'download-date': video.date,
        'user': video.userId,
        'directory': 'vimeo/' + video.date,
    });
    console.log("create video info in redis");
    client.quit();
}

module.exports = function removeVideoFromRedis(video) {

    const client = redis.createClient(port, host);

    client.on('error', function (err) {
        console.log('Error ' + err);
    });

    client.srem('vimeo-userVideoList', video.Id);
    client.srem('vimeo-userVideoList:' + video.userId, 'toprocess-' + video.videoId)
    client.sadd('vimeo-userVideoList:' + video.userId, video.videoId)
    console.log("delete old info")
    createVideoInfo(video);
    client.quit();
}
