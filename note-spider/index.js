const Jackd = require('jackd');
const beanstalkd = new Jackd();
const getCollect = require('./api/get_collect').get;
const postCollect = require('./api/create_collect').post;
const updateCollectStatus = require('./api/update_status_collect').post;
const spider = require('./spider/textit').get;

start()

async function work_once(beanstalkd) {
    var preid;
    try {
        const { id, payload } = await beanstalkd.reserve();
        preid = id;
        var data = JSON.parse(payload);
        var r = await getCollect(data.user_id, data.url_hash);
        if (r[1] != null) {
            console.log(r[1]);
            await beanstalkd.release(id, {delay: 10});
        } else {
            var data = r[0].collect;
            var r = await spider(data.url);
            if (r[1] != null) {
                let paylow = {user_id: data.user_id, url_hash: data.url_hash, url: data.url, crawl_status: "failed"}
                await beanstalkd.release(id, {delay: 10});    
            } else {
                let meta = r[0].dsec || "";
                let desc = "";
                let fulltext = r[0].fulltext || "";
                if (r[0].fulltext) {
                    desc = r[0].fulltext;
                } else {
                    desc = r[0].desc || "";
                }
                desc = desc.substring(0, 100);
                let user_id = data.user_id;
                let url_hash = data.url_hash;
                let title = data.title;
                let cover= data.cover;
                let description = desc;
                let meta_description = meta;
                let full_text = fulltext;
                let url = data.url;
                let site_domain = data.site_domain;
                let payload = { user_id, url_hash, title, cover, description, meta_description, full_text, url, site_domain }
                var r = await postCollect(payload);            
                if (r[1] != null) {
                    if (r[0].crawl_status == "failed") {
                        await beanstalkd.delete(id);
                    } else {
                        await beanstalkd.release(id, {delay: 10});
                    }
                } else {
                    console.log(r[0]);
                    await beanstalkd.delete(id);
                }
            }
        }
    } catch (err) {
        // Log error somehow
        console.error(err)
        if (preid) {
            await beanstalkd.release(preid, {delay: 10});
        }   
    }
}

async function start() {
  // Might want to do some error handling around connections
  await beanstalkd.connect({host: "192.168.1.137", "port": 11300});

  while (true) {
    await work_once(beanstalkd);
  }
}

// const getCollect = require('./api/get_collect').get;
// getCollect()