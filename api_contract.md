# API CONTRACT

## LoginAsAdmin

```
endpoint: /api/adm/login
request
{
    username: string
    password: string
}

response
set-cookie: artemis.sid
{
    success: bool
    servError: []string
    data: LoginAsAdminResponse {
        errMessage: []string
    }
}
```

## RegisterAsAdmin

```
endpoint: /api/adm/register
request
{
    username: string
    password: string
    emai: string
}

response
{
    success: bool
    servError: []string
    data: RegisterAsAdminResponse {
        errMessage: []string
        Message: []string
    }
}
```

### Logout (common)

```
endpoint:
  - /api/adm/logout
  - /api/st/logout (to be implemented)
cookie: artemis.sid
request
{}

response
{
    success: bool
    servError: []string
    data: LogoutResponse {
        errMessage: []string
        Message: []string
    }
}
```
