const cluster = require('cluster')
const os = require('os')
const link = require('./lib/link');
const http = require('./lib/http');
const html = require('./lib/html');
const log = require('./lib/log');
const redis = require('./lib/redis');
const express = require('express');
const cookieParser = require('cookie-parser');
const app = express();

app.use(cookieParser());

if (cluster.isMaster) {
    const cpuCount = os.cpus().length;
    for (let i = 0; i < cpuCount; i++) {
        cluster.fork();
    }
} else {
    app.get('/api/link/get_title', async(req, res) => {
        let url = req.query.url;

        log.log(`获取网址标题 ${url}`);

        // let uid = req.cookies['uid'];
        // if (uid.length !== 36) {
        //     return res.json({errcode: 40003, errmsg: '未登陆'});
        // }
        // let [session, sessionError] = await redis.checkAuth(uid);
        // if (session) {

        // } else {
        //     return res.json({errcode: 40003, errmsg: '未登陆'});   
        // }
        
        let [linkUrl, linkParseError] = await link.parse(url);
        if (linkParseError) {
            return res.json(linkParseError);
        }

        let [response, requestError] = await http.get(linkUrl);
        if (requestError) {
            return res.json(requestError);
        }

        let [title, htmlError] = await html.parse(response);
        if (htmlError) {
            return res.json(htmlError);
        }

        return res.json(title);
    })

    app.listen(8002);

    console.log('app is running on port', 8002);
}

cluster.on('exit', (worker) => {
    console.log('mayday! mayday! worker', worker.id, ' is no more!');
    cluster.fork();
})