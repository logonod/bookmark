const url = require('url');

exports.parse = async function(link) {
    link = link.trim();
    link = encodeURI(link);
    const u = url.parse(link);
    if (u.protocol !== 'https:' && u.protocol !== 'http:') {
        return [null, {errcode: 40400, errmsg: '仅支持收藏http和https协议的网页'}];
    }
    if (link.length > 1200) {
        return [null, {errcode: 40401, errmsg: '网址过长'}];
    }
    if (link.length === 0) {
        return [null, {errcode: 40402, errmsg: '网址不能为空'}];
    }
    if (u.host) {

    } else {
        return [null, {errcode: 40000, errmsg: '输入错误'}];
    }

    return [link, null]
};