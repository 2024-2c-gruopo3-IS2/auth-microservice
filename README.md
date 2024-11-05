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

curl -X POST https://auth-microservice-vvr6.onrender.com/auth/block-user -H "Content-Type: application/json" -d '{"email": "brandon@example.com"}'

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


