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
    page *int,
    sort *string) (
    models.ApiResponse[models.ResponseSearchVar],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `siteId` | `*string` | Query, Optional | - |
| `mVar` | `*string` | Query, Optional | - |
| `src` | [`*models.VarSourceEnum`](../../doc/models/var-source-enum.md) | Query, Optional | - |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br><br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br><br>**Constraints**: `>= 1` |
| `sort` | `*string` | Query, Optional | On which field the list should be sorted, -prefix represents DESC order<br><br>**Default**: `"timestamp"` |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseSearchVar](../../doc/models/response-search-var.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

limit := 100

page := 1

sort := "-site_id"

apiResponse, err := orgsVars.SearchOrgVars(ctx, orgId, nil, nil, nil, &limit, &page, &sort)
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
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |

