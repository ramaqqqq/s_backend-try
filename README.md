
# Test : Developed by Lutfi M

- Golang (gorilla mux)
- MySQL

## Application Manual

1. Please make DB with name "synapsisdb".
2. Please adjust the .env file for DB_HOST(localhost) and DB_PASSWORD according to your preferences.
2. How to Run :

   ```
   go run . / go run main.go
   ```

## Or by Dockerfile and Docker Compose

- sudo docker build -t backend-go-synapsis:v1 .
- or sudo docker pull lutfy/backend-synapsis:v1 </br>
<https://hub.docker.com/r/lutfy/backend-synapsis>
- docker-compose up -d
- optional : sudo docker-compose restart backend-go

## Support Document 
- Can be seen in the docs folder (erd file and create database)


## Endpoints List

- USER

[POST]       `localhost:7000/api/login` : Login user </br>
[POST]       `localhost:7000/api/register` : Register </br>
</br>

- ITEM

[POST]      `localhost:7000/api/item` : Add item</br>
[GET]       `localhost:7000/api/item` : View All item</br>
[PUT]       `localhost:7000/api/item/{item_id}/edit` : Edit Item by item_id</br>
[DELETE]    `localhost:7000/api/item/{item_id}/delete` : Delete Item by item_id</br>
</br>

- SHOPPING CART

[POST]      `localhost:7000/api/shopping-cart/{item_id}/{quantity}` : Add Shopping-cart</br>
[GET]       `localhost:7000/api/shopping-cart` : View All Shopping-cart</br>
[PUT]       `localhost:7000/api/shopping-cart/{purchase_id}/{item_id}/{quantity}/edit` : Edit shopping-cart by purchase_id, item_id, quantity</br>
[DELETE]    `localhost:7000/api/shopping-cart/{purchase_id}/delete` : Delete shopping-cart by purchase_id</br>
</br>

- PAYMENT

[POST]      `localhost:7000/api/payment/{purchase_id}` : Add Payment</br>
[GET]       `localhost:7000/api/payment` : View All Payment</br>
[DELETE]    `localhost:7000/api/payment/{payment_id}/delete` : Delete Payment by payment_id</br>
</br>

- TOP-UP

[POST]    `localhost:7000/api/topup/{amount}` : Transfer</br>
</br>

## Endpoint Details

***

### [POST]/api/login

- Summary  
Login

- Description  
Login

#### RequestBody

- application/json

```ts
{
    "email": "udin@gmail.com",
    "password":"123"
}
```

#### Responses

- 200 Login Successfully

`application/json`

```ts
{
    "body": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InVkaW5AZ21haWwuY29tIiwiZXhwIjoxNjk2MDg5NDkyLCJyb2xlIjoiIiwidXNlcl9pZCI6IjEiLCJ1c2VybmFtZSI6InVkaW4ifQ.ibo3vCWD9e2B-7mT0eM1u6LDtq9I8YJZsedGImasB2w",
        "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InVkaW5AZ21haWwuY29tIiwiZXhwIjoxNjk2MDg5NDkyLCJyb2xlIjoiIiwidXNlcl9pZCI6IjEiLCJ1c2VybmFtZSI6InVkaW4ifQ.ibo3vCWD9e2B-7mT0eM1u6LDtq9I8YJZsedGImasB2w",
        "users": {
            "email": "udin@gmail.com",
            "role": "customer",
            "user_id": 1,
            "username": "udin"
        }
    },
    "code": 200,
    "message": "Successfully"
}
```

***

### [POST]/api/register

- Summary  
Customer Register

- Description  
Customer Register

#### RequestBody

- application/json

```ts
{
    "username": "udin",
    "email": "udin@gmail.com",
    "password":"123"
}
```

#### Responses

- 201 Registration Successfully

`application/json`

```ts
{
    "body": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InVkaW5AZ21haWwuY29tIiwiZXhwIjoxNjk2MDg5NDgwLCJyb2xlIjoiIiwidXNlcl9pZCI6IjEiLCJ1c2VybmFtZSI6InVkaW4ifQ.dpXpXOmIxHDTsHq77iI_Dt7MtzzVgLagQULdBCnxwbg",
        "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6InVkaW5AZ21haWwuY29tIiwiZXhwIjoxNjk2MDg5NDgwLCJyb2xlIjoiIiwidXNlcl9pZCI6IjEiLCJ1c2VybmFtZSI6InVkaW4ifQ.dpXpXOmIxHDTsHq77iI_Dt7MtzzVgLagQULdBCnxwbg",
        "users": {
            "created": "2023-09-30T20:58:00+07:00",
            "email": "udin@gmail.com",
            "id": 1,
            "role": "customer",
            "username": "udin"
        }
    },
    "code": 200,
    "message": "Successfully"
}
```

***

### [POST]/api/item

- Summary  
Created Category Item

- Description  
Created Category Item

#### Headers

```ts
Authorization: string
```

#### RequestBody

- application/json

```ts
{
    "item_name" : "Sepatu Nike",
    "price": 2000,
    "Stock": 40
}
```

#### Responses

- 201 Success created Item

`application/json`

```ts
{
    "body": {
        "item_id": 1,
        "item_name": "Sepatu Nike",
        "item_price": 2000,
        "item_stock": 40
    },
    "code": 201,
    "message": "Successfully"
}
```

***

### [GET]/api/item

- Summary  
View All Item

- Description  
View All Item

#### Headers

```ts
Authorization: string
```

#### Responses

- 200 Success view all Item

`application/json`

```ts
{
    "body": [
        {
            "item_id": 1,
            "user_id": 1,
            "item_name": "Sepatu Nike",
            "price": 2000,
            "stock": 40
        }
    ],
    "code": 200,
    "message": "Successfully"
}
```

