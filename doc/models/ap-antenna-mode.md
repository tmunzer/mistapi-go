
# Ap Antenna Mode

*This model accepts additional fields of type interface{}.*

## Structure

`ApAntennaMode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AntMode` | [`*models.AntModeEnum`](../../doc/models/ant-mode-enum.md) | Optional | Antenna Mode for AP which supports selectable antennas. enum: `external`, `internal` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ant_mode": "external",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

