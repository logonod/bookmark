use note

db.webpage.insertMany([
    { 
    "title": "Multikey Indexes — MongoDB Manual",
    "cover": "https://download-cdn.pmcaff.com/coffee-f7c5894562b3e4101830749d243e50d6?imageView2/1/w/240/h/160/interlace/1/q/75/imageView2/0/format/jpg/interlace/1/q/100",
    "description": "As a result of changes to sorting behavior on array fields in MongoDB 3.6, when sorting on an array indexed with a multikey index the query plan includes a blocking SORT stage.",
    "meta_description": "",
    "full_text": "To index a field that holds an array value, MongoDB creates an index key for each element in the array. These multikey indexes support efficient queries against array fields. Multikey indexes can be constructed over arrays that hold both scalar values [1] (e.g. strings, numbers) and nested documents.",
    "url": "https://docs.mongodb.com/manual/core/index-multikey/",
    "url_hash": "d7a65f555b50bb07259fac1ba01d78299863078e000052",
    "site_domain": "docs.mongodb.com",
    "user_collected": 0,
    "created_at" : new Date(),
    "updated_at" : new Date()
    },
    { 
    "title": "硅谷早知道",
    "cover": "https://coffee.pmcaff.com//image/show?url=https%3A%2F%2Fdownload-cdn.pmcaff.com%2Fcoffee-e2e3935e1f2b2348d71e6cc6bdce0415%3Fcoffee-w%3D2000px%26coffee-h%3D1500px",
    "description": "太夸张了 一般好的架构师抢着要吧",
    "meta_description": "互联网即时资讯平台，随时获取新鲜事。",
    "full_text": "【“天枢”人工智能开源开放平台在杭州发布】“天枢”人工智能开源开放平台2日在浙江杭州正式发布。该平台具备高性能核心计算框架，提供一站式全功能AI开发套件，将提升人工智能技术的研发效率、扩大算法模型的应用范围，进一步构建人工智能生态“朋友圈”。",
    "url": "https://coffee.pmcaff.com/",
    "url_hash": "689d08df4fa33cd6d5ab7a5511dbb7c7644fdd03000026",
    "site_domain": "coffee.pmcaff.com",
    "user_collected": 0,
    "created_at" : new Date(),
    "updated_at" : new Date()
    },
    { 
    "title": "todos/todo.go at master · theaaf/todos",
    "cover": "https://img.pmcaff.com/FspC6MIuGOy7gLatXZ0AKT048F_T-picture",
    "description": "Using the Hello World guide, you’ll start a branch, write comments, and open a pull request.",
    "meta_description": "互联网即时资讯平台，随时获取新鲜事。",
    "full_text": "func (a *API) GetTodos(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {",
    "url": "https://github.com/theaaf/todos/blob/master/api/todo.go",
    "url_hash": "06b8d274af21f39d939c6cc7467432022ad683ff000055",
    "site_domain": "github.com",
    "user_collected": 0,
    "created_at" : new Date(),
    "updated_at" : new Date()
    }
])