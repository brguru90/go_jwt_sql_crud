SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiN0I1WldAREk4MzkuY29tIiwidXVpZCI6IjFjYjZkMjhjLWEyZWMtNDU1Yy04Y2U3LWQ1NWM5YjUzMmFiMiJ9LCJ1bmFtZSI6IjdCNVpXQERJODM5LmNvbSIsInRva2VuX2lkIjoiMWNiNmQyOGMtYTJlYy00NTVjLThjZTctZDU1YzliNTMyYWIyX2NwOFdycGlKMjZvdVhxZHc0RFd1OU00emVQM3V2VmZVbDdiVTc5K0I2OUZrUjJUZTlLOVhQbU1ZRGxqQWRDb3cxQkFxNFN6QVRiT2hKTGVqZmh4KzZBUkV1RjQzem5mRi9CU2d3Sm9SOUdHbWtsR2RwVFhFbjlSYm1KK1c1ajhJZHpRWjZnPT1fMTY0Njg1MzY3MTgyNCIsImV4cCI6MTY0Njg1NzI3MTgyNCwiaXNzdWVkX2F0IjoxNjQ2ODUzNjcxODI0LCJjc3JmX3Rva2VuIjoiIn19.BIy5Y136ZxdbONb8icmmCbZgM6nYcdrDXyrP-Hj2jRI"

csrf_token="9YAEgzpEFybPuhDBPnEfG1yyPZuqlneImk885XFpUaJKjOTiiOWEomE3eDuIZWT3CJBKX4lBppTaeLsoYaHrxDL+r4c6xguBB7iAtk5C5wgVaBVnvpwqVaLttanDXUzbX5gswQ=="

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"

