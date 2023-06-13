# API CONTRACT

## LoginAsAdmin

```
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
request
{
    username: string
    password: string
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
