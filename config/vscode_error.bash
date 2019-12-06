go.toolsGopath setting is not set. Using GOPATH /Users/apple/go
Installing 17 tools at /Users/apple/go/bin in module mode.
  gocode
  gopkgs
  go-outline
  go-symbols
  guru
  gorename
  gotests
  gomodifytags
  impl
  fillstruct
  goplay
  godoctor
  dlv
  gocode-gomod
  godef
  goreturns
  golint

Installing github.com/mdempsky/gocode FAILED
Installing github.com/uudashr/gopkgs/cmd/gopkgs FAILED
Installing github.com/ramya-rao-a/go-outline SUCCEEDED
Installing github.com/acroca/go-symbols SUCCEEDED
Installing golang.org/x/tools/cmd/guru SUCCEEDED
Installing golang.org/x/tools/cmd/gorename SUCCEEDED
Installing github.com/cweill/gotests/... SUCCEEDED
Installing github.com/fatih/gomodifytags SUCCEEDED
Installing github.com/josharian/impl FAILED
Installing github.com/davidrjenni/reftools/cmd/fillstruct FAILED
Installing github.com/haya14busa/goplay/cmd/goplay FAILED
Installing github.com/godoctor/godoctor FAILED
Installing github.com/go-delve/delve/cmd/dlv SUCCEEDED
Installing github.com/stamblerre/gocode FAILED
Installing github.com/rogpeppe/godef FAILED
Installing github.com/sqs/goreturns FAILED
Installing golang.org/x/lint/golint FAILED

10 tools failed to install.

gocode:
Error: Command failed: /usr/local/bin/go get -v github.com/mdempsky/gocode
go get github.com/mdempsky/gocode: module github.com/mdempsky/gocode: Get https://proxy.golang.org/github.com/mdempsky/gocode/@v/list: dial tcp 172.217.27.145:443: i/o timeout
go get github.com/mdempsky/gocode: module github.com/mdempsky/gocode: Get https://proxy.golang.org/github.com/mdempsky/gocode/@v/list: dial tcp 172.217.27.145:443: i/o timeout

gopkgs:
Error: Command failed: /usr/local/bin/go get -v github.com/uudashr/gopkgs/cmd/gopkgs
go get github.com/uudashr/gopkgs/cmd/gopkgs: module github.com/uudashr/gopkgs/cmd/gopkgs: Get https://proxy.golang.org/github.com/uudashr/gopkgs/cmd/gopkgs/@v/list: dial tcp 172.217.27.145:443: i/o timeout
go get github.com/uudashr/gopkgs/cmd/gopkgs: module github.com/uudashr/gopkgs/cmd/gopkgs: Get https://proxy.golang.org/github.com/uudashr/gopkgs/cmd/gopkgs/@v/list: dial tcp 172.217.27.145:443: i/o timeout

impl:
Error: Command failed: /usr/local/bin/go get -v github.com/josharian/impl
go get github.com/josharian/impl: module github.com/josharian/impl: Get https://proxy.golang.org/github.com/josharian/impl/@v/list: dial tcp 172.217.27.145:443: i/o timeout
go get github.com/josharian/impl: module github.com/josharian/impl: Get https://proxy.golang.org/github.com/josharian/impl/@v/list: dial tcp 172.217.27.145:443: i/o timeout

fillstruct:
Error: Command failed: /usr/local/bin/go get -v github.com/davidrjenni/reftools/cmd/fillstruct
go get github.com/davidrjenni/reftools/cmd/fillstruct: module github.com/davidrjenni/reftools/cmd/fillstruct: Get https://proxy.golang.org/github.com/davidrjenni/reftools/cmd/fillstruct/@v/list: dial tcp 172.217.27.145:443: i/o timeout
go get github.com/davidrjenni/reftools/cmd/fillstruct: module github.com/davidrjenni/reftools/cmd/fillstruct: Get https://proxy.golang.org/github.com/davidrjenni/reftools/cmd/fillstruct/@v/list: dial tcp 172.217.27.145:443: i/o timeout

goplay:
Error: Command failed: /usr/local/bin/go get -v github.com/haya14busa/goplay/cmd/goplay
go get github.com/haya14busa/goplay/cmd/goplay: module github.com/haya14busa/goplay/cmd/goplay: Get https://proxy.golang.org/github.com/haya14busa/goplay/cmd/goplay/@v/list: dial tcp 172.217.27.145:443: i/o timeout
go get github.com/haya14busa/goplay/cmd/goplay: module github.com/haya14busa/goplay/cmd/goplay: Get https://proxy.golang.org/github.com/haya14busa/goplay/cmd/goplay/@v/list: dial tcp 172.217.27.145:443: i/o timeout

godoctor:
Error: Command failed: /usr/local/bin/go get -v github.com/godoctor/godoctor
go get github.com/godoctor/godoctor: module github.com/godoctor/godoctor: Get https://proxy.golang.org/github.com/godoctor/godoctor/@v/list: dial tcp 172.217.27.145:443: i/o timeout
go get github.com/godoctor/godoctor: module github.com/godoctor/godoctor: Get https://proxy.golang.org/github.com/godoctor/godoctor/@v/list: dial tcp 172.217.27.145:443: i/o timeout

gocode-gomod:
Error: Command failed: /usr/local/bin/go get -v -d github.com/stamblerre/gocode
go get github.com/stamblerre/gocode: module github.com/stamblerre/gocode: Get https://proxy.golang.org/github.com/stamblerre/gocode/@v/list: dial tcp 172.217.24.17:443: i/o timeout
go get github.com/stamblerre/gocode: module github.com/stamblerre/gocode: Get https://proxy.golang.org/github.com/stamblerre/gocode/@v/list: dial tcp 172.217.24.17:443: i/o timeout

godef:
Error: Command failed: /usr/local/bin/go get -v github.com/rogpeppe/godef
go get github.com/rogpeppe/godef: module github.com/rogpeppe/godef: Get https://proxy.golang.org/github.com/rogpeppe/godef/@v/list: dial tcp 172.217.24.17:443: i/o timeout
go get github.com/rogpeppe/godef: module github.com/rogpeppe/godef: Get https://proxy.golang.org/github.com/rogpeppe/godef/@v/list: dial tcp 172.217.24.17:443: i/o timeout

goreturns:
Error: Command failed: /usr/local/bin/go get -v github.com/sqs/goreturns
go get github.com/sqs/goreturns: module github.com/sqs/goreturns: Get https://proxy.golang.org/github.com/sqs/goreturns/@v/list: dial tcp 172.217.24.17:443: i/o timeout
go get github.com/sqs/goreturns: module github.com/sqs/goreturns: Get https://proxy.golang.org/github.com/sqs/goreturns/@v/list: dial tcp 172.217.24.17:443: i/o timeout

golint:
Error: Command failed: /usr/local/bin/go get -v golang.org/x/lint/golint
go get golang.org/x/lint/golint: module golang.org/x/lint/golint: Get https://proxy.golang.org/golang.org/x/lint/golint/@v/list: dial tcp 172.217.24.17:443: i/o timeout
go get golang.org/x/lint/golint: module golang.org/x/lint/golint: Get https://proxy.golang.org/golang.org/x/lint/golint/@v/list: dial tcp 172.217.24.17:443: i/o timeout

