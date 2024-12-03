
# Stats Gateway Service Status

*This model accepts additional fields of type interface{}.*

## Structure

`StatsGatewayServiceStatus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AppidInstallResult` | `*string` | Optional | - |
| `AppidInstallTimestamp` | `*string` | Optional | - |
| `AppidStatus` | `*string` | Optional | - |
| `AppidVersion` | `*int` | Optional | - |
| `EwfStatus` | `*string` | Optional | - |
| `IdpInstallResult` | `*string` | Optional | - |
| `IdpInstallTimestamp` | `*string` | Optional | - |
| `IdpPolicy` | `*string` | Optional | - |
| `IdpStatus` | `*string` | Optional | - |
| `IdpUpdateTimestamp` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "appid_install_result": "appid_install_result6",
  "appid_install_timestamp": "appid_install_timestamp2",
  "appid_status": "appid_status2",
  "appid_version": 106,
  "ewf_status": "ewf_status6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

