# Self Alarms

```go
selfAlarms := client.SelfAlarms()
```

## Class Name

`SelfAlarms`


# List Alarm Subscriptions

Get List of all the subscriptions

```go
ListAlarmSubscriptions(
    ctx context.Context) (
    models.ApiResponse[[]models.ResponseSelfSubscription],
    error)
```

## Response Type

[`[]models.ResponseSelfSubscription`](../../doc/models/response-self-subscription.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := selfAlarms.ListAlarmSubscriptions(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

