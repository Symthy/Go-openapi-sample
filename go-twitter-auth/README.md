# go twitter auth

## 実行

conf/twitter-conf.yaml に 各Keyを指定

```
cd main
go run main.go
```

## URI

./main/main.go を確認要

## 認証 & Twitter API 実行までの流れ

- consumer key を元に request token を取得
- request token を元に authorization URL 取得＆302リダイレクト
- authorization URL で認可後は callback URL 呼ばれる
- callback 時に request token + oauth_verifier を元に access token 取得
- access token を元に twitter api 実行（今回は アカウント情報取得のみ）

## reference

動機元
- [Go言語でEchoを用いて認証付きWebアプリの作成](https://qiita.com/x-color/items/24ff2491751f55e866cf)

色々見たもの
- [GoでTwitterのOAuth1.0を用いた認証(2016)](https://christina04.hatenablog.com/entry/2016/07/11/193000)
- [GoでTwitter認証をまとめてみた](https://qiita.com/gorilla0513/items/a045c32bc531fdbb0d5e)
- [【第6回】Go言語（Golang）入門～Twitter API利用編～](https://rightcode.co.jp/blog/information-technology/golang-introduction-twitter-api)
- [Go言語でセッション管理](https://qiita.com/int_main_Nick/items/f8dee3c8c07b070b83f1)
- [Twitter APIの使い方まとめ（認証情報の種類・Goサンプル実装）](https://zenn.dev/nekoshita/articles/3c24c302a6a5ee)
- https://github.com/skanehira/vue-go-oauth/blob/a6b137cfca96580b930d1a51793aa7f74ba80890/api/handler/oauth.go
- https://github.com/wheatandcat/dotstamp_server/blob/master/controllers/twitter/oauth.go
- https://github.com/jun06t/oauth-twitter/blob/master/twitter.go
