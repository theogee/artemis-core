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
        usernameError: string
        passwordError: string
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
        message: []string
    }
}
```

### Logout (common)

```
endpoint: /api/logout
cookie: artemis.sid
request
{}

response
{
    success: bool
    servError: []string
    data: LogoutResponse {
        errMessage: []string
        message: []string
    }
}
```

### GetMeta

```
endpoint: /api/meta
cookie: artemis.sid
request
{}

response
{
    success: bool
    servError: []string
    data: GetMetaResponse {
        userType string
    }
}
```

### GetStudents

```
** to be implemented
endpoint: /api/students
cookie: artemis.sid
request
{}

response
{
    success: bool
    servError: []string
    data: GetStudentsResponse {
        errMessage: []string
        message: []string
        students: [
            student {
                studentID uint32
                givenName string
                surname   string
                SGUMajor  string
                SGUEmail  string
                MobileDE  string
                MobileID  string
            },
            ...
        ]
    }
}
```
