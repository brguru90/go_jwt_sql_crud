SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiYWEiLCJ1dWlkIjoiMWU2YjNiZGItNmM5MS00ZGVlLWIzOGUtZDYwMTVjMDIxZDg4In0sInVuYW1lIjoiYWEiLCJ0b2tlbl9pZCI6IjFlNmIzYmRiLTZjOTEtNGRlZS1iMzhlLWQ2MDE1YzAyMWQ4OF8wWFBKSjlKY2RvTkxQdlgxeW1WOVV5NnVmbzA5VFFodlRUUjc0QStTTitKb3lmdy9ZQnlGbmw2dWQ4aWxpb2M5cGtDa0t1cGpkVFBSOVM5c2hZYmxZTnBhSWpBOE0wYy9ZSWlBbFBQcHEvT0FQcVZsRWlVWjdrZGwybHNNS1pZb0NjZ2I1Zz09XzE2NDY5Mzk4MDcyNjgiLCJleHAiOjE2NDY5NDM0MDcyNjgsImlzc3VlZF9hdCI6MTY0NjkzOTgwNzI2OCwiY3NyZl90b2tlbiI6IiJ9fQ.zB7TRxbrgmxX6QnZBPXkKZERxKEaV88zQ0pOoe2_FAY"

csrf_token="uLuPFkT0TnERfD/iG2ou/DqLCF2JBFofUO1Rh33zsiy5k/6oiyZM8pNXc45cMX0xdHQh+Fr9XXH/hi2KJNLHW/V4ZWClcgO+TF5SOC7UvwJWVJW9RZxPc4t/tDMeuMWrpzVCDA=="

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"

