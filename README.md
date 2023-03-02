# How to run the solution

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


### login credentials to generate jwt token
url - http://localhost:8080/api/v1/users/auth
method - POST
request-object
{
"email":"u1@xm.com",
"password":"password"
}

### Add a company
url - http://localhost:8080/api/v1/companies
method - POST
request-object
{
    "name": "abc ltd",
    "description": "IT company",
    "amountOfEmployees": 6,
    "registered": true,
    "type": "corporate"
}
### Get a company
url - http://localhost:8080/api/v1/companies/{companyId}
method - GET

### Delete a company
url - http://localhost:8080/api/v1/companies/{companyId}
method - DELETE

### Patch a company
url - http://localhost:8080/api/v1/companies
method - PATCH
request-object
{
    "name": "abc ltd",
    "description": "IT company",
    "amountOfEmployees": 6,
    "registered": true,
    "type": "corporate"
}

# Note
## unfortunately didnt have enough time to complete the task, nevertheless it was a good exercise :)

