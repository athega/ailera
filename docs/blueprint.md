# FlockFlow

Backend for the FlockFlow iOS application

# Group Client

Endpoints used by the iOS application

## Login [/login]

+ Parameters
  + to (required, string, `test@example.com`)
  The email address to send the login link to

### POST

Login to FlockFlow via link in email

```
curl -X POST -d 'to=test@example.com' https://flockflow.herokuapp.com/login
```

+ Response 202

## Login [/login{?key}]

+ Parameters
  + key (required, string, `1234`)
  The key used to log in

### GET

Redirects to the FlockFlow application

The link is formatted like this:
```
flockflow://Login?token=<JWT>
```

+ Response 302 (application/json; charset=utf-8)
  Redirected to the application with a JWT
+ Response 401
+ Response 500

## Profile [/profile]

### POST

Update profile

```
curl -H 'Authorization: Bearer <JWT>' -X POST -d 'link=http://example.com&name=XYZ&phone=123' https://flockflow.herokuapp.com/profile
```

+ Response 200 (application/json; charset=utf-8)
+ Response 400
+ Response 401
+ Response 500

## Profile [/profile{?access_token}]

### GET

> Requires a valid token provided in the `Authorization` header or `access_token` query string parameter.

+ Response 200 (application/json; charset=utf-8)

# Group Info

Endpoints containing information about the FlockFlow backend

## Status [/__status]

### GET

+ Response 200 (application/json; charset=utf-8)

