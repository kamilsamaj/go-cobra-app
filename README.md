# go-cobra-app
Go CLI application built by [cobra](https://github.com/spf13/cobra)

# Prerequisites
* Install `gox` for easy cross-compilation with `go get -u github.com/mitchellh/gox`
* Optional: To enable `pre-commit` checks, install https://pre-commit.com/ and run `pre-commit install`

```shell
cobra init go-cobra-app --author "Kamil Samaj <kamilsamaj@gmail.com>" --pkg-name github.com/kamilsamaj/go-cobra-app
go mod init github.com/kamilsamaj/go-cobra-app
pre-commit install
```

# References
* https://github.com/spf13/cobra/tree/master/cobra
