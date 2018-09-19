# brunch-crawler
Brunch crawler in Go

## Setup

#### Install Go on MacOS X
```sh
# install via Homebrew
brew install go

# or download and install as MacOSX package from https://golang.org/dl/
curl -o go.pkg https://dl.google.com/go/go1.10.3.darwin-amd64.pkg
sudo open go.pkg
```

#### Configure GOPATH
```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
Recommended: In your `~/.profile` or `~/.bashrc` file, set the `$GOPATH` variable and add the `$GOPATH/bin` path to your `$PATH`.

#### Install Additional Dev Tools
```sh
# for golang lint
go get github.com/golang/lint/golint
# for docs
go get github.com/davecheney/godoc2md
# for testing
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen
```

#### Create Go App structure
```sh
mkdir -p $GOPATH/src/github.com/marcosdeseul
cd $GOPATH/src/github.com/marcosdeseul
git clone git@github.com:marcosdeseul/brunch-crawler.git
```

## Install package dependencies
```sh
go get -u github.com/golang/dep/cmd/dep
dep ensure
```
