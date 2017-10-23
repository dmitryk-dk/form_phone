# blacklist ui

### Run
```
docker run -d -p 3000:3000 -v ./db-config.json:/db-config.json --name blacklist dmitrykdk/form_phone go run main.go --db-config=/db-config.json
```

### DB config:
```
{
    "user": "user",
    "password": "passw0rd",
    "host": "tcp(localhost:3306)",
    "dbName": "database",
    "dbDriver": "mysql"
}
```