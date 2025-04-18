
# Service Policy

*This model accepts additional fields of type interface{}.*

## Structure

`ServicePolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.AllowDenyEnum`](../../doc/models/allow-deny-enum.md) | Optional | enum: `allow`, `deny` |
| `Antivirus` | [`*models.ServicePolicyAntivirus`](../../doc/models/service-policy-antivirus.md) | Optional | For SRX-only |
| `Appqoe` | [`*models.ServicePolicyAppqoe`](../../doc/models/service-policy-appqoe.md) | Optional | For SRX Only |
| `Ewf` | [`[]models.ServicePolicyEwfRule`](../../doc/models/service-policy-ewf-rule.md) | Optional | - |
| `Idp` | [`*models.IdpConfig`](../../doc/models/idp-config.md) | Optional | - |
| `LocalRouting` | `*bool` | Optional | access within the same VRF |
| `Name` | `*string` | Optional | - |
| `PathPreference` | `*string` | Optional | By default, we derive all paths available and use them. Optionally, you can customize by using `path_preference` |
| `Secintel` | [`*models.ServicePolicySecintel`](../../doc/models/service-policy-secintel.md) | Optional | For SRX Only |
| `ServicepolicyId` | `*uuid.UUID` | Optional | Used to link servicepolicy defined at org level and overwrite some attributes |
| `Services` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `SslProxy` | [`*models.ServicePolicySslProxy`](../../doc/models/service-policy-ssl-proxy.md) | Optional | For SRX-only |
| `Tenants` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
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
  "ewf": [
    {
      "alert_only": false,
      "block_message": "block_message0",
      "enabled": false,
      "profile": "standard",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "alert_only": false,
      "block_message": "block_message0",
      "enabled": false,
      "profile": "standard",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "alert_only": false,
      "block_message": "block_message0",
      "enabled": false,
      "profile": "standard",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "idp": {
    "alert_only": false,
    "enabled": false,
    "idpprofile_id": "00000e94-0000-0000-0000-000000000000",
    "profile": "profile8",
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

