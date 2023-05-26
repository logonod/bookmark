use note;

db.webpage.createIndex(
	{ url_hash: 1 }, { unique: true }
)
