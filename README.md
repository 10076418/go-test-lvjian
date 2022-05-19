## Mysql create Database
```
CREATE DATABASE `gotestdb`
```
## Mysql create Table
Please refer to 【gotestdb.sql】 file under the 【extra】 folder

## Postman job
Please refer to 【GOTest.postman_collection.json】 file under the 【extra】 folder


## 1. Signup API

**Url** : `http://127.0.0.1:8080/v1/signup`

**Method** : `POST`

**Params** : `FirstName`, `LastName`, `Email`, `Password`

**Auth Required** : `No`

**Request** :

```
{
    "FirstName" : "Baggio",
    "LastName":"Roberrt",
    "Email":"10076418@qq.com",
    "Password":"123456789test"
}
```
**Response** :

```
{
    "code": 200,
    "msg": "success",
    "data": {
        "FirstName": "Baggio",
        "LastName": "Roberrt",
        "Email": "10076418@qq.com"
    }
}
```

## 2. Signin API

**Url** : `http://127.0.0.1:8080/v1/signin`

**method** : `POST`

**Params** : `Email`, `Password`

**Auth Required** : `No`

**request** :

```
{
    "Email":"10076418@qq.com",
    "Password":"123456789test"
}
```
**response** :

```
{
    "code": 200,
    "msg": "success",
    "data": {
        "Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjEwMDc2NDE4QHFxLmNvbSIsIlBhc3N3b3JkIjoiMTIzNDU2Nzg5dGVzdCIsImV4cCI6MTY1Mjg2NTEwMywiaWF0IjoxNjUyODY0NTAzLCJuYmYiOjE2NTI4NjQ1MDN9.tsoRAiyFmZlqNWasIxadc2FywDoVcdwik8m8EMl5pLc"
    }
}
```

## 3. Profile API

**Url** : `http://127.0.0.1:8080/v1/profile`

**Method** : `GET`

**Params** : `NO`

**Auth Required** : `YES`
```
KEY:  Authorization  
VALUE: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjEwMDc2NDE4QHFxLmNvbSIsIlBhc3N3b3JkIjoiMTIzNDU2Nzg5dGVzdCIsImV4cCI6MTY1Mjg2NDk1MywiaWF0IjoxNjUyODY0MzUzLCJuYmYiOjE2NTI4NjQzNTN9.gvrkcbqg5R78GAR6xT5D7OIpLDxjvnD_oMGh83EGwsc
```

**Request** : 

**Response**  :

```
{
    "code": 200,
    "msg": "success",
    "data": {
        "FirstName": "lv",
        "LastName": "jian",
        "Email": "10076418@qq.com"
    }
}
```

## 4. Update API

**Url** : `http://127.0.0.1:8080/v1/profile/update`

**Method** : `POST`

**Params** : `FirstName`, `LastName`

**Auth Required** : `YES`
```
KEY:  Authorization  
VALUE: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IjEwMDc2NDE4QHFxLmNvbSIsIlBhc3N3b3JkIjoiMTIzNDU2Nzg5dGVzdCIsImV4cCI6MTY1Mjg2NDk1MywiaWF0IjoxNjUyODY0MzUzLCJuYmYiOjE2NTI4NjQzNTN9.gvrkcbqg5R78GAR6xT5D7OIpLDxjvnD_oMGh83EGwsc
```

**Request** :

```
{
    "FirstName":"lv",
    "LastName":"jian"
}
```

**Response**  :
```
{
    "code": 200,
    "msg": "success",
    "data": {
        "FirstName": "lv",
        "LastName": "jian",
        "Email": "10076418@qq.com"
    }
}
```