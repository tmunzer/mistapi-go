
# Marvis Self Driving

Self-driving network automation settings per domain

## Structure

`MarvisSelfDriving`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Wan` | [`*models.MarvisSelfDrivingDomain`](../../doc/models/marvis-self-driving-domain.md) | Optional | - |
| `Wired` | [`*models.MarvisSelfDrivingDomain`](../../doc/models/marvis-self-driving-domain.md) | Optional | - |
| `Wireless` | [`*models.MarvisSelfDrivingDomain`](../../doc/models/marvis-self-driving-domain.md) | Optional | - |

## Example (as JSON)

```json
{
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
```

