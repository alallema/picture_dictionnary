const axios = require('axios');
const cheerio = require('cheerio');
const puppeteer = require("puppeteer");
const pushList = require('./redis-cli.js');

const url = 'https://vimeo.com/user98620336/following';

getUsers = async () => {
let users = await axios(url)
  .then(response => {
    const html = response.data;
    const $ = cheerio.load(html);
    const paginationList = $('.pagination > ol > li > a');
    var pageNumber = 1;

    paginationList.each(function () {
      if (parseInt($(this).attr('data-page'), 10) > pageNumber)
        pageNumber = parseInt($(this).attr('data-page'), 10);
    });
    return pageNumber
  })
  .then(async value => {
    console.log('Number of pages to be scrapped: ', value);
    let data = [];
    for (i = 1; i < value; i++) {
      urlPage = 'https://vimeo.com/user98620336/following' + '/page:' + i + '/sort:date';
      let idList = [];
      await axios(urlPage)
        .then(response => {
          const html = response.data;
          const $ = cheerio.load(html);
          const browseList = $('.js-browse_list > li');

          browseList.each(function () {
            idList.push($(this).attr('id'));
          });
          return idList;
        })
        .catch(console.error);
        data.push(idList);
    }
    return data;
  })
  .catch(console.error);
  return users;
};

const getVideoId = async (url) => {
    const browser = await puppeteer.launch()
    const page = await browser.newPage()
  
    await page.goto(url)
    await page.waitFor(1000)
    // await page.click('#wrap > div.wrap_content.variant-v2 > main > div > div.profile_main.greyed_out_false > div > section.iris_p_content__main.sc-cSHVUG.bKlEbK > div.sc-frDJqD.bJHDuz > div > div.sc-caSCKo.bcVbzQ > div > button > span > span');

    const result = await page.evaluate(() => {
      let data = [];
      let elements = $('.iris_video-vital__overlay').toArray();
      for (var element of elements)
          data.push(element.getAttribute('href'));
      return { data }
    })
  
    browser.close()
    return result
}

getUsers().then(async users => {
    let videoList = [];
    for (i = 0; i < users.length; i++) {
        for (j = 0; j < users[i].length; j++) {
            let videoId = await getVideoId('https://vimeo.com/user' + users[i][j].substring(5))
            .then(value => {
                return value;
            })
            .catch(console.error);
            for (k = 0; k < videoId.data.length; k++) {
                var video = new Object();
                video.userId = users[i][j].substring(5)
                video.videoId = videoId.data[k];
                videoList.push(video);
            }
        }
    }
    console.log('Get all videos list: ', videoList);
    return (videoList);
})
.then(value => {
    pushList(value)
});