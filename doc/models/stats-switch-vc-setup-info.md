
# Stats Switch Vc Setup Info

## Structure

`StatsSwitchVcSetupInfo`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigType` | `*string` | Optional | - |
| `CurrentStats` | `*string` | Optional | - |
| `ErrMissingDevIdFpc` | `*bool` | Optional | - |
| `LastUpdate` | `*float64` | Optional | - |
| `RequestTime` | `*float64` | Optional | - |
| `RequestType` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "config_type": "nonprovisioned",
  "current_stats": "VCSETUP_WAITING",
  "request_type": "vc_create",
  "err_missing_dev_id_fpc": false,
  "last_update": 184.52,
  "request_time": 38.16
}
```

