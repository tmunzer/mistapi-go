
# Ap Iot

IoT AP settings

## Structure

`ApIot`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `A1` | [`*models.ApIotOutput`](../../doc/models/ap-iot-output.md) | Optional | IoT output AP settings |
| `A2` | [`*models.ApIotOutput`](../../doc/models/ap-iot-output.md) | Optional | IoT output AP settings |
| `A3` | [`*models.ApIotOutput`](../../doc/models/ap-iot-output.md) | Optional | IoT output AP settings |
| `A4` | [`*models.ApIotOutput`](../../doc/models/ap-iot-output.md) | Optional | IoT output AP settings |
| `DI1` | [`*models.ApIotInput`](../../doc/models/ap-iot-input.md) | Optional | IoT Input AP settings |
| `DI2` | [`*models.ApIotInput`](../../doc/models/ap-iot-input.md) | Optional | IoT Input AP settings |
| `DO` | [`*models.ApIotOutput`](../../doc/models/ap-iot-output.md) | Optional | IoT output AP settings |

## Example (as JSON)

```json
{
  "A1": {
    "enabled": false,
    "name": "name0",
    "output": false,
    "pullup": "internal",
    "value": 242
  },
  "A2": {
    "enabled": false,
    "name": "name8",
    "output": false,
    "pullup": "none",
    "value": 180
  },
  "A3": {
    "enabled": false,
    "name": "name6",
    "output": false,
    "pullup": "internal",
    "value": 118
  },
  "A4": {
    "enabled": false,
    "name": "name8",
    "output": false,
    "pullup": "internal",
    "value": 88
  },
  "DI1": {
    "enabled": false,
    "name": "name0",
    "pullup": "internal"
  }
}
```

