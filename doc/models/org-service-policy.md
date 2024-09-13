
# Org Service Policy

## Structure

`OrgServicePolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.AllowDenyEnum`](../../doc/models/allow-deny-enum.md) | Optional | enum: `allow`, `deny`<br>**Default**: `"allow"` |
| `Antivirus` | [`*models.OrgServicePolicyAntivirus`](../../doc/models/org-service-policy-antivirus.md) | Optional | for SRX-only |
| `Appqoe` | [`*models.ServicePolicyAppqoe`](../../doc/models/service-policy-appqoe.md) | Optional | For SRX Only |
| `CreatedTime` | `*float64` | Optional | - |
| `Ewf` | [`[]models.ServicePolicyEwfRule`](../../doc/models/service-policy-ewf-rule.md) | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Idp` | [`*models.IdpConfig`](../../doc/models/idp-config.md) | Optional | - |
| `LocalRouting` | `*bool` | Optional | access within the same VRF<br>**Default**: `false` |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PathPreference` | `*string` | Optional | by default, we derive all paths available and use them<br>optionally, you can customize by using `path_preference` |
| `Secintel` | [`*models.OrgServicePoliciesSecintel`](../../doc/models/org-service-policies-secintel.md) | Optional | For SRX Only |
| `Services` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `SslProxy` | [`*models.OrgServicePolicySslProxy`](../../doc/models/org-service-policy-ssl-proxy.md) | Optional | for SRX-only |
| `Tenants` | `[]string` | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "action": "allow",
  "local_routing": false,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "antivirus": {
    "avprofile_id": "00000282-0000-0000-0000-000000000000",
    "enabled": false,
    "profile": "profile4"
  },
  "appqoe": {
    "enabled": false
  },
  "created_time": 21.72,
  "ewf": [
    {
      "alert_only": false,
      "block_message": "block_message0",
      "enabled": false,
      "profile": "standard"
    },
    {
      "alert_only": false,
      "block_message": "block_message0",
      "enabled": false,
      "profile": "standard"
    },
    {
      "alert_only": false,
      "block_message": "block_message0",
      "enabled": false,
      "profile": "standard"
    }
  ]
}
```

