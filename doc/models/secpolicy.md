
# Secpolicy

Security Policy is designed to audit / catch discripancies between “what’s intended to be running” versus “what’s actually running” in a network. Many big organizations have separated Security and IT team (for good reasons). Each site can be assigned a security policy. Whenever an AP is provisioned, the configuration will be checked against the security policy. Any violations will be flagged in Device Config History where you can search for the when and where the violation occurs.

*This model accepts additional fields of type interface{}.*

## Structure

`Secpolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Wlans` | [`[]models.Wlan`](../../doc/models/wlan.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "created_time": 67.52,
  "modified_time": 11.44,
  "name": "name2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

