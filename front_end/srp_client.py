import srp
import requests
import sys
import json
username = sys.argv[1]
password = sys.argv[2]


## Make the Sign up Call
requests.post("http://localhost:8081/lb/signup",json={
    "Username": username,
    "Password": password
})

## Make the login call

is_login = requests.post("http://localhost:8081/lb/login",json={
    "Username": username,
    "Password": password
})
print(is_login)
## make the authentication request

