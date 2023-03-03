# Documentation

## To bring the services up use (on the console)
task up

## To bring the services down use 
task down

## To build an image 
task build

## To run Tests
task test

## testing
task test

## For test coverage
task cover

## For benchmarking
task bench

## To linter
task lint


# How to run the solution

## Health check
url - http://localhost:8080/api/v1/healthcheck

## login credentials to generate jwt token
url - http://localhost:8080/api/v1/users/auth <br>
method - POST <br>
request-object
```
{
"email":"u1@xm.com",
"password":"password"
}
```

output - jwt token

## Add a company
url - http://localhost:8080/api/v1/companies <br>
method - POST <br>
request-object
```
{
    "name": "abc ltd",
    "description": "IT company",
    "amountOfEmployees": 6,
    "registered": true,
    "type": "corporate"
}
```
output - http status created with link to resouce in header location object

## Get a company
url - http://localhost:8080/api/v1/companies/{companyId} <br>
method - GET

output - company object

## Delete a company
url - http://localhost:8080/api/v1/companies/{companyId} <br>
method - DELETE

output - http status accepted with link to resouce in header location object

## Patch a company
url - http://localhost:8080/api/v1/companies <br>
method - PATCH <br>
request-object
```
{
    "name": "abc ltd",
    "description": "IT company",
    "amountOfEmployees": 6,
    "registered": true,
    "type": "corporate"
}
```

output - http status accepted with link to resouce in header location object

# Note
## unfortunately didnt have enough time to complete the task, nevertheless it was a good exercise :)

