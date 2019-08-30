#!/bin/sh
openssl ecparam -name prime256v1 -genkey -noout -out priv.pem
openssl ec -in priv.pem -pubout -out pubkey.pem
