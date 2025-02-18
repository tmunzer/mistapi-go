# Orgs NACIDP

```go
orgsNACIDP := client.OrgsNACIDP()
```

## Class Name

`OrgsNACIDP`


# Validate Org Idp Credential

IDP Credential Validation. The output will be available through websocket. As there can be multiple command issued against the same device at the same time and the output all goes through the same websocket stream, `session` is introduced for demux.

#### Subscribe to Device Command outputs

`WS /api-ws/v1/stream`

```json
{
    "subscribe": "orgs/:org_id/mist_nac/test_idp"
}

```

### Response (no idp can be found)

```json
{
    "event": "data",
    "channel": "/orgs/:org_id/mist_nac/test_idp",
    "status": 
    "data": {
        "status": "failure",
        "error": "No matching IDP found"
    }
}

```

### Response OK

```json
{
    "event": "data",
    "channel": "/orgs/:org_id/mist_nac/test_idp",
    "status": 
    "data": {
        "status": "success",
        "idp_id": "915793c0-1355-4e98-b1c0-23df2227b357",
        "idp_type": "ldap",
        // more attributes will be added later
    }
}

```

### Response Invalid Credentials

```json
{
    "event": "data",
    "channel": "/orgs/:org_id/mist_nac/test_idp",
    "status": 
    "data": {
        "status": "failure",
        "error": "Invalid Credentials",
        "idp_id": "915793c0-1355-4e98-b1c0-23df2227b357",
        "idp_type": "ldap",
    }
}

```

```go
ValidateOrgIdpCredential(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.UsernamePassword) (
    models.ApiResponse[models.WebsocketSession],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.UsernamePassword`](../../doc/models/username-password.md) | Body, Optional | - |

## Response Type

[`models.WebsocketSession`](../../doc/models/websocket-session.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.UsernamePassword{
    Password:             models.ToPointer("xxxxxx"),
    Username:             models.ToPointer("suriyas@juniper.net"),
}

apiResponse, err := orgsNACIDP.ValidateOrgIdpCredential(ctx, orgId, &body)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

