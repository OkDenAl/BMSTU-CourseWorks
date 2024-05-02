rootUser = {
    user: "root",
    pwd: "password",
    roles: [
        { role: "clusterAdmin", db: "admin" },
        { role: "readAnyDatabase", db: "admin" },
        "readWrite"
    ]
};
db.createUser(rootUser);

user = {
    user: "user",
    pwd: "password",
    roles: [{ role: 'readWrite', db: 'story_stat' }]
};
db.createUser(user);
db.getSiblingDB('story_stat').createUser(user);
