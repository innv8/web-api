# WEB API - LESSON

[ Link to repo](https://github.com/innv8/web-api)

[Explanation to MVC](https://github.com/innv8/api-introduction#lesson-2--mvc)

---

This API will help us manage a list of users in a database.

### 1. Creating a database and a users table

user table

- id pk, int, auto increment
- first_name string
- last_name string
- phone_number string unique
- date_of_birth date
- created datetime (default current_timestamp)
- updated datetime (default null but on update, current_timestamp)

### 2. Creating an API that interacts with that table.

Endpoints

DONE :  1. /            [*]         - print "welcome home" 

DONE :  2. /user        [POST]      - create a new user

DONE :  3. /users       [GET]       - list all users/ search a user

DONE :  4. /users/3     [GET]       - return one user or 404 if the user is not in db

DONE :  5. /users/3     [PUT]       - update the details of user

DONE :  6. /users/3     [DELETE]    - Delete the user



---
## Resources

[Time](https://www.digitalocean.com/community/tutorials/how-to-use-dates-and-times-in-go)

[Regex validator](https://regex101.com/)

---
## Topics in the Advanced Web API

- Regexes in urls

- Validating payload from client

- Caching

- OAuth2

- Deploying on Sever