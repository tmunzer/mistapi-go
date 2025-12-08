
# Device Search Radius Stat

## Structure

`DeviceSearchRadiusStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthAccepts` | `*int` | Optional | Number of accepted authentication requests |
| `AuthRejects` | `*int` | Optional | Number of rejected authentication requests |
| `AuthServerStatus` | [`*models.DeviceSearchRadiusFilterStatusEnum`](../../doc/models/device-search-radius-filter-status-enum.md) | Optional | Status of the device search radius filter. enum: `up`, `down`, `unreachable` |
| `AuthTimeouts` | `*int` | Optional | Number of authentication timeouts |

## Example (as JSON)

```json
{
  "auth_accepts": 114,
  "auth_rejects": 84,
  "auth_server_status": "down",
  "auth_timeouts": 238
}
```

