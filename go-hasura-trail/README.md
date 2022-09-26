# golang Hasura trail

## コード

ref: https://hasura.io/blog/building-a-graphql-api-with-golang-postgres-and-hasura/

```
go install github.com/99designs/gqlgen@latest
```

tools.go を作成

```
go run github.com/99designs/gqlgen init
```

```
go run github.com/99designs/gqlgen generate
```

## docker での hasura 利用

ref: https://hasura.io/docs/latest/getting-started/docker-simple/

docker-compose.yml を取得

https://github.com/hasura/graphql-engine/blob/stable/install-manifests/docker-compose/docker-compose.yaml

起動

```
docker-compose up -d
```

DB 接続

http://localhost:8080/console

※ database url は docker-compose.yml に記載あり

## Try

実装して起動

※ pg: SSL is not enabled on the server が発生。 DB_URL の末尾に ?sslmode=disable を指定して解決

ref: https://bun.uptrace.dev/postgres/#pgdriver

```
go run ./server.go
```

コンソールから以下を実行

http://localhost:8080/console

```
mutation createMovie {
 createMovie(
   input: {
     title: "Rise of GraphQL Warrior Pt2"
     url: "https://riseofgraphqlwarriorpt2.com/"
   }
 ){
   id
 }
}
```

```
query getMovies {
 movies {
     id
     title
     url
     releaseDate
 }
}
```
