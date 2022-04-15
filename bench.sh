SQL_USER="7B5ZW@DI839.com"

echo "SQL_USER=$SQL_USER"


access_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NUb2tlbiI6eyJkYXRhIjp7ImVtYWlsIjoiYWEiLCJ1dWlkIjoiZTMwNjBkMzAtNGE2My00MWI5LWIxNmEtM2YxYmZhNDE0NzA1In0sInVuYW1lIjoiYWEiLCJ0b2tlbl9pZCI6ImUzMDYwZDMwLTRhNjMtNDFiOS1iMTZhLTNmMWJmYTQxNDcwNV9MblBoenFiNEtKM1VTNHQ4L0dKQ1dnVm1leXVvd2lLWld1Yi9YTDdZNU5HdHlINXNaRFBtQjdlR2ZKd25rUHc1NkJTKzlwaHc4SnBXcklEbklyL0dSL1U3clpNVVFjNUtaMjgzUWxjYnh1emZaM0QxNmNSb2crZ2Z6dXUwcFdVczVRTE1VQT09XzE2NTAwNDQwODg0MTciLCJleHAiOjE2NTAwNDc2ODg0MTcsImlzc3VlZF9hdCI6MTY1MDA0NDA4ODQxNywiY3NyZl90b2tlbiI6IjBwRVZGdXE1eDJ1TEZPQmloTTMvcVB1eTk0ejh4S2FHckZOSld5ZzR6QW9EMkExc1RUMXJFRTVqOXZ1aG56WW1JbXdkd0h5aFVBTytnRmhxb2NXSVoxL28rcTFCVkNwNWRzVGtrcjdEbFhEdkk2UVhmeEZkKzFhL2c4Nk5ENWlHUENBZHp3PT0ifX0.Fj_WOW6sh9_Ux108GJMLnT-B2wG-99ndnB5psMTLl_Q"

csrf_token="0pEVFuq5x2uLFOBihM3/qPuy94z8xKaGrFNJWyg4zAoD2A1sTT1rEE5j9vuhnzYmImwdwHyhUAO+gFhqocWIZ1/o+q1BVCp5dsTkkr7DlXDvI6QXfxFd+1a/g86ND5iGPCAdzw=="

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


# ab  -H "secret: 1234" -n 100000 -c 1000 -l "http://localhost:8000/api/del_user_cache/1"




