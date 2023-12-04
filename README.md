![image](https://raw.githubusercontent.com/logonod/bookmark/master/note-web-frontend/public/images/logo-icon.png)

# bookmark - 我的书签 基于go实现的在线书签管理工具

chrome 书签越来越多，为了快速搜索书签内容，实现了书签管理工具

Screenshot:

![image](https://raw.githubusercontent.com/logonod/bookmark/master/images/bookmark.png)

# keyword

* Go
* Mongodb
* Elasticsearch
* Beanstalkd
* Vue
* Puppeteer

# online demo

在线访客地址 http://bookmark.liuzeyu.me/

# main feature

- 书签打标签
- 按标签展示
- 收藏链接
- 全文爬取
- 全文检索
- 书签搜索

# requirement

- Golang 1.13.2
- Elasticsearch 7.1.0
- Beanstalkd
- Nodejs 10.9.0

# component

| Project | Feature |
|--|--|
| note-ms-server | 爬虫服务 |
| note-server | 后端服务 |
| note-spider | puppeteer爬虫 |
| note-title-server | 标题服务 |
| note-web-frontend | 静态页面 |

# setup

```bash
# install dependency
cd note-ms-server
dep ensure -update
cd ../note-server
dep ensure -update
cd ../note-title-server
npm install
cd ../note-web-frontend
npm install
cd ../note-spider
npm install
```

# compile and run
```bash
# compile and run go server
cd ../note-ms-server
go build main.go
./main serve
cd ../note-server
go build main.go
./main serve
# compile front-end statics
cd ../note-web-frontend
npm run build
# run spider and title server
cd ../note-title-server
npm run service
cd ../note-spider
node index.js
```

# reverse proxy setup

<summary>Nginx</summary>

设置8000和8002端口的反向代理服务以及Vue的前后端分离静态文件配置
```
server {

    listen 80;

    server_name bookmark.liuzeyu.me;

    access_log /var/log/bookmark.liuzeyu.me.log;

    location / {

        root   /var/www/bookmark;
        index  index.html index.htm;

    }

    location /api/user {

        proxy_pass http://127.0.0.1:8000;

    }

    location /api/link {

        proxy_pass http://127.0.0.1:8002;

    }

}
```

</details>

应用当前可以再localhost访问 http://localhost:80

# license 

GPL3
https://github.com/Illumina/licenses/blob/master/gpl-3.0.txt

本项目仅供学习娱乐，请勿滥用。请遵守知乎用户协议合理使用互联网

