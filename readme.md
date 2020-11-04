# Welcome to Quoverflow!

This is an API example written in go.

## Architecture

This project is an approach to "Clean Architecture" with the necessary layers to build complex enterprise applications.

### Layers
- Domain: Models & Repositories
- Application: Services
- Infrastructure: Configurations, Controllers, Middleware & Routes

### Requirements
- PostgreSQL database

### Set DB Environment Variables
- DATABASE_USER
- DATABASE_PASS
- DATABASE_URL
- DATABASE_PORT
- DATABASE_NAME

## Setup dependencies
$ go mod init
$ go mod vendor

### For running locally
$ go run go.main
Then API V1 will be availability on /v1/health

### For Swagger documentation 
$ go get -u github.com/swaggo/swag/cmd/swag
$ swag init
- Documentation is accessible on /swagger/index.html

### For Deployment on Heroku
1. Set DB environment variables
2. Create heroku app
$ heroku create quoverflow
  or if app already exists... $ heroku git:remote -a quoverflow
3. Push to heroku
$ git push heroku main

### For Deployment on ElasticBeanstalk
1. Set DB environment variables
2. Set System environment variables 
- APP_DEPLOY_DIR = /var/app/current
- APP_STAGING_DIR = /var/app/staging
- GIN_MODE = release
- GOPATH = /var/app/current
- PATH = /bin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:/usr/local/go/bin:/var/app/current
3. Push to ElasticBeanstalk from aws-cli
$ eb deploy "your-environmen" --profile "your-profile"

### Notes
- For deployment on ElasticBeanstalk rename main.go to application.go