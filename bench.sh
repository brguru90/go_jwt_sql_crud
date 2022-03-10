SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiYWEiLCJ1dWlkIjoiMWU2YjNiZGItNmM5MS00ZGVlLWIzOGUtZDYwMTVjMDIxZDg4In0sInVuYW1lIjoiYWEiLCJ0b2tlbl9pZCI6IjFlNmIzYmRiLTZjOTEtNGRlZS1iMzhlLWQ2MDE1YzAyMWQ4OF9CN1ZHZ3JwSG0yeG9za3liZVUzdWwrbzZuUzhpRktNR2dHZ2dxODh4MzlYU2FPREE0amxDbFRRTFlwSFZkSVdmcmZpeFhCRG5WK3dndXBlZUQzbjVxditRNk5iNjBqMEJ4Z0FlUzR4VlNMZUozSjJVb0JEWmFGSTYrL0FNZUxYVWMzaTZCUT09XzE2NDY5NDQ1NzcxODEiLCJleHAiOjE2NDY5NDgxNzcxODEsImlzc3VlZF9hdCI6MTY0Njk0NDU3NzE4MSwiY3NyZl90b2tlbiI6IiJ9fQ.Y2bte2hnPP_rTAuBtYg38t-FrSQ0Z3q9kRITEy1rhZo"

csrf_token="ohnzkw+e0hlvYi7C3697m1jrkbE8EZqWkCRbOxnepbJtdum4aWL0ddcUi8xWHGYwo1oUUxFRTWXqKU1knwP4kSVQYT6vhm0+XiyEkwVpzxBGkGjA2fw/TgiSEURdyYlJ7sRWgA=="

echo -e "\n====> /api/login_status/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/login_status/" 

echo -e "\n\n====>/api/user/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/user/" 

echo -e "\n\n\n\n\n\n\n\n"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"



