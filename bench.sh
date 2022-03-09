SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiN0I1WldAREk4MzkuY29tIiwidXVpZCI6IjFjYjZkMjhjLWEyZWMtNDU1Yy04Y2U3LWQ1NWM5YjUzMmFiMiJ9LCJ1bmFtZSI6IjdCNVpXQERJODM5LmNvbSIsInRva2VuX2lkIjoiMWNiNmQyOGMtYTJlYy00NTVjLThjZTctZDU1YzliNTMyYWIyX0c1eHc3ZzRySlhHRDZVaFlFVVA0OFBldkpPMFZzQ1pCWnRIajRQTzBBUXppZGt4REtyc2lkY2RKRUMwUlNQT3VwaVFHWU1VTUZVZS9nOFJYMEJPZVd0S2JEcG41T2ZJdXRlV09tTE0rKzRiazlpM2Y0akhIYS9hSjlCZ3EzUnM5OU1lMUl3PT1fMTY0Njg0NTc4MDk1NiIsImV4cCI6MTY0Njg0OTM4MDk1NiwiaXNzdWVkX2F0IjoxNjQ2ODQ1NzgwOTU2LCJjc3JmX3Rva2VuIjoiIn19.dKJSZde6S2MijW3eR4sK9e1aBdi49_E3ijw9rlOtYt0"

csrf_token="OdcDK5r1iPy7sPf4kw8oSUqyQrRInfqoMYryHBV/XLIDhlnwiruJWfdfDCwdfUMUBffagTf4JmoxXqci5HN78iUzXPZY1naiHJLjIVR44qCiSXtk7yxMELDc3xEwaucePo87CQ=="

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/login_status/" 

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1&limit=20"

ab -C "access_token=$access_token" -H "csrf_token: $csrf_token" -n 100000 -c 1000 -l "http://localhost:8000/api/user/?page=1000&limit=20"

