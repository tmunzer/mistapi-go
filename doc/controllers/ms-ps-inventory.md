# MS Ps Inventory

```go
mSPsInventory := client.MSPsInventory()
```

## Class Name

`MSPsInventory`


# Get Msp Inventory by Mac

Return the inventory record for a device MAC address owned by an organization under this MSP, including the owning organization and site when available.

```go
GetMspInventoryByMac(
    ctx context.Context,
    mspId uuid.UUID,
    deviceMac string) (
    models.ApiResponse[models.ResponseMspInventoryDevice],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mspId` | `uuid.UUID` | Template, Required | - |
| `deviceMac` | `string` | Template, Required | **Constraints**: *Pattern*: `^[0-9a-fA-F]{12}$` |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponseMspInventoryDevice](../../doc/models/response-msp-inventory-device.md).

## Example Usage

```go
ctx := context.Background()

mspId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

deviceMac := "0000000000ab"

apiResponse, err := mSPsInventory.GetMspInventoryByMac(ctx, mspId, deviceMac)
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |

