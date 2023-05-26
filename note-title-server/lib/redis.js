var redis = require("redis");
const to = require('./to');
const bluebirdPromise = require("bluebird");

var client = redis.createClient(6379, '127.0.0.1');
var get = bluebirdPromise.promisify(client.get).bind(client);

client.on('error', function(err){ 
	console.error('Redis error:', err); 
});

exports.checkAuth = async function(cookie) {
    r = await to(get(cookie));
    if (r[1]) {
    	return [null, {errcode: 40002, errmsg: '请稍后重试'}];
    }

    return [r[0], null]
};
