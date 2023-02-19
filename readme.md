### Go backend
Just a simple go backend testing - utilises GORM, echo, viper and bcrypt.
Goal is to make a simple CRUD backend with a database for User model.

Future plans: To include logging.

Ideal plans: To include JWT authentication. Most likely on hold first, will move to other projects after adding logging.

### Changelog

19 Feb 2023 - Added logging with zerolog.

18 Feb 2023 - Attempting logging with zap.

16 Feb 2023 - Adjusted bcrypt for password hashing. Tested with postman.

Added login function and tested with postman.

15 Feb 2023 - Added bcrypt for password hashing.

13 Feb 2023 - Added delete user function. Tested with Postman.

Edited response for some functions. Done with CRUD for users.

12 Feb 2023 - Adjusted update function.

Confirmed that update user works, tested with Postman.

7 Feb 2023 - Adjusting update function.

6 Feb 2023 - Adding update function.

5 Feb 2023 - Adjusted database and user controller code.

Confirmed that create and read users work, tested with Postman.

Added get user by id, email and username, tested with Postman.

4 Feb 2023 - Added future plans statement.

2 Feb 2023 - Added changelog to keep track of changes

1 Feb 2023 - Edited database code

31 Jan 2023 - Added users controller and route

30 Jan 2023 - Added echo package for routing

29 Jan 2023 - First commit, simple backend that adds data into the database with GORM