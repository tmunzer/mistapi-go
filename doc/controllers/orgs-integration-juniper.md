# Orgs Integration Juniper

```go
orgsIntegrationJuniper := client.OrgsIntegrationJuniper()
```

## Class Name

`OrgsIntegrationJuniper`

## Methods

* [Link Org to Juniper Juniper Account](../../doc/controllers/orgs-integration-juniper.md#link-org-to-juniper-juniper-account)
* [Unlink Org From Juniper Customer Id](../../doc/controllers/orgs-integration-juniper.md#unlink-org-from-juniper-customer-id)


# Link Org to Juniper Juniper Account

Link Juniper Accounts

```go
LinkOrgToJuniperJuniperAccount(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountJuniperConfig) (
    models.ApiResponse[models.AccountJuniperInfo],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AccountJuniperConfig`](../../doc/models/account-juniper-config.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.AccountJuniperInfo](../../doc/models/account-juniper-info.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.AccountJuniperConfig{
    Password:             "password",
    Username:             "john@nmo.com",
}

apiResponse, err := orgsIntegrationJuniper.LinkOrgToJuniperJuniperAccount(ctx, orgId, &body)
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
  "account": {
    "linked_by": "John Smith (john@abccorp.com)",
    "name": "ABC Corp"
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Account already linked | [`ResponseDetailStringException`](../../doc/models/response-detail-string-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Unlink Org From Juniper Customer Id

Unlink Juniper Customer ID
`linked_by` field is only required if there are duplicate account_names.

```go
UnlinkOrgFromJuniperCustomerId(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.AccountJuniperInfo) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.AccountJuniperInfo`](../../doc/models/account-juniper-info.md) | Body, Optional | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")



resp, err := orgsIntegrationJuniper.UnlinkOrgFromJuniperCustomerId(ctx, orgId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
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

