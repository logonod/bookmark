use note

db.user.insertMany([
	{ 
    "_id": ObjectId("5d9b29abb6fe0f893f79eddc"),
    "phone" : "15501265107", 
    "password" : "$2a$10$IJmrtyewUsZkCJJaTfmitelcve3I2qk86E1upyfdU6UmNwbzo9bYO", 
    "created_at" : new Date(),
    "updated_at" : new Date()
	}
])

db.userid_tag_map.insertMany([
	{ 
    "_id": ObjectId("5da2be9c6166091200474910"),
    "user_id" : ObjectId("5d9b29abb6fe0f893f79eddc"), 
    "tag_name" : "好人", 
    "created_at" : new Date(),
    "updated_at" : new Date()
	},
	{ 
    "_id": ObjectId("5da2be9c6166091200474911"),
    "user_id" : ObjectId("5d9b29abb6fe0f893f79eddc"), 
    "tag_name" : "坏人", 
    "created_at" : new Date(),
    "updated_at" : new Date()
	},
	{ 
    "_id": ObjectId("5da2be9c6166091200474912"),
    "user_id" : ObjectId("5d9b29abb6fe0f893f79eddc"), 
    "tag_name" : "普通人", 
    "created_at" : new Date(),
    "updated_at" : new Date()
	}
])

db.userid_tagids_collect_map.insertMany([
    { 
    "type": "webpage",
    "user_id" : ObjectId("5d9b29abb6fe0f893f79eddc"), 
    "tag_ids" : [ObjectId("5da2be9c6166091200474910")],
    "title": "Multikey Indexes — MongoDB Manual",
    "cover": "https://download-cdn.pmcaff.com/coffee-f7c5894562b3e4101830749d243e50d6?imageView2/1/w/240/h/160/interlace/1/q/75/imageView2/0/format/jpg/interlace/1/q/100",
    "description": "As a result of changes to sorting behavior on array fields in MongoDB 3.6, when sorting on an array indexed with a multikey index the query plan includes a blocking SORT stage.",
    "url": "https://docs.mongodb.com/manual/core/index-multikey/",
    "url_hash": "aaa",
    "site_domain": "docs.mongodb.com",
    "crawl_status": "pending",
    "user_collected": 0,
    "created_at" : new Date(),
    "updated_at" : new Date()
    },
    { 
    "type": "webpage",
    "user_id" : ObjectId("5d9b29abb6fe0f893f79eddc"), 
    "tag_ids" : [ObjectId("5da2be9c6166091200474910")],
    "title": "硅谷早知道",
    "cover": "https://coffee.pmcaff.com//image/show?url=https%3A%2F%2Fdownload-cdn.pmcaff.com%2Fcoffee-e2e3935e1f2b2348d71e6cc6bdce0415%3Fcoffee-w%3D2000px%26coffee-h%3D1500px",
    "description": "太夸张了 一般好的架构师抢着要吧",
    "url": "https://coffee.pmcaff.com/",
    "url_hash": "bbb",
    "site_domain": "coffee.pmcaff.com",
    "crawl_status": "pending",
    "user_collected": 0,
    "created_at" : new Date(),
    "updated_at" : new Date()
    },
    { 
    "type": "webpage",
    "user_id" : ObjectId("5d9b29abb6fe0f893f79eddc"), 
    "tag_ids" : [],
    "title": "todos/todo.go at master · theaaf/todos",
    "cover": "https://img.pmcaff.com/FspC6MIuGOy7gLatXZ0AKT048F_T-picture",
    "description": "Using the Hello World guide, you’ll start a branch, write comments, and open a pull request.",
    "url": "https://github.com/theaaf/todos/blob/master/api/todo.go",
    "url_hash": "ccc",
    "site_domain": "github.com",
    "crawl_status": "pending",
    "user_collected": 0,
    "created_at" : new Date(),
    "updated_at" : new Date()
    }
])