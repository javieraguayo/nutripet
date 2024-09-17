curls app mascotas

1. Registrar un usuario:

curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{
       "name": "Javi",
       "email": "javier.aguayo19@gmail.com",
       "password": "1234"
   }'

2. Iniciar sesión:


curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{
       "username": "javier.aguayo19@gmail.com",
       "password": "1234"
   }'

3. Crear una mascota:

 curl -X POST http://localhost:8080/pets -H "Content-Type: application/json" -d '{
       "name": "Kim",
       "userID": 1
   }'

4. Listar mascotas:ç

curl -X GET http://localhost:8080/pets

5. Crear una comida:

   curl -X POST http://localhost:8080/meals -H "Content-Type: application/json" -d '{
       "name": "catshow Food ",
       "petID": 1,
       "quantity": 20,
       "note": "Morning meal",
       "time": "2023-10-01T08:00:00Z"
   }'

6.Listar Comidas :

   curl -X GET http://localhost:8080/meals

7.Obtener la próxima notificación de comida para una mascota específica:

   curl -X GET http://localhost:8080/next-meal/1

8.Crear un plan:

   curl -X POST http://localhost:8080/plans -H "Content-Type: application/json" -d '{
       "name": "Weight Loss Plan"
   }'