
# Org Service Policy

*This model accepts additional fields of type interface{}.*

## Structure

`OrgServicePolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aamw` | [`*models.ServicePolicyAamw`](../../doc/models/service-policy-aamw.md) | Optional | For SRX Only |
| `Action` | [`*models.AllowDenyEnum`](../../doc/models/allow-deny-enum.md) | Optional | enum: `allow`, `deny` |
| `Antivirus` | [`*models.ServicePolicyAntivirus`](../../doc/models/service-policy-antivirus.md) | Optional | For SRX-only |
| `Appqoe` | [`*models.ServicePolicyAppqoe`](../../doc/models/service-policy-appqoe.md) | Optional | For SRX Only |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Ewf` | [`[]models.ServicePolicyEwfRule`](../../doc/models/service-policy-ewf-rule.md) | Optional | - |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Idp` | [`*models.IdpConfig`](../../doc/models/idp-config.md) | Optional | - |
| `LocalRouting` | `*bool` | Optional | access within the same VRF |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PathPreference` | `*string` | Optional | By default, we derive all paths available and use them, optionally, you can customize by using `path_preference` |
| `Secintel` | [`*models.ServicePolicySecintel`](../../doc/models/service-policy-secintel.md) | Optional | For SRX Only |
| `Services` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `SslProxy` | [`*models.ServicePolicySslProxy`](../../doc/models/service-policy-ssl-proxy.md) | Optional | For SRX-only |
| `Tenants` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "aamw": {
    "aamwprofile_id": "0000066a-0000-0000-0000-000000000000",
    "enabled": false,
    "profile": "executables",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "action": "allow",
  "antivirus": {
    "avprofile_id": "00000282-0000-0000-0000-000000000000",
    "enabled": false,
    "profile": "profile4",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "appqoe": {
    "enabled": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "created_time": 21.72,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

