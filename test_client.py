import requests

data = {
    "name" : "Ace Drogon",
    "email" : "acedrogon@gmail.com",
    "password" : "test123",
    "password_confirmation" : "test123",
}

url = "http://localhost:8080/api/users"
res = requests.post(url, data=data)
if res.ok:
    print(res.json())
