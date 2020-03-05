# tv rest api
An api with one endpoint :D

## How to run?
Some requirements are needed to run this app:
1. [PostgreSQL 10](https://www.postgresql.org/)
2. Go version 1.13 or later

### Steps to run
1. Create SQL data base: 
```sql 
  CREATE DATABASE restapi;
``` 
2. Create Users table in the database:
```sql
CREATE TABLE Users (
  id_user SERIAL PRIMARY KEY,
  full_name VARCHAR(50) NOT NULL,
  identification VARCHAR(15) NOT NULL,
  birth_date DATE
);
```
3. Download or clone this repo and extract it
4. Go to extracted folder
5. Next we need to set postgres connection string to our .env file, for instance: `CONNECTION_STR=postgres://postgres:your_postgres_passwrod@localhost/your_database_name?sslmode=disable`
6. In the folder extracted previously, open a cmd or terminal and run the command: `go run main.go`, this will open an http server listening at localhost:8080
7. Now our server is running :D

## Api Docs
This api has only a one endpoint
### User

**Endpoint:** `/user`  
**Method:** `POST`  
**Auth required:** NO

**Data constraints:**  
```json
{
  "fullName": "[required, 10 to 50 chars]",
  "identification": "[required, identification in plain text, 10 to 15 digits]",
  "birthDate": "[birthDate in plain text, format dd-mm-yyyy]"
}
```

**Data example:**
```json
{
  "fullName": "Carl Karlson",
  "identification": "1234567890",
  "birthDate": "07-10-1989"
}
```  
With empty birth date  
```json
{
  "fullName": "Carl Karlson",
  "identification": "1234567890",
  "birthDate": ""
}
```

**Success Response**  
`200 OK`  
```json
{
  "msg": "resource created",
}
```

**Error Response**  
Condition: if some server error happend  
`500 INTERNAL SERVER ERROR`  
```json
{
  "errors": [
    "Something went wrong :("
  ],
}
```

Condition: if the data sent to server has incorrect shape  
`400 BAD REQUEST`  
```json
{
  "errors": [
    "The data could not be processed"
  ],
}
```

Condition: if 'fullName' and 'identification' are empty  
`400 BAD REQUEST`  
```json
{
  "errors": [
    "The field FullName is required",
    "The field Identification is required"
  ],
}
```

Condition: if 'fullName' is empty  
`400 BAD REQUEST`  
```json
{
  "errors": [
    "The field FullName is required"
  ],
}
```

Condition: if 'fullName' is less than 10 chars  
`400 BAD REQUEST`  
```json
{
  "errors": [
    "The field FullName is too short, needs at least 10 elements/characters"
  ],
}
```

Condition: if 'fullName' is greater than 50 chars  
`400 BAD REQUEST`  
```json
{
  "errors": [
    "The field FullName is too long, maximun 50 elements/characters"
  ],
}
```

Condition: if 'identification' is empty  
`400 BAD REQUEST`  
```json
{
  "errors": [
    "The field Identification is required"
  ],
}
```

Condition: if 'identification' is less than 10 digits  
`400 BAD REQUEST`  
```json
{
  "errors": [
    "The field Identification is too short, needs at least 10 elements/characters"
  ],
}
```

Condition: if 'identification' is greater than 15 digits  
`400 BAD REQUEST`  
```json
{
  "errors": [
    "The field Identification is too long, maximun 15 elements/characters"
  ],
}
```

Condition: if 'birthDate' has incorrect format  
`400 BAD REQUEST`  
```json
{
  "errors": [
    "The field birthDate has incorrect format, it should be dd-mm-yyyy"
  ],
}
```

## TODO
1. Simplify verbose error messages
2. Extract validation logic to a function to reuse
3. Use `NullTime` instead of `interface{}` in user_handler.go -> line 41 -> `var birthDate interface{}`
