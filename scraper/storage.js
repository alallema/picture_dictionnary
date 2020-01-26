const fs = require('fs');
const {Storage} = require('@google-cloud/storage');
const storage = new Storage();
const vidl = require('vimeo-downloader');
const myBucket = storage.bucket('picture-dictionnary-to-process/');
const removeVideoFromRedis = require('./redis-cli.js');
var redis = require('redis');

const host = 'localhost';
const port = 6379;

function downloadVideo(video) {
    let file = myBucket.file('vimeo/' + video.date + '/video' + video.videoId + 'user' + video.userId + '.mp4');

    vidl('https://vimeo.com/' + video.videoId, { quality: '360p' })
        .pipe(file.createWriteStream({
            metadata: {
              contentType: 'video/mp4',
              metadata: {
                source: 'vimeo'
              }
            }
        }))
        .on('error', function(err) {
            console.log(err)
            return false;
        })
        .on('finish', function() {
            console.log('upload', file.name);
            return true;
          // The file upload is complete.
        });
    return false;
};

function launchDownloader() {
    const options = {year: 'numeric', month: 'numeric', day: 'numeric' };
    const client = redis.createClient(port, host);
    var today  = new Date();

    client.on('error', function (err) {
        console.log('Error ' + err);
    });

    client.smembers('vimeo-userVideoList', function (err, result) {
        if (result[0]) {
            var video = new Object();
            video.Id = result[0];
            video.userId = result[0].split('-')[0];
            video.videoId = result[0].split('-')[1];
            video.date = today.toLocaleDateString('de-DE', options);
            downloadVideo(video)
            removeVideoFromRedis(video)
            console.log(video)
        }
    })
    client.quit();
}

launchDownloader();