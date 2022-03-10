SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiYWEiLCJ1dWlkIjoiMWU2YjNiZGItNmM5MS00ZGVlLWIzOGUtZDYwMTVjMDIxZDg4In0sInVuYW1lIjoiYWEiLCJ0b2tlbl9pZCI6IjFlNmIzYmRiLTZjOTEtNGRlZS1iMzhlLWQ2MDE1YzAyMWQ4OF82WlpIbW1SOWZ5a3JtSHpsTW9jMHQvSFNXTmY3OUtSNjVIbUxZMnJiNVJNcUJKMlJoUDdhWUJxNDc2Z3NjbzRJemQzYmVyYVFFWUtOWEFtOHdsS05meGV0V0RrcUdadE5KRGlEYU1zVzh6ZENvMkE3QThER1UyL1RTeVRXNVNBNW1Hcjh3UT09XzE2NDY5MzQ3MzQ4MjUiLCJleHAiOjE2NDY5MzgzMzQ4MjUsImlzc3VlZF9hdCI6MTY0NjkzNDczNDgyNSwiY3NyZl90b2tlbiI6IiJ9fQ.fIROpKeTm_iWhsWLA7ns6mjSGyRijOd_PPzyjdci1mU"

csrf_token="8Q7t8ms+Bo6eJmsYf/yE+ErgK7OvRyp98xxBCCTRqpc1X23CGfiHyqkp0o2GOlr5fEQIowMouXxa5msRxqavYID4K4sob6ZbCmXaLd9ArHquw5QY7t0JVB0lr043Y27V4FJgZg=="

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"

