const {
    MONGO_USERNAME,
    MONGO_PASSWORD,
    MONGO_DATABASE,
} = process.env;

admin = db.getSiblingDB('admin');
admin.createUser({
    user: MONGO_USERNAME,
    pwd: MONGO_PASSWORD,
    roles: [{ role: 'read', db: MONGO_DATABASE }],
});

deeb = db.getSiblingDB(MONGO_DATABASE);

deeb.interests.drop();
deeb.interests.insertMany([
	{
		"slug": "example",
        "name": "Example",
        "summary": "An example of an interest",
        "aliases": ["demo", "test", "foo", "bar", "baz", "qux", "quux", "quuux"],
		"passion": 0.0,
		"elaborationUrl": "https://example.com/",
	},
]);

deeb.modes.drop();
deeb.modes.insertMany([
	{
		"id":  0,
		"mode": "polemicist",
		"data": {
			"name": {
				"nick":  "Ed",
				"first": "Edward",
				"last":  "Zitron",
				"slug":  "ed-zitron",
			},
			"portraitUri": "/static/img/me.png",
			"interests": ["privacy", "internet", "ai", "crypto"]
		}
	},
]);

