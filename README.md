# ECHO-CONTACTS

Echo-Contacts is repository that demonstrates the use of `echo` framework in creating contacts.

The repository uses `mysql` as db driver and the below given endpoints can be used to `Create` , `Read` , `Update` and `Delete`.

The endpoints uses jwt token based authorization and the token can be generated using the `/login` endpoint.

## SETUP

Before running the setup, ensure that you have dep as the dependency manager and run `dep ensure` to fetch the dependencies.

- Open the `/setup/setup.sh` and replace the `mysql` db username and password as per your system.

- Run the make command to execute the program. The make command will do the following.

- Setup your System.
  - Setup environment variables and create the database.
  - Build the Repository
  - Run the build file

## Endpoints

```json
[
  {
    "method": "POST",
    "path": "/login",
    "name": "github.com/binkkatal/echo-contacts/api.Login"
  },
  {
    "method": "GET",
    "path": "/contacts/index",
    "name": "github.com/binkkatal/echo-contacts/api.(*Api).Index-fm"
  },
  {
    "method": "DELETE",
    "path": "/contacts/:id",
    "name": "github.com/binkkatal/echo-contacts/api.(*Api).Delete-fm"
  },
  {
    "method": "PATCH",
    "path": "/contacts/:id/update",
    "name": "github.com/binkkatal/echo-contacts/api.(*Api).Update-fm"
  },
  {
    "method": "POST",
    "path": "/contacts/create",
    "name": "github.com/binkkatal/echo-contacts/api.(*Api).Create-fm"
  },
  {
    "method": "GET",
    "path": "/contacts/:id",
    "name": "github.com/binkkatal/echo-contacts/api.(*Api).Show-fm"
  }
]
```

## Accessing the endpoints

The api uses jwt token based authentication ,
The default `username` and `password` has been set as

```bash
  username -> bink
  password -> password
```

Using the above credentials , access the `/login` endpoint and generate the jwt tokens.

Pass the bearer token in the header as `Authorization: Bearer <JWT_TOKEN>` .
After that the endpoints will be accessible.
