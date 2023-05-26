const cheerio = require('cheerio');

exports.parse = async function(response) {
    let linkEntity = {
        type: null,
        url: null,
        title: null,
        encoding: null
    };

    try {
        let $ = cheerio.load(response.data.content);
        let title = $("title").text().trim() || response.link;
        linkEntity.type = response.link.type;
        linkEntity.url = response.link;
        linkEntity.title = title.trim();
        linkEntity.encoding = response.data.encoding;

        return [linkEntity, null];
    } catch(error) {
        console.error(error);
        return [null, {errcode: 40405, errmsg: '解析网页失败'}];
    }
};