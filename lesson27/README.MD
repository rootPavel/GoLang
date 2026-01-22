// Публичная страница. Все могут смотреть curl http://localhost:8080/public

// Секретна страница - только с токеном (валидным) curl http://localhost:8080/secret

// Аутентификация curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"Username": "Dima", "Password":"1234"}'

CMD - curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d "{\"Username\":\"Dima\",\"Password\":\"1234\"}"

// На секретную страницу с авторизацией (заменить токен своим): curl http://localhost:8080/secret -H "Authorization:eyJhbG..."