
```
root@c02bbd2123ed:/opt/projects# go -v
flag provided but not defined: -v
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        work        workspace maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.
```

```
root@c02bbd2123ed:/opt/projects# mkdir myapp                       
root@c02bbd2123ed:/opt/projects# cd myapp/
root@c02bbd2123ed:/opt/projects/myapp# go mod init myapp
go: creating new go.mod: module myapp
root@c02bbd2123ed:/opt/projects/myapp# go get github.com/labstack/echo/v4
go: downloading github.com/labstack/echo/v4 v4.11.4
go: downloading github.com/labstack/echo v3.3.10+incompatible
go: downloading github.com/labstack/gommon v0.4.2
go: downloading golang.org/x/crypto v0.17.0
go: downloading golang.org/x/net v0.19.0
go: downloading github.com/mattn/go-isatty v0.0.20
go: downloading github.com/valyala/fasttemplate v1.2.2
go: downloading github.com/mattn/go-colorable v0.1.13
go: downloading golang.org/x/sys v0.15.0
go: downloading golang.org/x/text v0.14.0
go: downloading github.com/valyala/bytebufferpool v1.0.0
go: added github.com/labstack/echo/v4 v4.11.4
go: added github.com/labstack/gommon v0.4.2
go: added github.com/labstack/echo/v4 v4.11.4
go: added github.com/labstack/gommon v0.4.2
go: added github.com/mattn/go-colorable v0.1.13
go: added github.com/mattn/go-isatty v0.0.20
go: added github.com/valyala/bytebufferpool v1.0.0
go: added github.com/valyala/fasttemplate v1.2.2
go: added golang.org/x/crypto v0.17.0
go: added golang.org/x/net v0.19.0
go: added golang.org/x/sys v0.15.0
go: added golang.org/x/text v0.14.0
root@c02bbd2123ed:/opt/projects/myapp# go run server.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.11.4
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
^Csignal: interrupt
root@c02bbd2123ed:/opt/projects/myapp# 
```