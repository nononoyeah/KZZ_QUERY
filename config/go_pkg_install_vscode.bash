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