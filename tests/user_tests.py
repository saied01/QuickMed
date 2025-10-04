import requests
import time

BASE_URL = "http://localhost:8081"
created_user_id = None

def test_signup():
    global created_user_id
    timestamp = int(time.time())
    payload = {
        "name": f"Test{timestamp}",
        "email": f"test{timestamp}@example.com",
        "password": "pass123",
        "age": 25,
    }
    r = requests.post(f"{BASE_URL}/signup", json=payload)
    assert r.status_code == 201

    # Ahora el handler devuelve el objeto user con id
    created_user_id = r.json()["ID"]  # depende de cómo serialices User en Go
    print("Signup OK, id =", created_user_id)

def test_get_user():
    global created_user_id
    r = requests.get(f"{BASE_URL}/users/{created_user_id}")
    assert r.status_code == 200
    print("Get user OK")

def test_delete_user():
    global created_user_id
    r = requests.delete(f"{BASE_URL}/users/{created_user_id}")
    assert r.status_code in (200, 204)
    print("Delete user OK")

if __name__ == "__main__":
    test_signup()
    test_get_user()
    test_delete_user()
