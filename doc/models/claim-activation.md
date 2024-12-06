
# Claim Activation

*This model accepts additional fields of type interface{}.*

## Structure

`ClaimActivation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Async` | `*bool` | Optional | whether to do a async claim process<br>**Default**: `false` |
| `Code` | `string` | Required | activation code |
| `DeviceType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Optional | enum: `ap`, `gateway`, `switch`<br>**Default**: `"ap"` |
| `Type` | [`models.ClaimTypeEnum`](../../doc/models/claim-type-enum.md) | Required | what to claim. enum: `all`, `inventory`, `license`<br>**Default**: `"all"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "async": false,
  "code": "code8",
  "device_type": "ap",
  "type": "all",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

