# Orgs Setting

```go
orgsSetting := client.OrgsSetting()
```

## Class Name

`OrgsSetting`

## Methods

* [Create Org Wireless Clients Blocklist](../../doc/controllers/orgs-setting.md#create-org-wireless-clients-blocklist)
* [Delete Org Wireless Clients Blocklist](../../doc/controllers/orgs-setting.md#delete-org-wireless-clients-blocklist)
* [Get Org Settings](../../doc/controllers/orgs-setting.md#get-org-settings)
* [Set Org Custom Bucket](../../doc/controllers/orgs-setting.md#set-org-custom-bucket)
* [Update Org Settings](../../doc/controllers/orgs-setting.md#update-org-settings)
* [Verify Org Custom Bucket](../../doc/controllers/orgs-setting.md#verify-org-custom-bucket)


# Create Org Wireless Clients Blocklist

Create Org Blacklist Client List.

If there is already a blacklist, this API will replace it with the new one.

Max number of blacklist clients is 1000.

Retrieve the current blacklisted clients from `blacklist_url` under Org:Setting

```go
CreateOrgWirelessClientsBlocklist(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.MacAddresses](../../doc/models/mac-addresses.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.MacAddresses{
    Macs:                 []string{
        "18-65-90-de-f4-c6",
        "84-89-ad-5d-69-0d",
    },
}

apiResponse, err := orgsSetting.CreateOrgWirelessClientsBlocklist(ctx, orgId, &body)
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
  "macs": [
    "18-65-90-de-f4-c6",
    "84-89-ad-5d-69-0d"
  ]
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


# Delete Org Wireless Clients Blocklist

Delete Org Blacklist Station Clients

```go
DeleteOrgWirelessClientsBlocklist(
    ctx context.Context,
    orgId uuid.UUID) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSetting.DeleteOrgWirelessClientsBlocklist(ctx, orgId)
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


# Get Org Settings

Get Org Settings

```go
GetOrgSettings(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.OrgSetting],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.OrgSetting](../../doc/models/org-setting.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSetting.GetOrgSettings(ctx, orgId)
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
  "auto_device_naming": {
    "enable": true,
    "rules": [
      {
        "match_device_type": "ap",
        "prefix": "MIST-",
        "src": "lldp_port_desc"
      }
    ]
  },
  "auto_deviceprofile_assignment": {
    "enable": true,
    "rules": [
      {
        "expression": "string",
        "model": "string",
        "prefix": "string",
        "src": "name",
        "subnet": "string",
        "suffix": "string",
        "value": "string"
      }
    ]
  },
  "auto_site_assignment": {
    "enable": true,
    "rules": [
      {
        "expression": "string",
        "model": "string",
        "prefix": "string",
        "src": "name",
        "subnet": "string",
        "suffix": "string",
        "value": "string"
      }
    ]
  },
  "cacerts": [
    "string"
  ],
  "cloudshark": {
    "apitoken": "string",
    "url": "string"
  },
  "created_time": 0,
  "device_cert": {
    "cert": "string",
    "key": "string"
  },
  "device_updown_threshold": 0,
  "disable_pcap": true,
  "id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "installer": {
    "allow_all_sites": true,
    "extra_site_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "grace_period": 0
  },
  "mgmt": {
    "mxtunnel_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "use_mxtunnel": true,
    "use_wxtunnel": true
  },
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "password_policy": {
    "enabled": true,
    "freshness": 0,
    "min_length": 8,
    "requires_special_char": true,
    "requires_two_factor_auth": true
  },
  "pcap": {
    "bucket": "string",
    "max_pkt_len": 0
  },
  "pcap_bucket_verified": true,
  "remote_syslog": {
    "enabled": true,
    "send_to_all_servers": true,
    "servers": [
      {
        "facility": "conflict-log",
        "host": "string",
        "port": 0,
        "protocol": "udp",
        "severity": "any",
        "tag": "string"
      }
    ]
  },
  "security": {
    "disable_local_ssh": true,
    "fips_zeroize_password": "string",
    "limit_ssh_access": true
  },
  "tags": [
    "string"
  ],
  "ui_idle_timeout": 0
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


# Set Org Custom Bucket

Provide Customer Bucket Name

Setting up Custom PCAP Bucket Involves the following:

* provide the bucket name
* we’ll attempt to write a file MIST_TOKEN
* you have to verify the ownership of the bucket by providing the content of the MIST_TOKEN

```go
SetOrgCustomBucket(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.PcapBucket) (
    models.ApiResponse[models.ResponsePcapBucketConfig],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.PcapBucket`](../../doc/models/pcap-bucket.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.ResponsePcapBucketConfig](../../doc/models/response-pcap-bucket-config.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.PcapBucket{
    Bucket:               "company-private-pcap",
}

apiResponse, err := orgsSetting.SetOrgCustomBucket(ctx, orgId, &body)
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
  "bucket": "company-private-pcap",
  "detail": "failed to write bucket - 403 AccessDenied"
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


# Update Org Settings

Update Org Settings

```go
UpdateOrgSettings(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OrgSetting) (
    models.ApiResponse[models.OrgSetting],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.OrgSetting`](../../doc/models/org-setting.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.OrgSetting](../../doc/models/org-setting.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.OrgSetting{
    ApUpdownThreshold:           models.NewOptional(models.ToPointer(0)),
    DeviceUpdownThreshold:       models.NewOptional(models.ToPointer(0)),
    DisablePcap:                 models.ToPointer(false),
    DisableRemoteShell:          models.ToPointer(false),
    GatewayUpdownThreshold:      models.NewOptional(models.ToPointer(10)),
    SwitchUpdownThreshold:       models.NewOptional(models.ToPointer(0)),
    UiIdleTimeout:               models.ToPointer(10),
    UiNoTracking:                models.ToPointer(false),
}

apiResponse, err := orgsSetting.UpdateOrgSettings(ctx, orgId, &body)
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
  "auto_device_naming": {
    "enable": true,
    "rules": [
      {
        "match_device_type": "ap",
        "prefix": "MIST-",
        "src": "lldp_port_desc"
      }
    ]
  },
  "auto_deviceprofile_assignment": {
    "enable": true,
    "rules": [
      {
        "expression": "string",
        "model": "string",
        "prefix": "string",
        "src": "name",
        "subnet": "string",
        "suffix": "string",
        "value": "string"
      }
    ]
  },
  "auto_site_assignment": {
    "enable": true,
    "rules": [
      {
        "expression": "string",
        "model": "string",
        "prefix": "string",
        "src": "name",
        "subnet": "string",
        "suffix": "string",
        "value": "string"
      }
    ]
  },
  "cacerts": [
    "string"
  ],
  "cloudshark": {
    "apitoken": "string",
    "url": "string"
  },
  "device_cert": {
    "cert": "string",
    "key": "string"
  },
  "disable_pcap": true,
  "installer": {
    "allow_all_sites": true,
    "extra_site_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "grace_period": 0
  },
  "mgmt": {
    "mxtunnel_ids": [
      "b069b358-4c97-5319-1f8c-7c5ca64d6ab1"
    ],
    "use_mxtunnel": true,
    "use_wxtunnel": true
  },
  "modified_time": 0,
  "msp_id": "b069b358-4c97-5319-1f8c-7c5ca64d6ab1",
  "name": "string",
  "password_policy": {
    "enabled": true,
    "freshness": 0,
    "min_length": 8,
    "requires_special_char": true,
    "requires_two_factor_auth": true
  },
  "pcap": {
    "bucket": "string",
    "max_pkt_len": 0
  },
  "pcap_bucket_verified": true,
  "remote_syslog": {
    "enabled": true,
    "send_to_all_servers": true,
    "servers": [
      {
        "facility": "change-log",
        "host": "string",
        "port": 0,
        "protocol": "udp",
        "severity": "critical",
        "tag": "string"
      }
    ]
  },
  "security": {
    "disable_local_ssh": true,
    "fips_zeroize_password": "string",
    "limit_ssh_access": true
  },
  "tags": [
    "string"
  ],
  "ui_idle_timeout": 0
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


# Verify Org Custom Bucket

Verify Customer PCAP Bucket

**Note**: If successful, a "VERIFIED" file will be created in the bucket

```go
VerifyOrgCustomBucket(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.PcapBucketVerify) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.PcapBucketVerify`](../../doc/models/pcap-bucket-verify.md) | Body, Optional | Request Body |

## Response Type

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.PcapBucketVerify{
    Bucket:               "company-private-pcap",
    VerifyToken:          "eyJhbGciOiJIUzI1J9.eyJzdWIiOiIxMjM0joiMjgxOG5MDIyfQ.2rzcRvMA3Eg09NnjCAC-1EWMRtxAnFDM",
}

resp, err := orgsSetting.VerifyOrgCustomBucket(ctx, orgId, &body)
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

