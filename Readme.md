# Go + OpenAPI による 超簡易APIサーバ (Sample)

Go + oapi-codegen (echo) のお試し用に作成

Qiita API を実行するのみ
- http://localhost:3030/authenticated_user → https://qiita.com/api/v2/authenticated_user
  - Header
    - Authorization: "Bearer <qiita access token>"
- http://localhost:3030/authenticated_user/items → https://qiita.com/api/v2/authenticated_user/items
  - Header
    - Authorization: "Bearer <qiita access token>"
  - Query parameters
    - page
    - per_page

ref: https://qiita.com/api/v2/docs

## コマンド

### 起動

```
cd src
go run main.go
```

### コード自動生成用
要自動化
```
oapi-codegen -generate "server" -package server ./doc/Qiita-proxy-client.yaml > ./src/autogen/server/api.gen.go

oapi-codegen -generate "types" -package server ./doc/Qiita-proxy-client.yaml > ./src/autogen/server/types.gen.go

oapi-codegen -generate "client" -package client ./doc/Qiita-proxy-client.yaml > ./src/autogen/client/api.gen.go

oapi-codegen -generate "types" -package client ./doc/Qiita-proxy-client.yaml > ./src/autogen/client/types.gen.go

oapi-codegen -generate "spec" -package spec ./doc/Qiita-proxy-client.yaml > ./src/autogen/spec/spec.gen.go
```

## お試し
openapi-generator 試すなら
```
node_modules/.bin/openapi-generator-cli generate -i ./doc/server-openapi.yml -g go-server -o ./test
```
```
docker run -v ${PWD}:/local openapitools/openapi-generator-cli:latest generate -i /local/doc/server-openapi.yaml -g go-server -o /local/out/go
```
