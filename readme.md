# Welcome to Quoverflow!

This is an API example written in go.

## Database Setup

This project uses a PostgreSQL database and needs to configure these Environment Variables:

- DATABASE_USER
- DATABASE_PASS
- DATABASE_URL
- DATABASE_PORT
- DATABASE_NAME

### Notes

The go.mod file was created with "go mod init" command

## Architecture

This project is an approach to "Clean Architecture" with the necessary layers to build complex enterprise applications.

### Layers
- Domain: Models & Repositories
- Application: Services
- Infrastructure: Configurations, Controllers, Middleware & Routes

