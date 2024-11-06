#/bin/bash
set -e

# CreateUser
curl -X POST http://localhost:8080/users \
    -H "Content-Type: application/json" \
    -d '{"Name": "hoge", "Email": "hoge@example.com"}'

# GetUser
curl http://localhost:8080/users/1