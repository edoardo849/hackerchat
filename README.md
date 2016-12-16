# go-chat

HackerChat is a revolutionary idea comparable only to the invention of hot water or the re-engineering of the wheel. HackerChat will solve any communication issue by putting a chat right where you need it: right in the terminal. On top of that, HackerChat comes with state-of-the-art filters to facilitate communication and reduce professional and cultural misunderstandings.

Over-engineered, under-tested and completely unstable, HackerChat has the security of a V8 mounted on a shopping cart: a great experience for your CPU and absolute hell for your users.

*Make development great again!*

## Team Members
- Dan Bond
- Tim Blackwell
- Edo Scalafiotti

# Requirements
- Docker & Docker Compose
- Go >= 1.7

# Installation

```bash

git clone https://github.com/edoardo849/hackerchat

# Install dependencies
go get ./...

# OR, install dependencies with glide
glide install

# Edit main.go to include the IP of your Kafka instance and modify accordingly docker-compose.yml

# Run the docker container with Kafka (locally or from a remote server)
make kafka-up

# Run the app
make run

# Kill Docker
make kafka-down

```
