# KZZ_QUERY

## 主要的环境变量介绍

使用 go env 命令查看go语言环境配置情况
GOROOT: Go的安装路径
GOBIN: Go的二进制文件存放的目录（默认为%GOROOT%/bin）
GOPATH: Go的工作空间，不能与GOROOT相同。
GO111MODULE：开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是 auto。

* off: 关闭模块支持，go会从GOPATH和vendor文件夹寻找包。
* on: 开启模块支持，go会忽略GOPATH和vendor文件夹，只根据go.mod下载依赖。
* auto: 自动，在 GOPATH/src外面且根目录有go.mod文件时，开启模块支持。

在使用模块的时候，GOPATH是无意义的，不过它还是会把下载的依赖储存在GOPATH/pkg/mod 中，也会把go install的结果放在 $GOPATH/bin 中。

在一个新的项目中，需要执行go mod init 来初始化创建文件go.mod，go.mod 中会列出所有依赖包的路径和版本
启用Go模块以后，使用go get xxx时会报错提示"go: cannot find main module; see 'go help modules'"，因为没有找到go.mod文件，所以会报错。你只要在项目根目录下生成一个go.mod文件就可以了。

## vscode配置go语言开发环境需要用到的包

不需要翻墙的可直接go get，否则需要手动下载对应的插件包，放到系统GOPATH/src下

* go tools: 工具类 包含guru、gorename等
 具体参见<https://github.com/golang/tools/>

* gopkgs：Add Imports 添加引用

  具体参见<https://github.com/uudashr/gopkgs/>

* go-outline：一个实用的工具，用于提取go文件中的JSON声明

  具体参见[https://github.com/ramya-rao-a/go-outline](https://github.com/ramya-rao-a/go-outline)

* go-symbols：树状大纲（全局变量、函数、类型、接口）

  一个实用的工具，用于提取go源文件树中的JSON声明
  具体参见[https://github.com/acroca/go-symbols](https://github.com/acroca/go-symbols)

* gotests：为你的源代码生成Go测试

  具体参见[https://github.com/cweill/gotests](https://github.com/cweill/gotests)

* gomodifytags：修改结构字段标签的工具

  具体参见[https://github.com/fatih/gomodifytags](https://github.com/fatih/gomodifytags)

* impl：生成实现接口的方法

  具体参见[https://github.com/josharian/impl](https://github.com/josharian/impl)

* fillstruct：用默认值填充结构文字

  具体参见[https://github.com/davidrjenni/reftools/tree/master/cmd/fillstruct](https://github.com/davidrjenni/reftools/tree/master/cmd/fillstruct)
  安装方法

  ```bash
  go get -u github.com/davidrjenni/reftools/cmd/fillstruct
  ```

* goplay：Go Playground客户端

  具体参见[https://github.com/haya14busa/goplay](https://github.com/haya14busa/goplay)

* godoctor：Golang重构引擎

  具体参见[https://github.com/godoctor/godoctor](https://github.com/godoctor/godoctor)

* dlv：Go编程语言的调试器

  具体参见[https://github.com/go-delve/delve](https://github.com/go-delve/delve)

* gocode：代码自动补全

  具体参见[https://github.com/mdempsky/gocode](https://github.com/mdempsky/gocode)

* godef：跳转到定义（ctrl+鼠标左击，查看定义的代码）

  具体参见[https://github.com/rogpeppe/godef](https://github.com/rogpeppe/godef)

* goreturns；自动补全return语句中的零值

  ``` go
  func F() (string, int, error) {
    return errors.New("foo")
    // 自动补全为 => return "", 0, errors.New("foo")
  }
  ```

  具体参见[https://github.com/sqs/goreturns](https://github.com/sqs/goreturns)

* golint：代码检测

  对在命令行中命名的Go源文件进行代码检测。
  具体参见[https://github.com/golang/lint](https://github.com/golang/lint)

* gocode-gomod：Go语言根据上下文自动补全的一个守护进程
  具体参见[https://github.com/stamblerre/gocode](https://github.com/stamblerre/gocode)

## go包文档搜索

可以查看主流包的文档介绍和使用示例

<https://godoc.org/>

示例：jwt包 <https://godoc.org/github.com/dgrijalva/jwt-go/>

## go标准库文档

<https://studygolang.com/pkgdoc>