# Development_Enviroment_Setting_1
#### Urls

```
https://golag.org/
https://atom.io/
https://git-scm.com/downloads
```

#### HomeBrew

* Go-lang
```
> brew install go

> go version
go version go1.12.7 darwin/amd64

> godoc -http=:8080
http://localhost:8080
```

* Atom
```
> brew cask install atom

```
* Git
```
> brew install git

> git --version
git version 2.22.0

```


#### Make Directories
```
Go-Tutorial
    - bin
    - pkg
    - src
```

#### Setting GOBIN, GOROOT - Free Setting
```
> vim ~/.bash_profile 

export GOPATH="/Users/adminstrator/Documents/Desktop/Go-Tutorial"
export GOBIN=$GOPATH/bin

export PATH=$GOPATH:$GOBIN

> source ~/.bash_profile
```