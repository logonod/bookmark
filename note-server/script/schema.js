use note;

db.userid_tag_map.createIndex(
	{ user_id: 1, tag_name: 1 }, { unique: true }
)

db.userid_tag_map.createIndex(
	{ user_id: 1, _id: -1, tag_name: 1 }
)

db.userid_tag_map.createIndex(
	{ _id: 1, tag_name: 1 }
)

db.userid_tag_map.createIndex(
	{ user_id: 1, tag_name: 1, _id: 1 }
)

db.userid_tagids_collect_map.createIndex(
	{ user_id: 1, tag_ids: 1, _id: -1 }
)

db.userid_tagids_collect_map.createIndex(
	{ user_id: 1, url_hash: 1 }, { unique: true }
)

user {
	_id,
	phone,
	password,
	created_at,
	updated_at
}

userid_tagname_map {
	_id,
	user_id,
	tag_name,
	created_at,
	updated_at
}

userid_tagname_post_map {
	_id,
	user_id,
	tag_name,
	type,
	title,
	cover,
	description,
	url,
	url_hash,
	site,
	site_favicon,
	created_at,
	updated_at
}