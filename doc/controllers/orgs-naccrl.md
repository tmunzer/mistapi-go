# Orgs NACCRL

```go
orgsNACCRL := client.OrgsNACCRL()
```

## Class Name

`OrgsNACCRL`

## Methods

* [Delete Org Nac Crl](../../doc/controllers/orgs-naccrl.md#delete-org-nac-crl)
* [Get Org Nac Crl](../../doc/controllers/orgs-naccrl.md#get-org-nac-crl)
* [Import Org Nac Crl](../../doc/controllers/orgs-naccrl.md#import-org-nac-crl)


# Delete Org Nac Crl

Delete NAC Org CRL file is a DELETE request to delete CRL file identified by its ID (ID assigned on file upload/creation)

```go
DeleteOrgNacCrl(
    ctx context.Context,
    orgId uuid.UUID,
    naccrlId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `naccrlId` | `uuid.UUID` | Template, Required | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

naccrlId := uuid.MustParse("00002294-0000-0000-0000-000000000000")

resp, err := orgsNACCRL.DeleteOrgNacCrl(ctx, orgId, naccrlId)
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
| 401 | Unauthorized | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 403 | Permission Denied | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Get Org Nac Crl

Returns all uploaded CRL file IDs with names for the orgI

```go
GetOrgNacCrl(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.ResponseNacCrlFiles],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.ResponseNacCrlFiles`](../../doc/models/response-nac-crl-files.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsNACCRL.GetOrgNacCrl(ctx, orgId)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |


# Import Org Nac Crl

Import NAC Org CRL file is a multipart POST which has a .crl file to allow user to manually upload a Certificate Revocation List file. We support one file upload per Issuer. Re-uploads for the same issuer will overwrite the existing file.

```go
ImportOrgNacCrl(
    ctx context.Context,
    orgId uuid.UUID,
    file *models.FileWrapper,
    json *string) (
    models.ApiResponse[models.NacCrlFile],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `file` | `*models.FileWrapper` | Form, Optional | a binary .crl file |
| `json` | `*string` | Form, Optional | json string with name for .crl file (optional) |

## Response Type

[`models.NacCrlFile`](../../doc/models/nac-crl-file.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")





apiResponse, err := orgsNACCRL.ImportOrgNacCrl(ctx, orgId, nil, nil)
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
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |

