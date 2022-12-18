# Golang 色々試しモノ置き場

各々の詳細は、各フォルダ下の README.md 参照

## go-config_load

go-ini による `<key>=<value>` 形式の conf ファイル読み込みお試し

go-ini とは疎結合で実装

## go-cross-compile

クロスコンパイル、Windows 環境でのコマンド実行挙動お試し

## go-custom-error

独自のエラー実装お試し

## go-openapi-samples

go-oapigen (Echo) の client/server 自動生成お試し

Qitta API をプロキシーするのみ

## go-opentelementry-trial

OpenTelemetry の チュートリアルお試し

## go-pararell

並列処理のお試し。goroutine、channel、waitgroup 等基本的な物

## go-twitter-auth

Echo + twitter oauth api (v1.1) での認証＆twitter api 実行の実装お試し

request token 及び access token は session で保持

認証後は、twitter api (account 取得) を実行可能

## go-win-service

OSS：kardianos/service を使用しての Windows サービス作成お試し

## go-zap-logger

OSS：Zap を使用してのログ出力＆フォーマット成形お試し
