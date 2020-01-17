import requests
import json

body = {
    "Value": 6
}
r = requests.post("http://localhost:8080", data=json.dumps(body))
print(r.text)

