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
<br/>
When creating user, Application automatically create basket object. It can be seen in user database, as list_id column. So when user
registered their account, it automatically create a basket object for them to add and remove item in their basket.

#### User Login Features
Login features in this application, can be used for customer to login with their account which their account already registered before.
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
<br/>

### User Add Product Features
This feature in this application can be used for user to add item in store to their account basket. To access this feature, user must
login first with their account in database. If user already login, user can use this feature. To access this feature, user can use
url below, <br/>
``` sql
http://localhost:2020/user-get-item?item_id=67df7eca-31cf-4be0-8899-366402650ceb
```
those url is used with url parameter, item_id. This item id is refer to item table in database. In url above, user want to add item with
id 67df7eca-31cf-4be0-8899-366402650ceb or _kertas A4_ to their basket. In API postman, if we got to url, it will give feedback like
this picture, <br/>
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/ff682880-2b38-4134-b844-895c9191964f)
<br/>
if user is not login yet, API will send message like this picture, <br/>
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/4db36f8b-0ba0-4904-8550-8aac4527138d)
<br/>
if url success, we can see in shop database, that item _kertas A4_ already added to user basket, <br/>
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/62332f5b-b575-48bb-b867-e7f037dce5d8)
</br>

#### User Delete Product Features
This feature is oposite with previous feature, where in this feature user can delete item in their account basket. To access this feature, same as previous feature, user must login first with their account in database. If user already login, user can accessed this feature thorugh this link, <br/>
``` sql
http://localhost:2020/user-delete-item?item_id=67df7eca-31cf-4be0-8899-366402650ceb
```
Same as previous feature, url is used with url parameter called item_id. This item id is refer to item table in database. In url above, 
user want to delete item with id 67df7eca-31cf-4be0-8899-366402650ceb or _kertas A4_ In their basket. if we test url in postman we will get feedback like this picture, <br/>
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/8f823f5d-a895-4a51-86a2-02b7c7d63514)
if user is not login yet, API will send message like this picture, <br/>
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/25e4a308-1182-41c6-b147-d3e67d3f9abc)
if url success, we can see in shop database, item _kertas A4_ already remove from user basket, <br/>
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/386ec17a-e638-4abe-9a37-94902b8e32a3)
<br/>

#### User See List Product In Basket
This feature is used for user to see all item in their basket. Same as previous feature, this feature also need for user to be logged in. User can access this feature, if already logged in, throug this link, <br/>
``` sql
http://localhost:2020/user-get-all-item
```
This url does not need url parameter because we want to see all items. For demonstration, we will login with user with username ivanindirsyah07 because this account alread has item in it basket. User, <br/>
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/cba5a240-1cac-439d-acb5-45016472a7af)
<br/>
If user already logged in, API will give feedback as picture below, 
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/c9b3b7c5-99c2-4f91-bcfe-d0622ea9cf8e)
<br/>
Same as data shop in database, <br/>
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/98e6ecdf-50ff-4fb3-b109-229baa9aaeac)
<br/>
If user was not logged in yet, API will give feedback as picture below, <br/>
![image](https://github.com/ivanpahlevi8/Challenge/assets/83549388/495ae7df-4ed9-485d-bc41-01d4d5ec9d39)
<br/>

#### User See List Product Based On Category In Store
