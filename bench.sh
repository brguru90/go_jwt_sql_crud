SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiYWEiLCJ1dWlkIjoiMDU0ZjY0MTMtOWNlNy00MjNiLTgxMzMtM2Y3YTFhN2FmMWJhIn0sInVuYW1lIjoiYWEiLCJ0b2tlbl9pZCI6IjA1NGY2NDEzLTljZTctNDIzYi04MTMzLTNmN2ExYTdhZjFiYV9naW1BVm5LMWkreEd3cG9uSEl0UkpCMmpBeWFCZUwza2hnUlhGTk1LWkUyVVIrYmE4MVNWU2tVenRwNVFlRTFlQ05rNEJhNXZmN1pWODdvd1pDQVJnZk4vWUpRcWJpKzRIT1hCbFp3U0ErUVY3SnNpdXE3Wk5wQkVrbW96Nm53UWVUYi9hdz09XzE2NDcxMDk5MzcyNjEiLCJleHAiOjE2NDcxMTM1MzcyNjEsImlzc3VlZF9hdCI6MTY0NzEwOTkzNzI2MSwiY3NyZl90b2tlbiI6IiJ9fQ.vvonY2evzgrDpmdqKwyfzhjxiK_1Ktskos46wG19ofU"

csrf_token="meIa4oguZ59K1rUv5MAJNfQr4nHkw6XQ0RWgop3Efqa2543FzHvamSNOn82OYrG7AEXpHDpUXbeSd4m/FiSXFa6sLL1Td8WtCa1y5VhdHDXJXnd55SvXwR8gdjFZOeleZnViqg=="

echo -e "\n====> /api/login_status/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/login_status/" 

echo -e "\n\n====>/api/user/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/user/" 

echo -e "\n\n\n\n\n\n\n\n"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"



