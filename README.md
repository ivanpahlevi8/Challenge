# Challenge

## Program Introduction
This repository contains application program called Online Store Application. Online Store Application was build using Go Language. Online
Store Application also implements some library that can be found at [github website](github.com). Some of those library can be seen at 
list below : <br/>
1. [Javascript Web Token](github.com/dgrijalva/jwt-go)
2. [Chi routing](github.com/go-chi/chi)
3. [pq postgres sql](github.com/lib/pq)
4. [scs session manager](github.com/alexedwards/scs/v2)
5. [uui token generator](github.com/google/uuid)
6. [crypto](golang.org/x/crypto)
7. [sys](golang.org/x/sys)

Based on library that imported for this project, the project itself have some features based on library that was implemented during project
development. Some of features that supported Online Store Application, especially for backend support, can be found at featured list below, <br/>

Features List : <br/>
1. Customer can view product list by product category
2. Customer can add product to shopping cart
3. Customers can see a list of products that have been added to the shopping cart
4. Customer can delete product list in shopping cart
5. Customers can checkout and make payment transactions
6. Login and register customers

## Program Explanation
Based on Application Features, there are several application program that can be use to support user experience based on backend development.
Some of those features is develop based on features list that already mentioned before. <br/>
Application use persistance database to save data. Database that used in this Application is Relational Database. This relational database
was implement to Application using postgres sql that can be found in [here](postgresql.org). Data was saved in three diffrent table in same database. The first table is used to save user data. The second table is used to save item or product data. The last table is used to save basket list of item from user. To connect with database sql, in this project, is using credential as below : <br/>
1. host, using local host with value "127.0.0.0"
2. port, using port number "5432"
3. user, using user "postgres"
4. password, password depends on database
5. dbname, "UserSynapsisDatabase"
<br/> 
Table in database can be seen below : </br>
- User table <br/>
CREATE TABLE usertable (
	id varchar(255),
	username varchar(255),
	password varchar(255),
	first_name varchar(255),
	last_name varchar(255),
	age int,
	list_id varchar(255)
); <br/>
[user table](https://github.com/ivanpahlevi8/Challenge/assets/83549388/457afb11-2db8-43d3-9d11-dac81a2968ca) <br/><br/>
- Item Table <br/>
CREATE TABLE itemtable (
	id varchar(255),
	item_name varchar(255),
	item_category varchar(255),
	item_price float(8),
	item_quantity INT
); <br/>
![item table](https://github.com/ivanpahlevi8/Challenge/assets/83549388/fcf44c14-814a-4271-a262-938bbf783af7) <br/><br/>
- Basket Table <br/>
CREATE TABLE shoptable (
	id varchar(255),
	all_items text[]
);<br/>
![shop table](https://github.com/ivanpahlevi8/Challenge/assets/83549388/e48cf46b-325a-4a48-949c-53d5f3a0c7e3) <br/><br/>

The feature that involve with persistance database can be seen in the next section. </br>

#### User Register Features
Register features in this application, can be used for customer ot user to create their account in database. This features is using Post request in API. In application, url for this feature is <br/>
``` sql
http://localhost:2020/add-user
```
API documentation can be seen in this picture, <br/>
![postman_register](https://github.com/ivanpahlevi8/Challenge/assets/83549388/ea439e97-894f-4ce2-9766-c6b6cc7dcfb2)
<br/>
And, when we check the database, we can see that data in table are added with new data, <br/>
![postgres user table register](https://github.com/ivanpahlevi8/Challenge/assets/83549388/3eb67421-0450-4521-8132-83d82dc13a8e)
<br/>
In picture above, we can see that the password already casting into other form. This is one of features in application to make user
user credential are safe and cannot be seen by other user or administrator. Password was casted using JWT library or Javascript Web Token.

#### User Login Features
Login features in this application, can be user for customer to login with their account which their account already registered before.
Login features is used to access some restricted url for user to manipulate data in their basket. Login features use Post request in API.
This feature can be accessed through this url, <br/>
``` sql
http://localhost:2020/login
```
API documentation using postman can be seen in this picture, <br/>
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/53f619be-a895-4f2e-b8a3-15aaefd99767)
Based on those image, if the user already registered and password is correct then user will get access. If username wrong, it will show
message like picture below :
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/c8576997-9b63-44f2-9225-fd0e12ff1be2)
<br/>
If password is wrong, it will show message like this picture :
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/2003e11c-2a2d-4d41-ba2c-e9eb48640eac)
<br/>
As we can see, the password was casting in database, so in order to compare user input and password in database, it must use JWT to compare
user input and password as token in database.

### 
