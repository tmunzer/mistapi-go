# Orgs Vars

```go
orgsVars := client.OrgsVars()
```

## Class Name

`OrgsVars`


# Search Org Vars

Search vars

Example: /api/v1/orgs/{org_id}/vars/search?vars=*

```go
SearchOrgVars(
    ctx context.Context,
    orgId uuid.UUID,
    siteId *string,
    mVar *string,
    src *models.VarSourceEnum,
    limit *int,
    sort *string,
    searchAfter *string) (
    models.ApiResponse[models.ResponseSearchVar],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | Filter results by site identifier. Accepts multiple comma-separated values. |
| `mVar` | `*string` | Query, Optional | Filter variable results by variable name. Accepts multiple comma-separated values. |
| `src` | [`*models.VarSourceEnum`](../../doc/models/var-source-enum.md) | Query, Optional | Filter results by source. enum: `deviceprofile`, `site` |
| `limit` | `*int` | Query, Optional | Maximum number of results to return per page<br><br>**Default**: `100`<br><br>**Constraints**: `>= 0` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |
| `searchAfter` | `*string` | Query, Optional | Pagination cursor for retrieving subsequent pages of results. This value is automatically populated by Mist in the `next` URL from the previous response and should not be manually constructed. |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSearchVar](../../doc/models/response-search-var.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

siteId := "00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002"

mVar := "guest_end,guest_net"

limit := 100

sort := "-site_id"

apiResponse, err := orgsVars.SearchOrgVars(ctx, orgId, &siteId, &mVar, nil, &limit, &sort, nil)
if err != nil {
    switch typedErr := err.(type) {
        case *errors.ResponseHttp400:
            log.Fatalln("ResponseHttp400Exception: ", typedErr)
        case *errors.ResponseHttp401:
            log.Fatalln("ResponseHttp401Exception: ", typedErr)
        case *errors.ResponseHttp403:
            log.Fatalln("ResponseHttp403Exception: ", typedErr)
        case *errors.ResponseHttp404:
            log.Fatalln("ResponseHttp404Exception: ", typedErr)
        case *errors.ResponseHttp429:
            log.Fatalln("ResponseHttp429Exception: ", typedErr)
        default:
            log.Fatalln(err)
    }
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "end": 1693952979,
  "limit": 10,
  "results": [
    {
      "created_time": 1618457655.384858,
      "modified_time": 1693610886.477805,
      "org_id": "0c160b7f-1027-4cd1-923b-744534c4b070",
      "site_id": "1519f016-4e41-47c0-a396-cce4d04bac0b",
      "src": "site",
      "var": "mvp"
    }
  ],
  "start": 1693949379,
  "total": 1
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

