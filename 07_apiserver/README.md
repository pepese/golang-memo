CLI 作成ツール [cobra](https://github.com/spf13/cobra) を利用して簡易な API サーバを作成してみる。  
仕様は以下。

- `apiserver start` コマンドでサーバが起動する
- apiserver は [郵便番号検索API](http://zipcloud.ibsnet.co.jp/doc/api) へ GET リクエストして住所を返却する
    - openapi/swagger を使う（かな？）
    - クエリパラメータは `zipcode` のみ
    - 例 `http://zipcloud.ibsnet.co.jp/api/search?zipcode=xxxxxxx`
    - apiserver の URL は `http://localhost/address?zipcode=xxxxxxx`
        - デフォルトの `zipcode` は `1638001` 東京都庁
- シグナル（ kill ）を受け取って Graceful Stop する

# 最初にちょっとメモ

## golang における OpenAPI/Swagger 周りの状況について。[2019/01/04]

Swaggerは2015年12月31日に、マイクロソフトやグーグルなどにより設立されたOAI(Open API Initiative)に寄贈され、Swagger SpecはOAS（OAS:Open API Specification）へと名称変更された。  
Swagger Specの1行目に記載する識別子が、`swagger: "2.0"` から `openapi: 3.0.0` になり、バージョンも2.xから3.xになった。  
Open API は バージョン 3.0 が初版であり、Swagger 3.0 と OAS:Open API Specification 3.0 は同じものと考えてよい。ただし、 Swagger Spec の 1 行目には `swagger: 3.0.0` とは書かず `openapi: 3.0.0` と書く。（ [参考](https://news.mynavi.jp/itsearch/article/devsoft/3854) ）

現状、 golang の Swagger コンバーター は以下のものがある

- [go-swagger](https://github.com/go-swagger/go-swagger)
    - 僅差で最もメジャーっぽいが、 現状 Swagger 2.0 で OAP 3.0 に対応中の模様
- [goa](https://github.com/goadesign/goa)
    - go-swagger より僅差でマイナー。独自の DSL で記述し、 Swagger Spec を出力できる。
- [go-openapi](https://github.com/go-openapi/spec)
    - k8s が採用している [go-restful](https://github.com/emicklei/go-restful) に対応

上記のような状況なので、ここでは、 最新ではない Swagger 2.0 ではあるが、 `go-swagger` を使おうかな、、、どないしょう。  
ちな k8s は現状 Swagger 2.0 。

# [cobra](https://github.com/spf13/cobra)

Cobra のセットアップと初期プロジェクトの作成。

```bash
$ go get -u github.com/spf13/cobra/cobra
$ which cobra
${GOPATH}/bin/cobra
```

Cobra Generator を使って、プロジェクトの初期化。  
`${GOPATH}/src` 配下からの絶対パスで実行する必要あり。  
また、空ディレクトリでないとエラーになるので注意。

```bash
$ cobra init github.com/pepese/golang-sample/04_apiserver
Your Cobra application is ready at
${GOPATH}/src/github.com/pepese/golang-sample/04_apiserver

Give it a try by going there and running `go run main.go`.
Add commands to it by running `cobra add [cmdname]`.

$ tree .
.
├── LICENSE
├── README.md
├── cmd
│   └── root.go
└── main.go

1 directory, 4 files
```

`start` コマンドを追加する。

```bash
$ pwd
${GOPATH}/src/github.com/pepese/golang-sample/04_apiserver

$ cobra add start
start created at ${GOPATH}/src/github.com/pepese/golang-sample/04_apiserver/cmd/start.go

$ tree .
.
├── LICENSE
├── README.md
├── cmd
│   ├── root.go
│   └── start.go // 追加された
└── main.go

1 directory, 5 files
```

ここで試しにビルドして実行。

```bash
$ go build main.go
$ ./main
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  04_apiserver [command]

Available Commands:
  help        Help about any command
  start       A brief description of your command

Flags:
      --config string   config file (default is $HOME/.04_apiserver.yaml)
  -h, --help            help for 04_apiserver
  -t, --toggle          Help message for toggle

Use "04_apiserver [command] --help" for more information about a command.

$ ./main start
start called
```

# [go-swagger](https://github.com/go-swagger/go-swagger)

go-swagger のセットアップ。

```bash
$ go get -u github.com/go-swagger/go-swagger/cmd/swagger
$ which swagger
${GOPATH}/bin/swagger
```

API 定義の雛形を作成。

```bash
$ swagger init spec \
  --title "A API Server Sample" \
  --version 0.1.0 \
  --scheme http \
  --consumes application/json \
  --produces application/jaon

$ cat swagger.yml
consumes:
- application/json
info:
  title: A API Server Sample
  version: 0.1.0
paths: {}
produces:
- application/jaon
schemes:
- http
swagger: "2.0"
```

雛形生成微妙じゃね、、、 path の定義も含めて書き直す。  
書き直したらバリデーションチェックしてサーバのコード生成。

```bash
$ swagger validate swagger.yml
yyyy/MM/dd hh:ss:mm
The swagger spec at "swagger.yml" is valid against swagger specification 2.0

$ swagger generate server -f swagger.yml -A apiserver
2019/01/04 11:09:10 validating spec ${GOPATH}/src/github.com/pepese/golang-sample/04_apiserver/swagger.yml
2019/01/04 11:09:11 preprocessing spec with option:  minimal flattening
2019/01/04 11:09:11 building a plan for generation
2019/01/04 11:09:11 planning definitions
2019/01/04 11:09:11 planning operations
2019/01/04 11:09:11 grouping operations into packages
2019/01/04 11:09:11 planning meta data and facades
2019/01/04 11:09:11 rendering 0 models
2019/01/04 11:09:11 rendering 1 operation groups (tags)
2019/01/04 11:09:11 rendering 1 operations for operations
2019/01/04 11:09:11 rendering 4 templates for operation Apiserver
2019/01/04 11:09:11 name field GetAddress
2019/01/04 11:09:11 package field operations
2019/01/04 11:09:11 creating generated file "get_address_parameters.go" in "restapi/operations" as parameters
2019/01/04 11:09:11 executed template asset:serverParameter
2019/01/04 11:09:11 name field GetAddress
2019/01/04 11:09:11 package field operations
2019/01/04 11:09:11 creating generated file "get_address_urlbuilder.go" in "restapi/operations" as urlbuilder
2019/01/04 11:09:11 executed template asset:serverUrlbuilder
2019/01/04 11:09:11 name field GetAddress
2019/01/04 11:09:11 package field operations
2019/01/04 11:09:11 creating generated file "get_address_responses.go" in "restapi/operations" as responses
2019/01/04 11:09:11 executed template asset:serverResponses
2019/01/04 11:09:11 name field GetAddress
2019/01/04 11:09:11 package field operations
2019/01/04 11:09:11 creating generated file "get_address.go" in "restapi/operations" as handler
2019/01/04 11:09:11 executed template asset:serverOperation
2019/01/04 11:09:11 rendering 0 templates for operation group Apiserver
2019/01/04 11:09:11 rendering support
2019/01/04 11:09:11 rendering 6 templates for application Apiserver
2019/01/04 11:09:11 name field Apiserver
2019/01/04 11:09:11 package field operations
2019/01/04 11:09:11 creating generated file "configure_apiserver.go" in "restapi" as configure
2019/01/04 11:09:11 executed template asset:serverConfigureapi
2019/01/04 11:09:11 name field Apiserver
2019/01/04 11:09:11 package field operations
2019/01/04 11:09:11 creating generated file "main.go" in "cmd/apiserver-server" as main
2019/01/04 11:09:11 executed template asset:serverMain
2019/01/04 11:09:11 name field Apiserver
2019/01/04 11:09:11 package field operations
2019/01/04 11:09:11 creating generated file "embedded_spec.go" in "restapi" as embedded_spec
2019/01/04 11:09:11 executed template asset:swaggerJsonEmbed
2019/01/04 11:09:11 name field Apiserver
2019/01/04 11:09:11 package field operations
2019/01/04 11:09:11 creating generated file "server.go" in "restapi" as server
2019/01/04 11:09:11 executed template asset:serverServer
2019/01/04 11:09:11 name field Apiserver
2019/01/04 11:09:11 package field operations
2019/01/04 11:09:11 creating generated file "apiserver_api.go" in "restapi/operations" as builder
2019/01/04 11:09:11 executed template asset:serverBuilder
2019/01/04 11:09:11 name field Apiserver
2019/01/04 11:09:11 package field operations
2019/01/04 11:09:11 creating generated file "doc.go" in "restapi" as doc
2019/01/04 11:09:11 executed template asset:serverDoc
2019/01/04 11:09:11 Generation completed!

For this generation to compile you need to have some packages in your GOPATH:

        * github.com/go-openapi/runtime
        * github.com/jessevdk/go-flags

You can get these now with: go get -u -f ./...
```

`restapi` `restapi/operations` `cmd/apiserver-server` が作られた。  
とりあえずサーバ側を動かしてみる。

```bash
$ dep init
$ go run cmd/apiserver-server/main.go --port=8080
$ curl http://localhost:8080/address
curl: (52) Empty reply from server
```

動かぬ。
- https://qiita.com/croquette0212/items/f7a21608b1eec1446c1c

go-restful で swagger 生成できやんのか？  
なんか、シグナルさばき出来てる。。。？ ごいすー  
`swagger generate server` で出力されたコード見てみる。