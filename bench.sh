SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiN0I1WldAREk4MzkuY29tIiwidXVpZCI6IjFjYjZkMjhjLWEyZWMtNDU1Yy04Y2U3LWQ1NWM5YjUzMmFiMiJ9LCJ1bmFtZSI6IjdCNVpXQERJODM5LmNvbSIsInRva2VuX2lkIjoiMWNiNmQyOGMtYTJlYy00NTVjLThjZTctZDU1YzliNTMyYWIyXzhxNWZHN2I1TDFoT05HWTJ1a0pSVG4yVGJKbW5kZkt0bklSaWJxMEVheW1FMi9PRVNibFcwM3V4cWpFNm5wOFREWmNPbTdHU2FQNFVBbnBaUlBwVzJBbUNleHh5VHhWb3orTEU1cXNXYWtWZS9wV1VIQkZ3emxjajFYRW1iWTBBN2pSRHZ3PT1fMTY0Njg1ODEyODUyMSIsImV4cCI6MTY0Njg2MTcyODUyMSwiaXNzdWVkX2F0IjoxNjQ2ODU4MTI4NTIxLCJjc3JmX3Rva2VuIjoiIn19.a9ATaGmleHUKeabBbz4axWG2ao4kGQQ36lkH6FdzENM"

csrf_token="1byp7KXT3Oh5YrOomAfJgHPLHgPD+t6vrYfPX0nO2jT0LyaZSIamYerE447rEDZg97w4E4L+uF4dYHpJdwmHJ/DavOyoB0lPXZi7RDKrucIZPUx6sr8EeEWXCSFCAYkDlqoxnw=="

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"

