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
        if r.json()["error"] != "":
            raise Exception(r.json()["error"])

    print("Sign up - OK")
except BaseException as e:
    print("Sign up - failed:", e)


# Wrong email
try:
    for u in USERS:
        data = {"email": u["name"], "password1": u["pwd"], "password2": u["pwd"], "name": u["name"]}
        r = requests.post('http://127.0.0.1:8001/api/sign-up', json=data)
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
        USERS[idx]["token"] = r.json()["data"]["token"]

    print("Sign in - OK")
except BaseException as e:
    print("Sign in - FAIL:", e)


# Wrong auth
try:
    for u in USERS:
        data = {"email": u["email"], "password": u["pwd"]+u["pwd"]}
        r = requests.post('http://127.0.0.1:8001/api/sign-in', json=data)
        if r.json()["error"] != "login and password does not match":
            raise Exception(r.json()["error"])

    print("Sign in / login and password does not match - OK")
except BaseException as e:
    print("Sign in / login and password does not match - FAIL:", e)


# Profile
try:
    for u in USERS:
        r = requests.get('http://127.0.0.1:8001/api/profile', headers={'Authorization': 'access_token ' + u["token"]})
        if r.json()["data"]["email"] != u["email"]:
            raise Exception("data not equal")

    print("Profile - OK")
except BaseException as e:
    print("Profile - FAIL:", e)


# Profile - wrong token
try:
    for u in USERS:
        r = requests.get('http://127.0.0.1:8001/api/profile', headers={'Authorization': 'access_token fake.token.jwt'})
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
        if len(r.json()["data"]) < 1:
            raise Exception("User not found")

    print("Auto complete- OK")
except BaseException as e:
    print("Auto complete- FAIL:", e)


