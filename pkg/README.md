# pkg

go公共pkg。

## usage

其他项目使用，需要指定特定分支以减少冗余文件，加快下载速度。
```shell
$ go get github.com/gomicroim/gomicroim/pkg@gh-pages
```

## publish

公共模块修改后，需要推送修改到分支：
```shell
$ sh publish.sh
```