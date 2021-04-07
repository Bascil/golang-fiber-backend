# golang-fiber-backend-with-roles-and-permissions

A Go (golang) backend application developed on Fiber Framework

> This project is intended for developers who wish to expand their skills in using go to develop enterprise applications. This application uses MySQL database

It can be used as a fully functional backend with a working authentication system as well as user management with roles and permission

## Installation

Run the following command as a user with sudo privileges to download and extract the Go binary archive in the /usr/local directory

```shell
wget -c https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
```

Setup environment variables on your .bashrc or .bash_profile as follows:

```shell
export GOPATH=$HOME/path/to/your/project
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```
## Environment Setup

Add your environment variables to your `.env` file in the root of your project:

```shell
DB_NAME=
DB_USER=
DB_PASS=
JWT_SECRET=
```

Then in your Go application you can do something like

```go
package main

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  dbName := os.Getenv("DB_NAME")
  dbUser:= os.Getenv("DB_USER")
}
```

## Install Dependencies

Install the neccessary go dependencies. Make sure you have set GOPATH and PATH variables correctly

```shell
go get github.com/gofiber/fiber
go get github.com/gofiber/fiber/middleware/cors
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/dgrijalva/jwt-go
go get encoding/csv
```

## Watching File Changes

We will use realize as a dev dependency to watch file changes.

```shell
go get github.com/oxequa/realize

```
### Update realize.yml file

We will use realize as a dev dependency to watch file changes. Replace go-demos with the name of your go project folder

```shell
settings:
  legacy:
    force: false
    interval: 0s
schema:
- name: backend
  path: .
  commands:
    run:
      status: true
  watcher:
    extensions:
    - go
    paths:
    - /
    ignored_paths:
    - .git
    - .realize
    - vendor
```

Go to the root of your project and run
```shell
realize start

```

## Contributing

Contributions are most welcome! 
*code changes without tests will not be accepted*

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


