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
