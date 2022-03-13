SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiYWEiLCJ1dWlkIjoiZTVlZDAxOWUtODJkYi00ZTgxLTlmNjQtNjkzMWRkNzI2MjFjIn0sInVuYW1lIjoiYWEiLCJ0b2tlbl9pZCI6ImU1ZWQwMTllLTgyZGItNGU4MS05ZjY0LTY5MzFkZDcyNjIxY19UaFpDNTkvOUlRbWo1ZEtJcUtYbEVXU1VhUlRlL0RqQzE2UVBhbEsya1kxcDZoaXBGYWxiMjBNOHpXK1ZtZFEyQ3ZFcHR6RVdTNlBiY3VPUGcwci9wamRKQ0k1b0QwSzlIWUJ6TXZQU2ZTOXVzODhnbTV3RGVhZFRZWTZyanhtVEF6bjV6dz09XzE2NDcxNzMxMDI1MjUiLCJleHAiOjE2NDcxNzY3MDI1MjUsImlzc3VlZF9hdCI6MTY0NzE3MzEwMjUyNSwiY3NyZl90b2tlbiI6IiJ9fQ.Nh9qoR-rY7vHieGPpMTVZTTHM-eXpMTcWjTjWVZMoOI"

csrf_token="48XR+KizmNBn+FeEdxiD2g/dGRR6KVgbMKzpIig/kx3xIzBnJjt+0qKY9qYngid2ub0+6v88OYC11jYpZk6/WoK/P2mfYLZfOf0gp6g29OYJOD/KXs32k7263h8FVy+bXN7Ynw=="

echo -e "\n====> /api/login_status/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/login_status/" 

echo -e "\n\n====>/api/user/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/user/" 

echo -e "\n\n\n\n\n\n\n\n"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"



