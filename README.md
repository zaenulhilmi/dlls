# Dating App

## Requirements
    - Go 

## Running the App
    
```bash
go run main.go
```
The app will run on port 8010

## Running the Tests

```bash
go test ./...
```
This will run all the tests in the app

### Testing Watcher (For Development Purposes)

if you want to run the tests and watch for changes in the files, 
you can use nodemon to run the tests and watch for changes in the files.
make sure you have nodemon installed globally

```bash
nodemon --exec "go test ./..." --ext "go"
```

### Dependency Graph

![Dependency Graph](./dependency-graph.png)

## API Endpoints

### Login
```bash
curl --location 'http://localhost:8010/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "email-0@example.com",
    "password": "password"
}'    
```

### Sign Up
```bash
curl --location 'http://localhost:8010/api/v1/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "admin@example.com",
    "password": "password"
}'
```

### Preview Target Profile
```bash
curl --location 'http://localhost:8010/api/v1/preview-next-profile' \
--header 'Authorization: Bearer <JWT_TOKEN>'
```

### Swipe Profile
#### Like
```bash
curl --location 'http://localhost:8010/api/v1/like' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <JWT_TOKEN>' \
--data '{
    "target_id": "user-84"
}'
```
#### Pass
```bash
curl --location 'http://localhost:8010/api/v1/pass' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <JWT_TOKEN>' \
--data '{
    "target_id": "user-84"
}'
```

### Purchase Premium
```bash
curl --location --request POST 'http://localhost:8010/api/v1/subscribe' \
--header 'Authorization: Bearer <JWT_TOKEN>' \
--data ''
```

## Functionalities

- Sign up & Login to the App 
    A simple sign up and login using email and password, it will return a JWT token that will be used for all other requests.

- User Able to only view, swipe left (pass) and swipe right (like) 10 other dating profiles in total (pass + like) in 1 day.
    This will be a actions table that will store 
    - user ID
    - target user ID
    - action (pass or like)
    - datetime
- Same profiles can’t appear twice in the same day.
    Need to check if the user has already swiped on the target user ID in the same day. 


- User Able to purchase premium packages that unlocks one premium feature of your choosing. A few examples:
    - No swipe quota for user
    - Verified label for user

    simplifying this to just a boolean value in the user table.
    The case for swiping quota will be handled in the service layer by checking if the user has a premium account or not.



