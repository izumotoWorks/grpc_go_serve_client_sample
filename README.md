# gRPC on Go & Windows 動作 サンプル

殴り書きなのでそのうち整理。

## Use version

- go version go1.21.7 windows/amd64

## 謝辞

このサンプルは、以下の記事を参考に作成しています。

> 参考
> (作ってわかる！ はじめての gRPC)[https://zenn.dev/hsaki/books/golang-grpc-starting]

## protoc メモ

```
protoc
--go_out=<出力ディレクトリ>
--go-grpc_out=<出力ディレクトリ>
--proto_path=<プロトファイルのディレクトリ>
<プロトファイルのパス>

```

Windows のシェルスクリプトで proto からファイルを自動生成したかったので以下のように設定。

シェルスクリプトファイルは同梱してませんので
必要であれば各自コピペで作ってください。

```ps1
protoc --go_out=src\grpc\gen --go_opt=paths=source_relative `
    --go-grpc_out=src\grpc\gen --go-grpc_opt=paths=source_relative `
    --proto_path=grpc\proto `
    grpc\proto\helloworld.proto
```

windows の ps を使う場合は上記の表記で
/src/grpc/gen の中に生成できる

生成したい場所は任意で変えてみてください。

## grpcurl メモ

grpcurl をウィンドウズに入れるにはバイナリを落として gopath のところに直接置くのでおｋ

実際に grpc を叩くときにダブルクォートはエスケープしないといけない

```
# ng command
grpcurl -plaintext -d '{"name":"hogehoge"}' localhost:8080 grpc.Greeter.SayHello
Error invoking method "grpc.Greeter.SayHello": error getting request data: invalid character 'n' looking for beginning of object key string
```

```
# ok command
# req
grpcurl -plaintext -d '{\"name\":\"hogehoge\"}' localhost:8080 grpc.Greeter.SayHello
```

```
# res
{
"message": "Hello, hogehoge!"
}
```
