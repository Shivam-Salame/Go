# Go

cmds:

go env
go mod init <module-name>
go get -u <package-name>
go mod verify

# list dependent packages
go list 

# list all packages
go list all

# list versions of package
go list -m -versions github.com/gorilla/mux

# 
go mod tidy

#
go mod why github.com/gorilla/mux

#
go mod graph

# change version of go
go mod edit -go 1.16

# change version of module
go mod edit -module 1.16

# to have all modules locally like node modules
go mod vendor

# to pull packages from locally
go run -mod=vendor main.go
