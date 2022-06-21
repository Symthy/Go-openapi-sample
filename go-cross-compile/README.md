# Cross Compile Sample & exec.Cmd with Windows

build

```sh
./build.sh
```

change vscode gopls buildFlags to windows or linux

```sh
./change_build_flag.sh
```

## exec.Command with windows

- パスの `\` をエスケープしなかった場合、`\` が消える

```
error: exec: "D:dev_gitgolang-practicesgo-cross-compilesubmain.exe": file does not exist
```

- オプション指定(仮)色々試した結果

```golang
name := flag.String("exec", cmdPath, "exec command")
flag.Parse()
opts := []string{
	"-a", "x\"xx", "-b", "y\\yy", "-c", "zz/z", "-d", "t\"ru\"e", "-e", "20'22",
}
```

エスケープされていない `\` や `"` を go による外部コマンド実行時のコマンド引数に渡してみる

エスケープされていない `"` が exec.Command に渡ると変な変換がされる… （これが cmd.exe による独自パース処理の産物か？）

※ Main 側では、`exec.Command(*name, opts...)` 実行。エスケープされていない `\` や `"` を Sub 側の `exec.Command` の 第二引数に渡すためのもの

```
Sub Args: D:\dev_git\golang-practices\go-cross-compile\sub\main.exe -a x"xx -b "y\y y" -c z:z/z -d t"ru"e -e 2^0'2-2
[Exec Command (use Args)]
Sub Stdout:
D:\dev_git\golang-practices\go-cross-compile\sub\main.exe -a x\xx -b \y\y y" -c z:z/z -d t\ru"e -e 2^0'2-2
```

`syscall.SysProcAttr` の `CmdLine` に `strings.Join(opts, " ")` で１文字列にして `exec.Command(*name)`の戻り値（`exec.Cmd`）の `SysProcAttr`に代入して実行にすると、`syscall.SysProcAttr` に渡した分は、`"` がなくなり、変な変換を抑止できる

```
Sub Args: D:\dev_git\golang-practices\go-cross-compile\sub\main.exe -a x"xx -b "y\y y" -c z:z/z -d t"ru"e -e 2^0'2-2
[Exec Command (use syscall.SysProcAttr)]
Sub Stdout:
-a xxx -b yy y -c z:z/z -d true -e 2^0'2-2
```

## ビルド制約

ビルド制約で、特定プラットフォーム用の処理を実装可能

ただし、gopls が複数プラットフォーム用のソースを同時には解析はできない

gopls

ビルドタグは、2 つ同時には設定できない。1 つしか設定できない＝ linux,windows 等複数プラットフォームに対する個別処理を実装する際には、随時 buildFlags を切り替えるしかなさそう…

[x/tools/gopls: improve handling for build tags #29202](https://github.com/golang/go/issues/29202#issuecomment-1013455808)

## ref

ビルド制約

- [GitHub Actions で Go のソースコードをクロスコンパイルするときに、ビルドが失敗する理由とその対策](https://developer.hatenastaff.com/entry/2021/04/23/093000)

gopls

- [gopls returns the error "gopls: no packages returned: packages.Load error" for github.com/Shopify/sarama](https://golangshowcase.com/question/gopls-returns-the-error-gopls-no-packages-returned-packages-load-error-for-github-com-shopify-sarama)

exec.Command

- [Go の exec.Command を調査する](https://qiita.com/TsuyoshiUshio@github/items/22cafc8a4dc097add73b)

shell 作成時の参考

- [シェルスクリプト/文法/文字列操作/部分一致で検索する](https://yanor.net/wiki/?%E3%82%B7%E3%82%A7%E3%83%AB%E3%82%B9%E3%82%AF%E3%83%AA%E3%83%97%E3%83%88/%E6%96%87%E6%B3%95/%E6%96%87%E5%AD%97%E5%88%97%E6%93%8D%E4%BD%9C/%E9%83%A8%E5%88%86%E4%B8%80%E8%87%B4%E3%81%A7%E6%A4%9C%E7%B4%A2%E3%81%99%E3%82%8B)
