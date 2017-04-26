# lazygit

*lazygit* is a short-hand command for doing 

```
git add -A
git add -u
git commit -m "SOME MESSAGE"
git pull
git push
```

It is a dangerous practice, but it saves a lot of time. Also it encourages [committing early and often](http://www.databasically.com/2011/03/14/git-commit-early-commit-often/).

### Usage

```
lazygit Your commit message
```

Note that you don't have to put quotes around the commit message

### Installation

```
go get -u github.com/JoshNavi/lazygit
cd lazygit
go build
```

Should work with most versions of go

### Credit

This is just a golang port of @gka's [git-go nodejs script](https://github.com/gka/git-go).