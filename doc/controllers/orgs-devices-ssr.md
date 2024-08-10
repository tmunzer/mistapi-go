# Orgs Devices-SSR

```go
orgsDevicesSSR := client.OrgsDevicesSSR()
```

## Class Name

`OrgsDevicesSSR`


# Get Org 128 T Registration Commands

128T devices can be managed/adopted by Mist.

```go
GetOrg128TRegistrationCommands(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ReponseRouter128tRegisterCmd],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ReponseRouter128tRegisterCmd`](../../doc/models/reponse-router-128-t-register-cmd.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsDevicesSSR.GetOrg128TRegistrationCommands(ctx, orgId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "conductor_cmd": "register mist eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ",
  "registration_code": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ",
  "router_shell_cmd": "128agent register --registration-code eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJvcmdfaWQiOiIyODE4ZTM4Ni04ZGVjLTI1NjItOWVkZS01YjhhMGZiYmRjNzEiLCJzdmMiOiIxMjhyb3V0ZXIiLCJwcm92aWRlciI6ImF3cyIsImVudiI6ImxvY2FsIiwiZXB0ZXJtX3VybCI6IndzczovL2xvY2FsL3dzIiwiaWF0IjoxNjEzODQ3NDg0LCJleHAiOjE2NDUzODM0ODR9.YnhgThKYAj1uaooi6j-zY8dMipp5YqJxnn79B9TB5XQ"
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

