## SQL 
There are couple of approaches
- using the native driver
- using the api

## Setting up the database

https://towardsdatascience.com/local-development-set-up-of-postgresql-with-docker-c022632f13ea

```
docker run -d --name dev-postgres -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=user -p 5432:5432 postgres
```
```
docker inspect dev-postgres -f "{{json .NetworkSettings.Networks }}"
```
```
docker run -p 80:80 -e 'PGADMIN_DEFAULT_EMAIL=user@domain.local' -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' \--name dev-pgadmin  -d dpage/pgadmin4
```
### Connecting to the database

We need 4 things
- host
- port
- username
- password

```
sql.Open("Drivername","user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
```
This returns the two value
- database connection
- err

In long running activity we can use the db.Ping() to check the connection 

GO has two subsittion 
- Where col1= - used in the place of 
- VALUES($1,$2) in the case of inserts or updates.

## Create Syntax
```
CREATE TABLE table_name (

    column1 datatype constrain,

    column2 datatype constrain,

    column3 datatype constrain,

    ....

);
```
For creating and inserting two variable has to be created

## Query
There are two functions
- query 
- query row
## Truncate and delete 

```
EmptyTable, _ := db.Exec("TRUNCATE TABLE test")
DropTable, _ := db.Exec("DROP TABLE test")


```
## Update
***
eg-1 Creating table and inserting value
13.01 - creating connection
eg-2 query example
eg-3 prepare + RowQuery
eg-4 update
eg-5 delete and truncate
***