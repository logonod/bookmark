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

exports.get = async function(user_id, url_hash) {
    let r, strBody, charset, tmpCharset, bodyDecoded;
    r = await to(request({
        url: "http://127.0.0.1:8001/api/spider/collect/get",
        headers: {
            'Secret-Key': 'amVmZmdlZWs='
        },
        timeout: 5000,
        encoding: null,
        strictSSL: false,
        json: {"user_id": user_id, "url_hash": url_hash}
    }));
    if (r[1]) {
        // http get error
        return [null, {errcode: 40403, errmsg: '请求失败'}];
    } else {
        if (r[0].statusCode !== 200) {
            return [null, {errcode: 40405, errmsg: '状态码错误'}];
        }
        // http success
        return [r[0].body, null];
    }
};