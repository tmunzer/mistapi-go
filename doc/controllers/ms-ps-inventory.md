# MS Ps Inventory

```go
mSPsInventory := client.MSPsInventory()
```

## Class Name

`MSPsInventory`


# Get Msp Inventory by Mac

Get Inventoy By device MAC address

```go
GetMspInventoryByMac(
    ctx context.Context,
    mspId uuid.UUID,
    deviceMac string) (
    models.ApiResponse[models.ResponseMspInventoryDevice],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | - |

## Response Type

[`models.ResponseMspInventoryDevice`](../../doc/models/response-msp-inventory-device.md)

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

apiResponse, err := mSPsInventory.GetMspInventoryByMac(ctx, mspId, deviceMac)
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
  "mac": "5c5b35000018",
  "model": "AP200",
  "org_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "serial": "FXLH2015150025",
  "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
  "type": "ap"
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

