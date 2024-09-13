
# Secintel Profile

## Structure

`SecintelProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `Profiles` | [`[]models.SecintelProfileProfile`](../../doc/models/secintel-profile-profile.md) | Optional | - |

## Example (as JSON)

```json
{
  "name": "secintel-custom",
  "profiles": [
    {
      "action": "strict",
      "category": "DNS"
    },
    {
      "action": "strict",
      "category": "DNS"
    }
  ]
}
```

