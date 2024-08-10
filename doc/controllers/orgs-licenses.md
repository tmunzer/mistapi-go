# Orgs Licenses

```go
orgsLicenses := client.OrgsLicenses()
```

## Class Name

`OrgsLicenses`

## Methods

* [Claim Org License](../../doc/controllers/orgs-licenses.md#claim-org-license)
* [Get Org Licences by Site](../../doc/controllers/orgs-licenses.md#get-org-licences-by-site)
* [Get Org Licences Summary](../../doc/controllers/orgs-licenses.md#get-org-licences-summary)
* [Move or Delete Org License to Another Org](../../doc/controllers/orgs-licenses.md#move-or-delete-org-license-to-another-org)


# Claim Org License

Claim Org licenses / activation codes

```go
ClaimOrgLicense(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.ClaimActivation) (
    models.ApiResponse[models.ResponseClaimLicense],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.ClaimActivation`](../../doc/models/claim-activation.md) | Body, Optional | Request Body |

## Response Type

[`models.ResponseClaimLicense`](../../doc/models/response-claim-license.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.ClaimActivation{
    Code:       "ZHT3K-H36DT-MG85D-M61AC",
    Type:       models.ClaimTypeEnum("all"),
}

apiResponse, err := orgsLicenses.ClaimOrgLicense(ctx, orgId, &body)
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
  "inventory_added": [
    {
      "mac": "5c5b35000018",
      "magic": "6JG8EPTFV2A9Z2N",
      "model": "AP41",
      "serial": "FXLH2015150025",
      "type": "ap"
    }
  ],
  "inventory_duplicated": [
    {
      "mac": "5c5b35000012",
      "magic": "DVH4VSNMSZPDXBR",
      "model": "AP41",
      "serial": "FXLH2015150027",
      "type": "ap"
    }
  ],
  "license_added": [
    {
      "end": 1520380800,
      "quantity": 180,
      "start": 1504828800,
      "type": "SUB-MAN"
    },
    {
      "end": 1520380800,
      "quantity": 120,
      "start": 1504828800,
      "type": "SUB-LOC"
    }
  ],
  "license_duplicated": [
    {
      "end": 1520380800,
      "quantity": 180,
      "start": 1504828800,
      "type": "SUB-MAN"
    }
  ],
  "license_error": [
    {
      "order": "00000464",
      "reason": ""
    }
  ]
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | invalid key (or already used) | `ApiError` |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Org Licences by Site

Get Licenses Usage by Sites
This shows license usage (i.e. needed) based on the features enabled for the site.

```go
GetOrgLicencesBySite(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[[]models.LicenseUsageOrg],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`[]models.LicenseUsageOrg`](../../doc/models/license-usage-org.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsLicenses.GetOrgLicencesBySite(ctx, orgId)
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
[
  {
    "fully_loaded": {
      "SUB-LOC": 30,
      "SUB-MAN": 80
    },
    "num_devices": 80,
    "site_id": "4ac1dcf4-9d8b-7211-65c4-057819f0862b",
    "usages": {
      "SUB-LOC": 30,
      "SUB-MAN": 60
    }
  }
]
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Bad Syntax | [`ResponseHttp400Exception`](../../doc/models/response-http-400-exception.md) |
| 401 | Unauthorized | [`ResponseHttp401ErrorException`](../../doc/models/response-http-401-error-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403ErrorException`](../../doc/models/response-http-403-error-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429ErrorException`](../../doc/models/response-http-429-error-exception.md) |


# Get Org Licences Summary

Get the list of licenses

```go
GetOrgLicencesSummary(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.License],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

[`models.License`](../../doc/models/license.md)

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsLicenses.GetOrgLicencesSummary(ctx, orgId)
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
  "amendments": [
    {
      "created_time": 1644684842,
      "end_time": 1744156799,
      "id": "ff0a14f6-1234-5678-90ab-c8e64d4bc6c6",
      "modified_time": 1644684842,
      "quantity": -1,
      "start_time": 1632873600,
      "subscription_id": "VNA-000000af",
      "type": "SUB-VNA"
    },
    {
      "created_time": 1644684842,
      "end_time": 1744156799,
      "id": "c1c28812-1234-5678-90ab-dc95680da61e",
      "modified_time": 1644684842,
      "quantity": -1,
      "start_time": 1632873600,
      "subscription_id": "MAN-000008be",
      "type": "SUB-MAN"
    },
    {
      "created_time": 1644684842,
      "end_time": 1744243199,
      "id": "96c0a41f-1234-5678-90ab-afe74817e9fd",
      "modified_time": 1644684842,
      "quantity": -1,
      "start_time": 1586476800,
      "subscription_id": "EX24-000000bc",
      "type": "SUB-EX24"
    }
  ],
  "entitled": {
    "SUB-ENG": 26,
    "SUB-EX24": 9,
    "SUB-MAN": 26,
    "SUB-VNA": 26
  },
  "licenses": [
    {
      "created_time": 1555353534,
      "end_time": 1586822399,
      "id": "693a41a6-1234-5678-90ab-f53dbd3a31c0",
      "modified_time": 1555353534,
      "order_id": "00000000",
      "org_id": "9777c1a0-1234-5678-90ab-02e208b2d34f",
      "quantity": 2,
      "remaining_quantity": 0,
      "start_time": 1555286400,
      "subscription_id": "VNA-000000aa",
      "type": "SUB-VNA"
    },
    {
      "created_time": 1576132516,
      "end_time": 1586822399,
      "id": "656607cf-1234-5678-90ab-fc9035614ea5",
      "modified_time": 1576132516,
      "order_id": "00000000",
      "org_id": "9777c1a0-1234-5678-90ab-02e208b2d34f",
      "quantity": 8,
      "remaining_quantity": 0,
      "start_time": 1576022400,
      "subscription_id": "VNA-000000ab",
      "type": "SUB-VNA"
    },
    {
      "created_time": 1579204568,
      "end_time": 1730764800,
      "id": "db50d0bc-1234-5678-90ab-e439958cb06b",
      "modified_time": 1579204568,
      "order_id": "00000000",
      "org_id": "9777c1a0-1234-5678-90ab-02e208b2d34f",
      "quantity": 2,
      "remaining_quantity": 2,
      "start_time": 1572998400,
      "subscription_id": "MAN-000000ac",
      "type": "SUB-MAN"
    },
    {
      "created_time": 1579204568,
      "end_time": 1730764800,
      "id": "2ff9e84a-1234-5678-90ab-fb9ec0726e01",
      "modified_time": 1579204568,
      "order_id": "00000000",
      "org_id": "9777c1a0-1234-5678-90ab-02e208b2d34f",
      "quantity": 2,
      "remaining_quantity": 2,
      "start_time": 1572998400,
      "subscription_id": "ENG-000000ad",
      "type": "SUB-ENG"
    },
    {
      "created_time": 1579204568,
      "end_time": 1730764800,
      "id": "16df7ea6-1234-5678-90ab-78018cd4024d",
      "modified_time": 1579204568,
      "order_id": "00000000",
      "org_id": "9777c1a0-1234-5678-90ab-02e208b2d34f",
      "quantity": 2,
      "remaining_quantity": 2,
      "start_time": 1572998400,
      "subscription_id": "VNA-000000ae",
      "type": "SUB-VNA"
    },
    {
      "created_time": 1586237081,
      "end_time": 1744243199,
      "id": "1b6f68d5-1234-5678-90ab-70d3e6d18c73",
      "modified_time": 1586237081,
      "order_id": "00000000",
      "org_id": "9777c1a0-1234-5678-90ab-02e208b2d34f",
      "quantity": 14,
      "remaining_quantity": 14,
      "start_time": 1586563200,
      "subscription_id": "VNA-000000af",
      "type": "SUB-VNA"
    },
    {
      "created_time": 1586237097,
      "end_time": 1744243199,
      "id": "1375c9bf-1234-5678-90ab-9c636708c89e",
      "modified_time": 1586237097,
      "order_id": "00000000",
      "org_id": "9777c1a0-1234-5678-90ab-02e208b2d34f",
      "quantity": 14,
      "remaining_quantity": 14,
      "start_time": 1586563200,
      "subscription_id": "MAN-000000ba",
      "type": "SUB-MAN"
    },
    {
      "created_time": 1586237137,
      "end_time": 1744243199,
      "id": "5974e979-1234-5678-90ab-438f833ec1c9",
      "modified_time": 1586237137,
      "order_id": "00000000",
      "org_id": "9777c1a0-1234-5678-90ab-02e208b2d34f",
      "quantity": 14,
      "remaining_quantity": 14,
      "start_time": 1586563200,
      "subscription_id": "ENG-000000bb",
      "type": "SUB-ENG"
    },
    {
      "created_time": 1629947267,
      "end_time": 1744243199,
      "id": "340a9cb3-1234-5678-90ab-b009344dbf3c",
      "modified_time": 1629947267,
      "order_id": "00000000",
      "org_id": "9777c1a0-1234-5678-90ab-02e208b2d34f",
      "quantity": 10,
      "remaining_quantity": 9,
      "start_time": 1586476800,
      "subscription_id": "EX24-000000bc",
      "type": "SUB-EX24"
    },
    {
      "created_time": 1632941870,
      "end_time": 1744156799,
      "id": "9b599b0f-1234-5678-90ab-406081b58e7f",
      "modified_time": 1632941870,
      "order_id": "00000000",
      "org_id": "9777c1a0-1234-5678-90ab-02e208b2d34f",
      "quantity": 10,
      "remaining_quantity": 10,
      "start_time": 1632873600,
      "subscription_id": "ENG-000000bd",
      "type": "SUB-ENG"
    },
    {
      "created_time": 1632941882,
      "end_time": 1744156799,
      "id": "d6d8ead3-1234-5678-90ab-98badeac7287",
      "modified_time": 1632941882,
      "order_id": "00000000",
      "org_id": "9777c1a0-1234-5678-90ab-02e208b2d34f",
      "quantity": 11,
      "remaining_quantity": 9,
      "start_time": 1632873600,
      "subscription_id": "MAN-000008be",
      "type": "SUB-MAN"
    }
  ],
  "summary": {
    "SUB-ENG": 18,
    "SUB-EX24": 3,
    "SUB-MAN": 22,
    "SUB-VNA": 20
  }
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


# Move or Delete Org License to Another Org

Move, Undo Move or Delete Org License to Another Org
If the admin has admin privilege against the `org_id` and `dst_org_id`, he can move some of the licenses to another Org. Given that:

1. the specified license is currently active
2. there’s enough licenses left in the specified license (by subscription_id)
3. there will still be enough entitled licenses for the type of license after the amendment

```go
MoveOrDeleteOrgLicenseToAnotherOrg(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OrgLicenseAction) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.OrgLicenseAction`](../../doc/models/org-license-action.md) | Body, Optional | Request Body |

## Response Type

``

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.OrgLicenseAction{
    Notes:          models.ToPointer("customer notes"),
    Op:             models.OrgLicenseActionOperationEnum("annotate"),
    SubscriptionId: models.ToPointer("SUB-000144"),
}

resp, err := orgsLicenses.MoveOrDeleteOrgLicenseToAnotherOrg(ctx, orgId, &body)
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

