
# Tunnel Provider Options Zscaler Sub Location

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelProviderOptionsZscalerSubLocation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AupAcceptanceRequired` | `*bool` | Optional | **Default**: `true` |
| `AupExpire` | `*int` | Optional | days before AUP is requested again<br>**Default**: `1` |
| `AupSslProxy` | `*bool` | Optional | proxy HTTPs traffic, requiring Zscaler cert to be installed in browser<br>**Default**: `false` |
| `DownloadMbps` | `*int` | Optional | the download bandwidth cap of the link, in Mbps |
| `EnableAup` | `*bool` | Optional | if `use_xff`==`true`, display Acceptable Use Policy (AUP) |
| `EnableCaution` | `*bool` | Optional | when `enforce_authentication`==`false`, display caution notification for non-authenticated users<br>**Default**: `false` |
| `EnforceAuthentication` | `*bool` | Optional | **Default**: `false` |
| `Subnets` | `[]string` | Optional | - |
| `UploadMbps` | `*int` | Optional | the download bandwidth cap of the link, in Mbps |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "aup_acceptance_required": true,
  "aup_expire": 1,
  "aup_ssl_proxy": false,
  "download_mbps": 200,
  "enable_caution": false,
  "enforce_authentication": false,
  "subnets": [
    "172.16.8.0/24",
    "172.16.32.0/23"
  ],
  "upload_mbps": 200,
  "enable_aup": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

