const to = require('./to');
const bluebirdPromise = require("bluebird");
const request = bluebirdPromise.promisify(require("request").get);
const typeis = require('type-is');
const iconvLite = require('iconv-lite')

function get(res, field) {
    return res.headers[field.toLowerCase()] || '';
}

function contentType(res) {
    return get(res, 'content-type').split(';').filter(item => item.trim().length !== 0).join(';');
}

function getHeaderCharset(str) {
    if (str == null) return null;
    var charset = str.match(/charset=["']?([^>"'\s]+)/i);
    if (charset instanceof Array && charset.length >= 2) return charset[1];

    return null;
}

function getBodyCharset(str) {
    if (str == null) return null;
    var charset = str.match(/<meta\s*http-equiv=["']?Content-Type["']?.*charset=["']?([^>"'\s]+)[\s\S]*>/i);
    if (charset instanceof Array && charset.length >= 2) return charset[1];
    charset = str.match(/<meta\s*charset=["']?([^>"\s]+)[\s\S]*>/i);
    if (charset instanceof Array && charset.length >= 2) return charset[1];
    charset = str.match(/<meta\s*.*charset=["']?([^>"'\s]+)[\s\S]*>/i);
    if (charset instanceof Array && charset.length >= 2) return charset[1];

    return null;
}

exports.get = async function(linkUrl) {
    const response = {
        link: null,
        data: null
    };
    response.link = linkUrl;
    let r, strBody, charset, tmpCharset, bodyDecoded;
    r = await to(request({
        url: linkUrl,
        headers: {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36'
        },
        timeout: 5000,
        encoding: null,
        strictSSL: false
    }));
    if (r[1]) {
        // http get error
        return [null, {errcode: 40403, errmsg: '请求失败'}];
    } else {
        if (r[0].statusCode !== 200) {
            return [null, {errcode: 40405, errmsg: '状态码错误'}];
        }
        // http success
        if (typeis.is(r[0], ['html'])) {
            // is html
            strBody = r[0].body instanceof Buffer ? r[0].body.toString() : r[0].body;
            charset = getHeaderCharset(contentType(r[0]));
            tmpCharset = getBodyCharset(strBody);
            if (tmpCharset) {
                charset = tmpCharset;
            }
            charset = charset || 'utf-8';
            bodyDecoded = iconvLite.decode(r[0].body, charset);
            response.data = {
                'encoding': charset,
                'content': bodyDecoded
            };
            return [response, null];
        } else {
            // not html
            return [null, {errcode: 40404, errmsg: '仅支持收藏网页'}];
        }
    }
};