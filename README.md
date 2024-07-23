# A Web App with Golang and Gin Framework

This is a simple educational web app some main features are:
- Written in **Golang** and **Gin** framework
- Command line support like help, version and serve by **kobra** package.

# Installation
After cloning and install dependencies, for runnig app you need run mysql db too, for example:
```
docker run --name mysql-db -e MYSQL_ROOT_PASSWORD=mysql_root_password -e MYSQL_DATABASE=mysql_db -e MYSQL_USER=mysql_user -e MYSQL_PASSWORD=mysql_password -p 3306:3306 -d mysql
```


# Usage
After installation, provide environ variables in a .env file according to env.example.Then to see all available commands of this app type:
`go run main.go ` you should see something like this:

```
This app leverage fast gin framework to serve a blog sample web application.
        You can see source code at https://github.com/RasoulZamani/hiGin

Usage:
  help [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  serve       Run server and serve web app
  version     Print the version number of App

Flags:
  -h, --help   help for help

Use "help [command] --help" for more information about a command.
```

For example to run web app: `go run main.go serve` and then go to the host:port you had declared in .env.