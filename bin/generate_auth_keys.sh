#!/usr/bin/env bash
openssl genrsa -out educonn_auth.rsa 4096
openssl rsa -in educonn_auth.rsa -pubout > educonn_auth.rsa.pub
