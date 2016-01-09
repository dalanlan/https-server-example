#!/bin/bash

# one-way authentication
# generate private key file (server)
# 2048 -- key length
openssl genrsa -out server.key 2048

# generate public key file (server)
# openssl rsa -in server.key -out server.key.public

# generate digital cert (server)
openssl req -new -x509 -key server.key -out server.crt -days 365



