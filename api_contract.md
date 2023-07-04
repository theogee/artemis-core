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
endpoint: /api/students
cookie: artemis.sid
request
{
    limit int64 DEFAULT 20
    page int64 DEFAULT 1
}

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

### GetSGUMajors

```
endpoint: /api/sgu_majors
cookie: artemis.sid
request
{}

response
{
    success: bool
    servError: []string
    data: GetSGUMajorsResponse {
        errMessage: []string
        message: []string
        majors: [
            SGUMajor {
                majorID int
                majorName string
                majorCode string
            },
            ...
        ]
    }
}
```

### GetExchangeYear

```
endpoint: /api/exchange_year
cookie: artemis.sid
request
{}

response
{
    success: bool
    servError: []string
    data: GetExchangeYearResponse {
        errMessage: []string
        message: []string
        exchangeYear: [
            2023,
            2024,
            ...
        ]
    }
}
```
