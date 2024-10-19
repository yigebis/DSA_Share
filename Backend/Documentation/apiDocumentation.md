# ShopOps API Documentation

## Overview
The `DSAShare` API allows shop owners to register, log in, verify emails. All requests and responses use JSON format, and the API requires authentication via JWT tokens for specific routes.

### Base URL
`http://localhost:8080/`

---

## Endpoints

### **User Registration**

#### `POST /register`
Registers a new user (shop owner).

- **Request Body:**
  ```json
  {
    "first_name": "string",
    "last_name": "string",
    "sex" : "string",
    "email": "string",
    "password": "string"
  }

- **Validation:**
  "first_name" : required, minimum_length = 1, maximum_length = 50
  "last_name" : required, minimum_length = 1, maximum_length = 50
  "sex" : required, can be of types "M" and "F"
  "email" : required, must be an email
  "password" : required, length >= 8, lower_case >= 2, upper_case >= 1, special >= 1


- **Responses:**
    -200 OK - Registration successful, verification email sent.
    -400 Bad Request - Invalid request payload.
    -409 Conflict - User already exists or pending verification.


### **User Login**

#### `POST /login/email`
Logins a user (shop owner/employee).

- **Request Body:**
  ```json
  {
    "identifier": "string",
    "password": "string"
  }

- **Validation:**
  "identifier" : required
  "password" : required


- **Responses:**
    -200 OK - Returns JWT access and refresher tokens.
    ```json
    {
        "token" : "string",
        "refresher" : "string"
    }
    -400 Bad Request - Invalid request payload  
                     - account not verified 
                     - invalid email or password


#### `POST /login/user_name`
Logins a user (shop owner/employee).

- **Request Body:**
  ```json
  {
    "identifier": "string",
    "password": "string"
  }

- **Validation:**
  "identifier" : required
  "password" : required


- **Responses:**
    -200 OK - Returns JWT access and refresher tokens.
    ```json
    {
        "token" : "string",
        "refresher" : "string"
    }
    -400 Bad Request - Invalid request payload  
                     - account not verified 
                     - invalid user_name or password  


### **Email verification**

#### `GET /verify`
verifies a registered user (shop owner).

- **Query Parameters:**
  -email
  -token

- **Responses:**
    -200 OK - Email verified successfully.
    -400 Bad Request - Invalid verification token.
    -409 Conflict - Email already verified.


