
# Org Setting Marvis

## Structure

`OrgSettingMarvis`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SelfDriving` | [`*models.MarvisSelfDriving`](../../doc/models/marvis-self-driving.md) | Optional | Self-driving network automation settings per domain |

## Example (as JSON)

```json
{
  "self_driving": {
    "wan": {
      "enabled": false
    },
    "wired": {
      "enabled": false
    },
    "wireless": {
      "enabled": false
    }
  }
}
```

