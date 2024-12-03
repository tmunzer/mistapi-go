
# Ap Iot

IoT AP settings

*This model accepts additional fields of type interface{}.*

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
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "A1": {
    "enabled": false,
    "name": "name0",
    "output": false,
    "pullup": "internal",
    "value": 242,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "A2": {
    "enabled": false,
    "name": "name8",
    "output": false,
    "pullup": "none",
    "value": 180,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "A3": {
    "enabled": false,
    "name": "name6",
    "output": false,
    "pullup": "internal",
    "value": 118,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "A4": {
    "enabled": false,
    "name": "name8",
    "output": false,
    "pullup": "internal",
    "value": 88,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "DI1": {
    "enabled": false,
    "name": "name0",
    "pullup": "internal",
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

