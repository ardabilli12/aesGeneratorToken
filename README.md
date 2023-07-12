# AES Generator Token CLI

The AES Generator Token CLI is a command-line tool written in Go that generates AES tokens for encrypting Company and User (paper internal usage only).

## Prerequisites

- Go version 1.16 or higher installed
- MySQL database credentials (username, password, host, port, and database name) in a `.env` file

## Configuration

Before running the CLI, make sure to set up the necessary environment variables in the `.env` file. The following variables are required:

- `MYSQL_USERNAME`: The username for the MySQL database
- `MYSQL_PASSWORD`: The password for the MySQL database
- `MYSQL_HOST`: The host address of the MySQL database
- `MYSQL_PORT`: The port number of the MySQL database
- `MYSQL_DATABASE`: The name of the MySQL database

## Usage

The CLI has the following command:

### Generate

Generates an AES token for the specified email address.
> ~ go run main.go generate --email user.email@example.com

Replace `<email>` with the email address for which you want to generate the AES token.

## Troubleshooting

- If you encounter any issues, make sure you have the correct MySQL database credentials set up in the `.env` file.
- Double-check that you have Go version 1.16 or higher installed.

## Author
Ardabilli12 - (ahmad.ardabilli@paper.id)