
# Idp Profile Matching

*This model accepts additional fields of type interface{}.*

## Structure

`IdpProfileMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AttackName` | `[]string` | Optional | - |
| `DstSubnet` | `[]string` | Optional | - |
| `Severity` | [`[]models.IdpProfileMatchingSeverityValueEnum`](../../doc/models/idp-profile-matching-severity-value-enum.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "attack_name": [
    "attack_name7",
    "attack_name8"
  ],
  "dst_subnet": [
    "dst_subnet9",
    "dst_subnet0"
  ],
  "severity": [
    "major"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

