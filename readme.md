# Welcome to Quoverflow!

This is an API example written in go.

## Architecture

This project is an approach to "Clean Architecture" with the necessary layers to build complex enterprise applications.

### Layers
- Domain: Models & Repositories
- Application: Services
- Infrastructure: Configurations, Controllers, Middleware & Routes

## Setup and Deployment to ElasticBeanstalk

## DB Environment Variables:
- DATABASE_USER
- DATABASE_PASS
- DATABASE_URL
- DATABASE_PORT
- DATABASE_NAME

## System Environment Variables:
- APP_DEPLOY_DIR = /var/app/current
- APP_STAGING_DIR = /var/app/staging
- GIN_MODE = release
- GOPATH = /var/app/current
- PATH = /bin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:/usr/local/go/bin:/var/app/current

### Notes
- This project uses a PostgreSQL database
- The go.mod file was created with "go mod init" command
- Swagger documentation was created with "swagger generate spec -o ./swagger.yalm --scan-models" command
- Documentation is accessible on /swagger/index.html
- API V1 availability on /v1/health  
