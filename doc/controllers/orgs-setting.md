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

Replace the organization wireless client blocklist with the supplied client MAC addresses. The list can contain up to 1000 MAC addresses; retrieve the current list from the `blacklist_url` field in organization settings.

```go
CreateOrgWirelessClientsBlocklist(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.MacAddresses) (
    models.ApiResponse[models.MacAddresses],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.MacAddresses`](../../doc/models/mac-addresses.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Delete Org Wireless Clients Blocklist

Clear the organization wireless client blocklist by removing all blocked client MAC addresses.

```go
DeleteOrgWirelessClientsBlocklist(
    ctx context.Context,
    orgId uuid.UUID) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance.

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

resp, err := orgsSetting.DeleteOrgWirelessClientsBlocklist(ctx, orgId)
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
    fmt.Println(resp.StatusCode)
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


# Get Org Settings

Return organization-wide settings, including feature flags, automatic device assignment rules, management connectivity, packet capture, security controls, and integration configuration.

```go
GetOrgSettings(
    ctx context.Context,
    orgId uuid.UUID) (
    models.ApiResponse[models.OrgSetting],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.OrgSetting](../../doc/models/org-setting.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

apiResponse, err := orgsSetting.GetOrgSettings(ctx, orgId)
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
  "auto_device_naming": {
    "enable": true,
    "rules": [
      {
        "match_device": "ap",
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
    "expiry_in_days": 180,
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Set Org Custom Bucket

Start custom packet capture bucket setup by saving the bucket name and having Mist write a `MIST_TOKEN` file to the bucket. Complete ownership verification with the verify endpoint by submitting the token contents.

```go
SetOrgCustomBucket(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.PcapBucket) (
    models.ApiResponse[models.ResponsePcapBucketConfig],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.PcapBucket`](../../doc/models/pcap-bucket.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

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
  "bucket": "company-private-pcap",
  "detail": "failed to write bucket - 403 AccessDenied"
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


# Update Org Settings

Update organization-wide settings such as automatic device assignment rules, management connectivity, packet capture, password policy, security controls, tags, and integration options.

```go
UpdateOrgSettings(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.OrgSetting) (
    models.ApiResponse[models.OrgSetting],
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.OrgSetting`](../../doc/models/org-setting.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

This method returns an [`ApiResponse`](../../doc/api-response.md) instance. The `Data` property of this instance returns the response data which is of type [models.OrgSetting](../../doc/models/org-setting.md).

## Example Usage

```go
ctx := context.Background()

orgId := uuid.MustParse("000000ab-00ab-00ab-00ab-0000000000ab")

body := models.OrgSetting{
    AllowMist:                    models.ToPointer(false),
    ApUpdownThreshold:            models.NewOptional(models.ToPointer(0)),
    DeviceUpdownThreshold:        models.NewOptional(models.ToPointer(0)),
    DisablePcap:                  models.ToPointer(false),
    DisableRemoteShell:           models.ToPointer(false),
    GatewayUpdownThreshold:       models.NewOptional(models.ToPointer(10)),
    SwitchUpdownThreshold:        models.NewOptional(models.ToPointer(0)),
    UiIdleTimeout:                models.ToPointer(10),
    UiNoTracking:                 models.ToPointer(false),
}

apiResponse, err := orgsSetting.UpdateOrgSettings(ctx, orgId, &body)
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
  "auto_device_naming": {
    "enable": true,
    "rules": [
      {
        "match_device": "ap",
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
    "expiry_in_days": 365,
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
| 401 | Unauthorized | [`ResponseHttp401Exception`](../../doc/models/response-http-401-exception.md) |
| 403 | Permission Denied | [`ResponseHttp403Exception`](../../doc/models/response-http-403-exception.md) |
| 404 | Not found. The API endpoint doesn’t exist or resource doesn’ t exist | [`ResponseHttp404Exception`](../../doc/models/response-http-404-exception.md) |
| 429 | Too Many Request. The API Token used for the request reached the 5000 API Calls per hour threshold | [`ResponseHttp429Exception`](../../doc/models/response-http-429-exception.md) |


# Verify Org Custom Bucket

Verify ownership of a custom packet capture bucket by submitting the token read from the `MIST_TOKEN` file. If verification succeeds, Mist creates a `VERIFIED` file in the bucket.

```go
VerifyOrgCustomBucket(
    ctx context.Context,
    orgId uuid.UUID,
    body *models.PcapBucketVerify) (
    http.Response,
    error)
```

## Authentication

This endpoint requires [apiToken](../../doc/auth/custom-header-signature.md) **OR** [csrfToken](../../doc/auth/custom-header-signature-1.md)

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orgId` | `uuid.UUID` | Template, Required | - |
| `body` | [`*models.PcapBucketVerify`](../../doc/models/pcap-bucket-verify.md) | Body, Optional | Request Body |

## Response Type

**200**: OK

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
    fmt.Println(resp.StatusCode)
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

