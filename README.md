# go-template
template project using Golang , Cmake , docker easy to start
#Library use.
1. Logging: Logrus
2. Api: go-gin
3. Gomodule
4. Environment: Viper

# defined:
default module's name is "app", update in file `go.mod` (ignore editing file `go.sum`)
# using
1. From root folder. The place contains Makefile. run command
for production: `make install env=prd` <br>
for staging: `make install env=stg` <br>
for dev: `make install env=dev` <br>
for local: `make install env=local` <br>

# setup path for go
MacOS.
```
export GO111MODULE="on"
export GOPATH=$HOME/go-installed-folder
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOROOT/bin
```
