SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiNzJMNUJAUlVUSEkuY29tIiwidXVpZCI6Ijg2M2RmNWZmLTgyNTgtNDdjYS04YWJkLTUwYjVhYjkwMzM1ZiJ9LCJ1bmFtZSI6IjcyTDVCQFJVVEhJLmNvbSIsInRva2VuX2lkIjoiODYzZGY1ZmYtODI1OC00N2NhLThhYmQtNTBiNWFiOTAzMzVmX2czRlJndWFvNmp6cDJhY081ZmZuOEVWVDFZdlU0Uk5RdzQwSHRnT1AvdVpkVlZFdlNiVmxQSXI5eE0xV1kxeG5FOG9PNmRTM0crVkgvUG9nU2J3OFNQV1FVL3g1LzI2alNMNFF1eStZVm55M0ZFTzZsQndqZVV2bDBld2hHc2hJeEl1WlVnPT1fMTY0NzgzNjg0MzQ2MiIsImV4cCI6MTY0Nzg0MDQ0MzQ2MiwiaXNzdWVkX2F0IjoxNjQ3ODM2ODQzNDYyLCJjc3JmX3Rva2VuIjoiIn19.KnrFNb-U8xTLzJq-RLO7JjtXt9H5OfN1ZTAZlUGoRlI"

csrf_token="249GjlqaJFRnZTekttQtxNrkAs8HMT3ORZO2JT8w0CDzU3kYOCZ8xvQvF6m3io9yUkPzgCxkrr21YUrrbdL0/AhSQMy32ovXF5ZP2QnXkZSC2Nzoz4jrBf4FbYfZKKinT63zow=="

echo -e "\n====> /api/login_status/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/login_status/" 

echo -e "\n\n====>/api/user/\n"
curl  -i --cookie  "access_token=$access_token" --header "csrf_token: $csrf_token"  "http://localhost:8000/api/user/" 

echo -e "\n\n\n\n\n\n\n\n"

ulimit -n 1000000


ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"



