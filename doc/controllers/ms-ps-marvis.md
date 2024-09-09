# MS Ps Marvis

```go
mSPsMarvis := client.MSPsMarvis()
```

## Class Name

`MSPsMarvis`


# Count Msps Marvis Actions

Count Marvis actions

```go
CountMspsMarvisActions(
    ctx context.Context,
    mspId uuid.UUID,
    distinct *models.MspMarvisSuggestionsCountDistinctEnum,
    limit *int,
    page *int) (
    models.ApiResponse[models.ResponseCountMarvisActions],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `distinct` | [`*models.MspMarvisSuggestionsCountDistinctEnum`](../../doc/models/msp-marvis-suggestions-count-distinct-enum.md) | Query, Optional | **Default**: `"org_id"` |
| `limit` | `*int` | Query, Optional | **Default**: `100`<br>**Constraints**: `>= 0` |
| `page` | `*int` | Query, Optional | **Default**: `1`<br>**Constraints**: `>= 1` |

## Response Type

[`models.ResponseCountMarvisActions`](../../doc/models/response-count-marvis-actions.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

distinct := models.MspMarvisSuggestionsCountDistinctEnum("org_id")

limit := 100

page := 1

apiResponse, err := mSPsMarvis.CountMspsMarvisActions(ctx, mspId, &distinct, &limit, &page)
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
  "distinct": "status",
  "limit": 1000,
  "results": [
    {
      "count": 24,
      "status": "002e176a-0000-000-1111-002e208b20e1"
    },
    {
      "count": 12,
      "status": "2d3f176a-0000-000-2222-002e208f176a"
    },
    {
      "count": 15,
      "status": "08b2176a-0000-000-3333-002e208b2d3f"
    }
  ],
  "total": 3
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

