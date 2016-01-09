#!/bin/bash

# client verify server

# generate CA private key
openssl genrsa -out ca.key 2048

# generate digital certificates
openssl req -x509 -new -nodes -key ca.key -subj "/CN=*.tunnel.tonybai.com" -days 5000 -out ca.crt

# generate server private key
openssl genrsa -out serverBi.key 2048

# generate server digital cert request file *.csr
openssl req -new -key serverBi.key -subj "/CN=localhost" -out serverBi.csr

# generate server digital cert
openssl x509 -req -in serverBi.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out serverBi.crt -days 5000

# server verify client

# generate client private key
openssl genrsa -out client.key 2048

# generate client digital cert
openssl req -new -key client.key -subj "/CN=tonybai_cn" -out client.csr
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 5000

# regenerate client cert with Extend key usage enclosed
cat <<EOF > client.ext 
extendedKeyUsage=clientAuth
EOF
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile client.ext -out client.crt -days 5000