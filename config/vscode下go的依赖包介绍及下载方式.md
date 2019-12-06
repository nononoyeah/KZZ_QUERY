# Markdown heading

**注意：**

  **1.  以下所有关于gitclone和goinstall的操作都要在$GOPATH/src目录下完成。如果采用推荐的go get 方式获取则没有限制。**

  **2.  系统环境为MacOS**

<table>
  <tr>
    <th>名称</th>
    <th>介绍</th>
    <th>描述</th>
    <th>获取方式</th>
  </tr>
  <tr>
    <td rowspan="2">go tools</td>
    <td rowspan="2">工具类 <br/>包含guru、gorename等</td>
    <td rowspan="2">https://github.com/golang/tools</td>
    <td>go get -u golang.org/x/tools/...</td>
  </tr>
  <tr>
    <td>git clone https://github.com/golang/tools.git golang.org/x/tools</td>
  </tr>
  <tr>
    <td>gopkgs</td>
    <td>添加导入的包</td>
    <td>https://github.com/uudashr/gopkgs</td>
    <td>go get -u github.com/uudashr/gopkgs/cmd/gopkgs</td>
  </tr>
  <tr>
    <td>go-outline</td>
    <td>一个实用的工具，用于提取go文件中的JSON声明</td>
    <td>https://github.com/ramya-rao-a/go-outline</td>
    <td>go get -u github.com/ramya-rao-a/go-outline</td>
  </tr>
  <tr>
    <td>go-symbols</td>
    <td>用于提取go源文件树中的JSON声明</td>
    <td>https://github.com/acroca/go-symbols</td>
    <td>go get -u github.com/newhook/go-symbols</td>
  </tr>
  <tr>
    <td>gotests</td>
    <td>为你的源代码生成Go测试</td>
    <td>https://github.com/cweill/gotests</td>
    <td>go get -u github.com/cweill/gotests/...</td>
  </tr>
  <tr>
    <td>gomodifytags</td>
    <td>修改结构字段标签的工具<br>比如添加json标签</td>
    <td>https://github.com/fatih/gomodifytags</td>
    <td>go get github.com/fatih/gomodifytags</td>
  </tr>
  <tr>
    <td>impl</td>
    <td>生成实现接口的方法</td>
    <td>https://github.com/josharian/impl</td>
    <td>go get -u github.com/josharian/impl</td>
  </tr>
  <tr>
    <td>fillstruct</td>
    <td>用默认值填充结构文字</td>
    <td>https://github.com/davidrjenni/reftools/tree/master/cmd/fillstruct</td>
    <td>go get -u github.com/davidrjenni/reftools/cmd/fillstruct</td>
  </tr>
  <tr>
    <td>goplay</td>
    <td>Go Playground客户端</td>
    <td>https://github.com/haya14busa/goplay</td>
    <td>go get -u github.com/haya14busa/goplay</td>
  </tr>
  <tr>
    <td>godoctor</td>
    <td>Golang重构引擎</td>
    <td>https://github.com/godoctor/godoctor</td>
    <td>git clone "https://github.com/godoctor/godoctor" github.com/godoctor/godoctor</td>
  </tr>
  <tr>
    <td>dlv</td>
    <td>为你的源代码生成Go测试</td>
    <td>https://github.com/go-delve/delve</td>
    <td>go get -u github.com/go-delve/delve/cmd/dlv <br>//MacOS</td>
  </tr>
  <tr>
    <td>gocode</td>
    <td>代码自动补全</td>
    <td>https://github.com/mdempsky/gocode</td>
    <td>go get -u github.com/mdempsky/gocode</td>
  </tr>
  <tr>
    <td>godef</td>
    <td>跳转到定义</td>
    <td>https://github.com/rogpeppe/godef</td>
    <td>git clone "https://github.com/rogpeppe/godef" github.com/rogpeppe/godef</td>
  </tr>
  <tr>
    <td>goreturns</td>
    <td>自动补全return语句中的零值</td>
    <td>https://github.com/sqs/goreturf</td>
    <td>go get -u github.com/sqs/goreturns</td>
  </tr>
  <tr>
    <td>golint</td>
    <td>对在命令行中命名的Go源文件进行代码检测</td>
    <td>https://github.com/golang/lint</td>
    <td>go get -u golang.org/x/lint/golint</td>
  </tr>
  <tr>
    <td>gocode-gomod</td>
    <td>Go语言根据上下文自动补全的一个守护进程</td>
    <td>https://github.com/stamblerre/gocode</td>
    <td>go get -u -v github.com/stamblerre/gocode</td>
  </tr>
</table>

>上面介绍了vscode中常见工具和依赖包的github地址以及基本的包获取方式
通常优先采取推荐方法：go get 获取包，这样不用执行手动克隆代码（git clone）和手动安装（go install）.
>下面贴一个，自己写的批量安装工具和依赖的包的bash脚本，可以参考一下。自测通过了的。

```bash
#! /bin/bash

echo 查看当前go配置环境
go env

cd /Users/apple/go/src #进入你的GOPATH目录下的src文件夹下

#read -p "输入你的GOPATH路径: " gopath
#cd $gopath/src/

echo "当前位置:" 
pwd

# 插件依赖工具
git clone https://github.com/golang/tools.git golang.org/x/tools
# 包含以下工具
# go get -u -v "golang.org/x/tools/cmd/guru"
# go get -u -v "golang.org/x/tools/cmd/gorename"


# 获取包的源码并编译 -v显示执行的命令（不用单独执行git clone + go install）
go get -u -v "github.com/go-delve/delve/cmd/dlv"
go get -u -v "github.com/uudashr/gopkgs/cmd/gopkgs"
go get -u -v "github.com/haya14busa/goplay"
go get -u -v "github.com/davidrjenni/reftools"
go get -u -v "github.com/stamblerre/gocode"


#下载依赖包到$GOPATH/src文件夹下的指定目录
# 克隆代码到指定路径下
git clone "https://github.com/golang/lint.git" golang.org/x/lint

git clone "https://github.com/mdempsky/gocode" github.com/mdempsky/gocode
git clone "https://github.com/ramya-rao-a/go-outline" github.com/ramya-rao-a/go-outline #go get 可以成功
git clone "https://github.com/acroca/go-symbols" github.com/acroca/go-symbols #go get 可以成功
git clone "https://github.com/cweill/gotests" github.com/cweill/gotests #go get 可以成功
git clone "https://github.com/fatih/gomodifytags" github.com/fatih/gomodifytags #go get 可以成功
git clone "https://github.com/josharian/impl" github.com/josharian/impl
git clone "https://github.com/godoctor/godoctor" github.com/godoctor/godoctor
git clone "https://github.com/rogpeppe/godef" github.com/rogpeppe/godef
git clone "https://github.com/sqs/goreturns" github.com/sqs/goreturns

# #执行go install
go install golang.org/x/lint

go install github.com/mdempsky/gocode
go install github.com/ramya-rao-a/go-outline
go install github.com/acroca/go-symbols
go install github.com/cweill/gotests/...
go install github.com/fatih/gomodifytags
go install github.com/josharian/impl
go install github.com/godoctor/godoctor
go install github.com/go-delve/delve/cmd/dlv
go install github.com/rogpeppe/godef
go install github.com/sqs/goreturns

#根据需要安装所需工具(可以到golang.org/x/tools文件夹里去查看有哪些工具)
go install golang.org/x/tools/cmd/guru
go install golang.org/x/tools/cmd/gorename

#注：优先使用go-get方法去获取工具包和依赖包，
# 如果无法获取，尝试到github使用git clone克隆对应的源代码到制定目录，在使用go install进行安装

```