
# Tunnel Provider Options Zscaler

For zscaler-ipsec and zscaler-gre

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelProviderOptionsZscaler`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AupBlockInternetUntilAccepted` | `*bool` | Optional | **Default**: `false` |
| `AupEnabled` | `*bool` | Optional | Can only be `true` when `auth_required`==`false`, display Acceptable Use Policy (AUP)<br><br>**Default**: `false` |
| `AupForceSslInspection` | `*bool` | Optional | Proxy HTTPs traffic, requiring Zscaler cert to be installed in browser<br><br>**Default**: `false` |
| `AupTimeoutInDays` | `*int` | Optional | Required if `aup_enabled`==`true`. Days before AUP is requested again<br><br>**Constraints**: `>= 1`, `<= 180` |
| `AuthRequired` | `*bool` | Optional | Enable this option to enforce user authentication<br><br>**Default**: `false` |
| `CautionEnabled` | `*bool` | Optional | Can only be `true` when `auth_required`==`false`, display caution notification for non-authenticated users<br><br>**Default**: `false` |
| `DnBandwidth` | `models.Optional[float64]` | Optional | Download bandwidth cap of the link, in Mbps. Disabled if not set<br><br>**Constraints**: `>= 0.1`, `<= 99999` |
| `IdleTimeInMinutes` | `*int` | Optional | Required if `surrogate_IP`==`true`, idle Time to Disassociation<br><br>**Constraints**: `>= 0`, `<= 43200` |
| `OfwEnabled` | `*bool` | Optional | If `true`, enable the firewall control option<br><br>**Default**: `false` |
| `SubLocations` | [`[]models.TunnelProviderOptionsZscalerSubLocation`](../../doc/models/tunnel-provider-options-zscaler-sub-location.md) | Optional | `sub-locations` can be used for specific uses cases to define different configuration based on the user network |
| `SurrogateIP` | `*bool` | Optional | Can only be `true` when `auth_required`==`true`. Map a user to a private IP address so it applies the user's policies, instead of the location's policies<br><br>**Default**: `false` |
| `SurrogateIPEnforcedForKnownBrowsers` | `*bool` | Optional | Can only be `true` when `surrogate_IP`==`true`, enforce surrogate IP for known browsers |
| `SurrogateRefreshTimeInMinutes` | `*int` | Optional | Required if `surrogate_IP_enforced_for_known_browsers`==`true`, must be lower or equal than `idle_time_in_minutes`, refresh Time for re-validation of Surrogacy<br><br>**Constraints**: `>= 1`, `<= 43200` |
| `UpBandwidth` | `models.Optional[float64]` | Optional | Download bandwidth cap of the link, in Mbps. Disabled if not set<br><br>**Constraints**: `>= 0.1`, `<= 99999` |
| `XffForwardEnabled` | `*bool` | Optional | Location uses proxy chaining to forward traffic<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "aup_block_internet_until_accepted": false,
  "aup_enabled": false,
  "aup_force_ssl_inspection": false,
  "auth_required": false,
  "caution_enabled": false,
  "dn_bandwidth": 200,
  "ofw_enabled": false,
  "surrogate_IP": false,
  "up_bandwidth": 200,
  "xff_forward_enabled": false,
  "aup_timeout_in_days": 180,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