***

### [PUT]/api/item/{item_id}/edit

- Summary  
Update Item By item_id

- Description  
Update Item By item_id

#### Headers

```ts
Authorization: string
```

#### RequestBody

- application/json

```ts
{
    "item_name" : "Adidas",
    "price": 2000,
    "Stock": 40
}
```

#### Responses

- 200 Success update Item by item_id

`application/json`

```ts
{
    "body": {
        "item_id": 1,
        "user_id": 1,
        "item_name": "Adidas",
        "price": 2000,
        "stock": 40
    },
    "code": 200,
    "message": "Successfully"
}
```

***

### [DELETE]/api/item/{item_id}/delete

- Summary  
Delete Item

- Description  
Delete Item

#### Headers

```ts
Authorization: string
```

#### Responses

- 200 Success delete Item

`application/json`

```ts
{
    "code": 200,
    "deleted": "success to deleted",
    "message": "Succesfully"
}
```

***

### [POST]/api/shooping-cart/{item_id}/{quantity}

- Summary  
Created Shooping-cart

- Description  
Created Shooping-cart

#### Headers

```ts
Authorization: string
```

#### Responses

- 201 Success created Shooping-cart

`application/json`

```ts
{
    "body": {
        "item_id": 1,
        "purchase_date": "2023-09-30T21:13:04.634305149+07:00",
        "purhase_id": 1,
        "quantity": 10,
        "total_price": 20000,
        "user_id": 1
    },
    "code": 201,
    "message": "Successfully"
}
```

***

### [GET]/api/shopping-cart

- Summary  
View All Shooping-cart

- Description  
View All Shooping-cart

#### Headers

```ts
Authorization: string
```

#### Responses

- 200 Success view all Shooping-cart

`application/json`

```ts
{
    "body": [
        {
            "purchase_id": 1,
            "user_id": 1,
            "item_id": 1,
            "quantity": 10,
            "total_price": 20000,
            "purchase_date": "2023-09-30T21:13:04+07:00",
            "Payment": {
                "payment_id": 0,
                "user_id": 0,
                "amount": 0,
                "status": "",
                "order_id": "",
                "snap_url": "",
                "payment_date": "0001-01-01T00:00:00Z",
                "purchase_id": 0
            }
        }
    ],
    "code": 200,
    "message": "Successfully"
}
```

***

### [PUT]/api/shopping-cart/{purchase_id}/{item_id}/{quantity}/edit

- Summary  
Update Shooping-cart By purchase_id, item_id, quantity

- Description  
Update Shooping-cart By purchase_id, item_id, quantity

#### Headers

```ts
Authorization: string
```

#### RequestBody

- application/json

```ts
{
    "body": {
        "item_id": 1,
        "purchase_id": 1,
        "quantity": 20,
        "total_price": 40000,
        "user_id": 1
    },
    "code": 200,
    "message": "Successfully"
}
```

#### Responses

- 200 Success updated Shooping-cart By purchase_id, item_id, quantity

`application/json`

```ts
{
    "body": {
        "item_id": 1,
        "user_id": 1,
        "item_name": "Adidas",
        "price": 2000,
        "stock": 40
    },
    "code": 200,
    "message": "Successfully"
}
```

***

### [DELETE]/api/shopping-cart/{purchase_id}/delete

- Summary  
Deleted Shooping-cart

- Description  
Deleted Shooping-cart

#### Headers

```ts
Authorization: string
```

#### Responses

- 200 Success deleted Shooping-cart

`application/json`

```ts
{
    "body": "success to deleted purchase",
    "code": 200,
    "message": "Deleted"
}
```

***

### [POST]/api/payment/{purchase_id}

- Summary  
Saved payment

- Description  
Saved payment

#### Headers

```ts
Authorization: string
```

#### Responses

- 201 Success Save payment

`application/json`

```ts
{
    "body": {
        "amount": 40000,
        "order_id": "28156188-d86a-4ab8-99ce-baac918e8c15",
        "payment_date": "2023-09-30T21:31:37.217301048+07:00",
        "payment_id": 1,
        "purchase_id": 1,
        "status": "Pembayaran berhasil",
        "user_id": 1
    },
    "code": 201,
    "message": "Successfully"
}
```

***

### [GET]/api/payment

- Summary  
View All payment

- Description  
View All payment

#### Headers

```ts
Authorization: string
```

#### Responses

- 200 Success view all payment

`application/json`

```ts
{
    "body": [
        {
            "payment_id": 1,
            "user_id": 1,
            "amount": 40000,
            "status": "Pembayaran berhasil",
            "order_id": "28156188-d86a-4ab8-99ce-baac918e8c15",
            "snap_url": "by saldo user",
            "payment_date": "2023-09-30T21:31:37+07:00",
            "purchase_id": 1
        }
    ],
    "code": 200,
    "message": "Successfully"
}
```

***

### [DELETE]/api/payment/{payment_id}/delete

- Summary  
Deleted Payment

- Description  
Deleted Payment

#### Headers

```ts
Authorization: string
```

#### Responses

- 200 Success deleted Payment

`application/json`

```ts
{
    "body": "success to deleted payment",
    "code": 200,
    "message": "Deleted"
}
```

***

### [POST]/api/topup/{amout}

- Summary  
Add saldo

- Description  
Add saldo

#### Headers

```ts
Authorization: string
```

#### Responses

- 201 Success Add saldo

`application/json`

```ts
{
    "body": {
        "amount": 20000000,
        "id": "1"
    },
    "code": 201,
    "message": "Successfully"
}
```

