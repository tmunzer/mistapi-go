
# Msp

*This model accepts additional fields of type interface{}.*

## Structure

`Msp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowMist` | `*bool` | Optional | - |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `LogoUrl` | `*string` | Optional | For advanced tier (uMSPs) only |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | - |
| `Tier` | [`*models.MspTierEnum`](../../doc/models/msp-tier-enum.md) | Optional | enum: `advanced`, `base`<br>**Default**: `"base"` |
| `Url` | `*string` | Optional | For advanced tier (uMSPs) only |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "tier": "base",
  "allow_mist": false,
  "created_time": 109.3,
  "logo_url": "logo_url0",
  "modified_time": 225.66,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

