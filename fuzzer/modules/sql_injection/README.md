# SQL Injection
The goal of this module is to demonstrate how in a CI environment there is a way to automate detection of SQL injection vulnerabilities with minimal input from a developer. The goal is to reduce software developer cycles while at the same time ensuring API security.

In this module we demonstrate two different SQL injection vulnerability detection methods. These two `strategies` are the `union` strategy and the `create_table` strategy.


## Union Strategy
The union strategy works by injecting a SQL string with a UNION to access data that a user should not have access to in the database. This attack is largely based on lab6.

Imagine that the server side executes the following SQL statement to retrieve a user's information.
```postgresql
SELECT
    u.info
FROM users u
WHERE u.name = '{username}'
AND u.password = '{password}'
```

Imagine that the `username` and `password` are passed in as a raw string without any validation. An attacker can then pass in the following strings as parameters:
- username: pwned
- password: ' union select info where 1=1; --

This would render the SQL statement:
```postgresql
SELECT
    u.info
FROM users u
WHERE u.name = 'pwned'
AND u.password = '' union select info where 1=1; -- '
```

The UNION in this statement has a `where 1=1` which always evalutes to true, effectively dumping out all user `info` records. The `--` comments out the rest of the SQL statement so that the server side service doesn't throw a database exception due to malformed SQL.

We take this attack strategy and build code to automatically detect it in `strategies/lib/union_strategy.py`. We run into two issues that are solved in the `create_table` strategy.
1. We need some way to generate valid API calls. For this example we are hard coding a known vulnerable endpoint. Obviously this is not a realistic pattern since knowing a vulnerable API endpoint exists makes the need to detect it less valuable.
1. The strategy also requires a valid response parser be able to be defined. If you look at the function `api.lab_6.parse_response()` you can see the parser we defined for lab6. There are a couple disadvantages to this approach. One is that we need to define a parser for each API endpoint. This becomes hard to manage as API endpoints grows. The other disadvantage is that the union attack has to somehow inject some information into the response. This isn't always guaranteed.


We solve problem one, API call generation, as part of this project. We talk about the details more in other documentation so will not go into our approach here. The other issue around detection is solved in the `sql_injection` method.

## SQL Injection Strategy
In the SQL injection strategy we hook the module into the top level runner that automatically generates API calls based on the OpenAPI specification. This solves the first problem of generating valid API calls. We make a custom `request_generator` that always passes in a SQL injection attack with a `create table uuid` if the parameter is a string. The generator can be found at `strategies.create_table_strategy.CreateTableGenerator`. We then make API calls with this string. Since we're in a CI environment, we have elevated permissions to the backend database. We can then look at the database to see if the uuid table was created. If the uuid table was created, we know the attack was successful. We can then log that successful attack and report the vulnerability.

With the elevated permissions on the database, we're able to come up with a way to detect SQL injection in a more automated fashion. Currently the module is built for OpenAPI APIs and postgreSQL databases. Some possible extensions are to create API call generators for other specs like graphql. You can also extend the database code to work with other SQL syntaxes.  