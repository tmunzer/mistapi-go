
# Org Service Policy

## Structure

`OrgServicePolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.AllowDenyEnum`](../../doc/models/allow-deny-enum.md) | Optional | **Default**: `"allow"` |
| `Appqoe` | [`*models.ServicePolicyAppqoe`](../../doc/models/service-policy-appqoe.md) | Optional | For SRX Only |
| `CreatedTime` | `*float64` | Optional | - |
| `Ewf` | [`[]models.ServicePolicyEwfRule`](../../doc/models/service-policy-ewf-rule.md) | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Idp` | [`*models.IdpConfig`](../../doc/models/idp-config.md) | Optional | - |
| `LocalRouting` | `*bool` | Optional | access within the same VRF<br>**Default**: `false` |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PathPreferences` | `*string` | Optional | by default, we derive all paths available and use them<br>optionally, you can customize by using `path_preference` |
| `Services` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Tenants` | `[]string` | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "action": "allow",
  "local_routing": false,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "appqoe": {
    "enabled": false
  },
  "created_time": 18.54,
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
  "id": "00000e28-0000-0000-0000-000000000000"
}
```

