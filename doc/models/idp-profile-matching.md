
# Idp Profile Matching

## Structure

`IdpProfileMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AttackName` | `[]string` | Optional | - |
| `DstSubnet` | `[]string` | Optional | - |
| `Severity` | [`[]models.IdpProfileMatchingSeverityValueEnum`](../../doc/models/idp-profile-matching-severity-value-enum.md) | Optional | - |

## Example (as JSON)

```json
{
  "attack_name": [
    "attack_name5"
  ],
  "dst_subnet": [
    "dst_subnet1"
  ],
  "severity": [
    "critical",
    "info"
  ]
}
```

