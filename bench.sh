SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiRjQ0SkFASFA3U0IuY29tIiwidXVpZCI6Ijc5NjUyYjYyLTJjYmUtNGFmNy05OWJiLWZjMjkzMDVhYTc1MSJ9LCJ1bmFtZSI6IkY0NEpBQEhQN1NCLmNvbSIsInRva2VuX2lkIjoiNzk2NTJiNjItMmNiZS00YWY3LTk5YmItZmMyOTMwNWFhNzUxX0c2OUtOZVF4Uzk3QzZLUDA2Zmd5Ui9lRkVENDh5L3JSNXlNbGRhVnlrUWFvd0ZJdjF1Vlp1R0ZWdXp2VjdYSWpocXJOVkVmTUNxV212TjJ0aGprTVliK2wwd1NtbTdTcmlGcjBsS2R5MEtuazlVOGpTUEloLzdMNVR0YUFRaE9vWTJjd3hBPT1fMTY0Njg5NDg4NzE2MyIsImV4cCI6MTY0Njg5ODQ4NzE2MywiaXNzdWVkX2F0IjoxNjQ2ODk0ODg3MTYzLCJjc3JmX3Rva2VuIjoiIn19.8WImv5hiEznBhYxtb8jIx5RzQsWIUVfk1sNOtxszySU"

csrf_token="oV9hTM5mbJ9acU9ms7c5IqrwSH66nvB0zjN+VVH2LEmW77ixoCEQXk2gfpt7h/CKKpBYyL3WgASqpoQv4vxc6ex/uMlzpR8d2AX5PUs9IkLRpyKKJoKKJ/WajQwkOp1cJ6LJzQ=="

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"

