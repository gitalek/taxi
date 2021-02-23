![build](https://github.com/gitalek/taxi/workflows/build/badge.svg)  
![test](https://github.com/gitalek/taxi/workflows/test/badge.svg)  
![lint](https://github.com/gitalek/taxi/workflows/lint/badge.svg)  

# taxi

### API V1

```sh
$ curl -X POST localhost:9090/calcprice -d '{"coordinates":[[8.681495,49.41461],[8.686507,49.41843]]}'
$ {"price":23080}
```
### API V2

```sh
$ curl -X POST localhost:9090/v2/calcprice -d '{"coordinates":[{"lat":8.681495,"lon":49.41461},{"lat":9.686507,"lon":49.41843}]}'
$ curl -X POST localhost:9090/v2/calcprice -d '{"coordinates":[{"lat":49.41461,"lon":8.681495},{"lat":49.420318,"lon":9.687872}], "strategy":1, "rate":"business"}';
$ {"price":9280.345000000001}
```

# DB migrations (goose)

```sh
$ GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=user password=password dbname=taxi_db sslmode=disable host=localhost port=54320" goose status
```
