import requests
import random
import string

USERS = []


def rnd_name():
    return ''.join(random.choice(string.ascii_uppercase + string.digits) for _ in range(10))


def rnd_email():
    return rnd_name() + "@" + rnd_name() + ".ru"


for i in range(5):
    USERS.append({"name": rnd_name(), "pwd": rnd_name(), "email": rnd_email(), "token": ""})


# Normal registration
try:
    for u in USERS:
        data = {"email": u["email"], "password1": u["pwd"], "password2": u["pwd"], "name": u["name"]}
        r = requests.post('http://127.0.0.1:8001/api/sign-up', json=data)
        print(r.json())
        if r.json()["error"] != "":
            raise Exception(r.json()["error"])

    print("Sign up - OK")
except BaseException as e:
    print("Sign up - failed:", e)


# Wrong json
try:
    for u in USERS:
        data = {"email": 123, "password1": 123, "password2": 123, "name": 123}
        r = requests.post('http://127.0.0.1:8001/api/sign-up', json=data)
        print(r.json())
        if "cannot unmarshal" not in r.json()["error"]:
            raise Exception(r.json()["error"])

    print("Sign up / json error - OK")
except BaseException as e:
    print("Sign up / json error - FAIL:", e)


# Wrong email
try:
    for u in USERS:
        data = {"email": u["name"], "password1": u["pwd"], "password2": u["pwd"], "name": u["name"]}
        r = requests.post('http://127.0.0.1:8001/api/sign-up', json=data)
        print(r.json())
        if r.json()["error"] != "valid email address is required":
            raise Exception(r.json()["error"])

    print("Sign up / valid email address is required - OK")
except BaseException as e:
    print("Sign up / valid email address is required - FAIL:", e)

# Wrong password
try:
    for u in USERS:
        data = {"email": u["email"], "password1": u["pwd"], "password2": u["pwd"]+u["pwd"], "name": u["name"]}
        r = requests.post('http://127.0.0.1:8001/api/sign-up', json=data)
        print(r.json())
        if r.json()["error"] != "password and confirmation is not equal":
            raise Exception(r.json()["error"])

    print("Sign up / password and confirmation is not equal - OK")
except BaseException as e:
    print("Sign up / password and confirmation is not equal - FAIL:", e)

# Wrong name
try:
    for u in USERS:
        data = {"email": u["email"], "password1": u["pwd"], "password2": u["pwd"], "name": "X"}
        r = requests.post('http://127.0.0.1:8001/api/sign-up', json=data)
        print(r.json())
        if r.json()["error"] != "user name is required":
            raise Exception(r.json()["error"])

    print("Sign up / user name is required - OK")
except BaseException as e:
    print("Sign up / user name is required - FAIL:", e)


# Wrong email (used)
try:
    for u in USERS:
        data = {"email": u["email"], "password1": u["pwd"], "password2": u["pwd"], "name": u["name"]}
        r = requests.post('http://127.0.0.1:8001/api/sign-up', json=data)
        print(r.json())
        if r.json()["error"] != "email already used":
            raise Exception(r.json()["error"])

    print("Sign up / email already used - OK")
except BaseException as e:
    print("Sign up / email already used - FAIL:", e)


# Sign in
try:
    for idx in range(len(USERS)):
        u = USERS[idx]
        data = {"email": u["email"], "password": u["pwd"]}
        r = requests.post('http://127.0.0.1:8001/api/sign-in', json=data)
        print(r.json())
        USERS[idx]["token"] = r.json()["data"]["Token"]

    print("Sign in - OK")
except BaseException as e:
    print("Sign in - FAIL:", e)


# Wrong auth
try:
    for u in USERS:
        data = {"email": u["email"], "password": u["pwd"]+u["pwd"]}
        r = requests.post('http://127.0.0.1:8001/api/sign-in', json=data)
        print(r.json())
        if r.json()["error"] != "login and password does not match":
            raise Exception(r.json()["error"])

    print("Sign in / login and password does not match - OK")
except BaseException as e:
    print("Sign in / login and password does not match - FAIL:", e)


# Profile
try:
    for u in USERS:
        r = requests.get('http://127.0.0.1:8001/api/profile', headers={'Authorization': 'access_token ' + u["token"]})
        print(r.json())
        if r.json()["data"]["email"] != u["email"]:
            raise Exception("data not equal")

    print("Profile - OK")
except BaseException as e:
    print("Profile - FAIL:", e)


# Profile - wrong token
try:
    for u in USERS:
        r = requests.get('http://127.0.0.1:8001/api/profile', headers={'Authorization': 'access_token fake.token.jwt'})
        print(r.json())
        if r.json()["error"] != "Malformed authentication token":
            raise Exception(r.json()["error"])

    print("Profile / Malformed authentication token- OK")
except BaseException as e:
    print("Profile / Malformed authentication token- FAIL:", e)


# Auto complete
try:
    for u in USERS:
        data = {"name": u["name"][:5]}
        r = requests.post('http://127.0.0.1:8001/api/transfer/autocomplete', headers={'Authorization': 'access_token ' + u["token"]}, json=data)
        print(r.json())
        if len(r.json()["data"]) < 1:
            raise Exception("User not found")

    print("Auto complete- OK")
except BaseException as e:
    print("Auto complete- FAIL:", e)


# Transfer
try:
    recipient = USERS[0]
    for u in USERS[1:]:
        data = {"recipient": recipient["email"], "amount": 400}
        r = requests.post('http://127.0.0.1:8001/api/transfer/create', headers={'Authorization': 'access_token ' + u["token"]}, json=data)
        if r.json()["error"] != "":
            raise Exception(r.json()["error"])

    print("Transfer - OK")
except BaseException as e:
    print("Transfer - FAIL:", e)


# Transfer - low balance
try:
    recipient = USERS[0]
    for u in USERS[1:]:
        data = {"recipient": recipient["email"], "amount": 400}
        r = requests.post('http://127.0.0.1:8001/api/transfer/create', headers={'Authorization': 'access_token ' + u["token"]}, json=data)
        print(r.json())
        if r.json()["error"] != "balance is too low":
            raise Exception(r.json()["error"])

    print("Transfer / balance is too low - OK")
except BaseException as e:
    print("Transfer / balance is too low - FAIL:", e)


# History
try:
    for u in USERS:
        data = {
            "timestampMin": 1548618776, "timestampMax": None,
            "amountMin": None, "amountMax": 500, "Name": u["name"][:1],
            "sort": "date", "size": 10, "offset": 0}
        r = requests.post('http://127.0.0.1:8001/api/transfer/history', headers={'Authorization': 'access_token ' + u["token"]}, json=data)
        print(r.json())
        if len(r.json()["data"]) < 1:
            raise Exception(r.json()["error"])

    print("History - OK")
except BaseException as e:
    print("History - FAIL:", e)

