# FlockFlow

Backend for the FlockFlow iOS application

# Group Client

Endpoints used by the iOS application

## Login [/login{?key}]

Redirects to the FlockFlow application

+ Parameters
  + key (required, string, `1234`)
  The key used to log in

### GET

+ Response 302 (application/json; charset=utf-8)
  Redirected to the application with a JWT
+ Response 401
+ Response 500

## Profile [/profile]

### GET

> Requires a valid token provided in the `Authorization` header or `access_token` query string parameter.

+ Response 200 (application/json; charset=utf-8)

# Group Info

Endpoints containing information about the FlockFlow backend

## Status [/__status]

### GET

+ Response 200 (application/json; charset=utf-8)

