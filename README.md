# auth-microservice

curl -X POST https://auth-microservice-vvr6.onrender.com/auth/signup \
-H "Content-Type: application/json" \
-d '{"email": "test@example.com", "password": "securepassword", "is_admin": false}'

curl -X POST https://auth-microservice-vvr6.onrender.com/auth/signin \
-H "Content-Type: application/json" \
-d '{"email": "test@example.com", "password": "securepassword", "is_admin": false}'

curl -X GET https://auth-microservice-vvr6.onrender.com/auth/get-email-from-token \
-H "Content-Type: application/json" \
-d '{"token": "your_token_here"}'

curl -X POST http://0.0.0.0:8080/auth/block-user -H "Content-Type: application/json" -d '{"email": "martinscazzola01@gmail.com", "reason": "Lenguaje Inapropiado", "days": 3}'

curl -X POST https://auth-microservice-vvr6.onrender.com/auth/unblock-user -H "Content-Type: application/json" -d '{"email": "brandon@example.com"}'

curl -X POST http://0.0.0.0:8080/auth/request-password-reset \
-H "Content-Type: application/json" \
-d '{"email": "user@example.com"}'

curl -X POST http://0.0.0.0:8080/auth/password-reset \
-H "Content-Type: application/json" \
-d '{
    "email": "tu_email@example.com",
    "password": "tu_nueva_contraseña",
    "token": "tu_token_aqui"
}'

curl -X POST http://0.0.0.0:8080/auth/generate-pin -H "Content-Type: application/json" -d '{
    "email": "martinscazzola01@gmail.com"
}'

curl -X POST http://0.0.0.0:8080/auth/verify-pin -H "Content-Type: application/json" -d '{
    "email": "martinscazzola01@gmail.com", "pin":"THLcOf"
}'

curl -X POST http://0.0.0.0:8080/auth/signin-with-google -H "Content-Type: application/json" -d '{
    "email": "martinscazzola01@gmail.com"
}'

curl -X POST http://0.0.0.0:8080/auth/signup \
-H "Content-Type: application/json" \
-d '{"email": "martinscazzola01@gmail.com", "password": "Martin0!", "is_admin": false}'


eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1hcnRpbnNjYXp6b2xhMDFAZ21haWwuY29tIiwiZXhwIjoxNzMyMjMyMzk1LCJpYXQiOjE3MzIxNDU5OTV9.0543M_GF9OC-dsCjgmTeDMnNNdNs6_YzEqXzQFSybsQ