# auth-server-go
This is a simple auth server implementation in Go

### Basic Introduction

The aim of this project is to learn about the go eco-system and server code in go and implement the basic authentication using the SRP protocal where we do not store the password at all.

SRP is a protocol where we do not store user password anywhere and only store a random salt to derive a hash which would be same for the combination of password and username.

Advantages:
Disadvantages: