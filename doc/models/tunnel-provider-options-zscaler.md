
# Tunnel Provider Options Zscaler

for zscaler-ipsec and zscaler-gre

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelProviderOptionsZscaler`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AupAcceptanceRequired` | `*bool` | Optional | **Default**: `true` |
| `AupExpire` | `*int` | Optional | days before AUP is requested again<br>**Default**: `1` |
| `AupSslProxy` | `*bool` | Optional | proxy HTTPs traffic, requiring Zscaler cert to be installed in browser<br>**Default**: `false` |
| `DownloadMbps` | `*int` | Optional | the download bandwidth cap of the link, in Mbps |
| `EnableAup` | `*bool` | Optional | if `use_xff`==`true`, display Acceptable Use Policy (AUP)<br>**Default**: `false` |
| `EnableCaution` | `*bool` | Optional | when `enforce_authentication`==`false`, display caution notification for non-authenticated users<br>**Default**: `false` |
| `EnforceAuthentication` | `*bool` | Optional | **Default**: `false` |
| `Name` | `*string` | Optional | - |
| `SubLocations` | [`[]models.TunnelProviderOptionsZscalerSubLocation`](../../doc/models/tunnel-provider-options-zscaler-sub-location.md) | Optional | if `use_xff`==`true` |
| `UploadMbps` | `*int` | Optional | the download bandwidth cap of the link, in Mbps |
| `UseXff` | `*bool` | Optional | location uses proxy chaining to forward traffic |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "aup_acceptance_required": true,
  "aup_expire": 1,
  "aup_ssl_proxy": false,
  "download_mbps": 200,
  "enable_aup": false,
  "enable_caution": false,
  "enforce_authentication": false,
  "name": "guest-wifi",
  "upload_mbps": 200,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

