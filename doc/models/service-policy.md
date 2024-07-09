
# Service Policy

## Structure

`ServicePolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.AllowDenyEnum`](../../doc/models/allow-deny-enum.md) | Optional | **Default**: `"allow"` |
| `Appqoe` | [`*models.ServicePolicyAppqoe`](../../doc/models/service-policy-appqoe.md) | Optional | For SRX Only |
| `Ewf` | [`[]models.ServicePolicyEwfRule`](../../doc/models/service-policy-ewf-rule.md) | Optional | - |
| `Idp` | [`*models.IdpConfig`](../../doc/models/idp-config.md) | Optional | - |
| `LocalRouting` | `*bool` | Optional | access within the same VRF<br>**Default**: `false` |
| `Name` | `*string` | Optional | - |
| `PathPreferences` | `*string` | Optional | by default, we derive all paths available and use them<br>optionally, you can customize by using `path_preference` |
| `ServicepolicyId` | `*uuid.UUID` | Optional | used to link servicepolicy defined at org level and overwrite some attributes |
| `Services` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Tenants` | `[]string` | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "action": "allow",
  "local_routing": false,
  "appqoe": {
    "enabled": false
  },
  "ewf": [
    {
      "alert_only": false,
      "block_message": "block_message0",
      "enabled": false,
      "profile": "strict"
    },
    {
      "alert_only": false,
      "block_message": "block_message0",
      "enabled": false,
      "profile": "strict"
    },
    {
      "alert_only": false,
      "block_message": "block_message0",
      "enabled": false,
      "profile": "strict"
    }
  ],
  "idp": {
    "alert_only": false,
    "enabled": false,
    "idpprofile_id": "00000e94-0000-0000-0000-000000000000",
    "profile": "profile8"
  }
}
```

