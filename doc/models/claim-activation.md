
# Claim Activation

## Structure

`ClaimActivation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `string` | Required | activation code |
| `DeviceType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Optional | **Default**: `"ap"` |
| `Type` | [`models.ClaimTypeEnum`](../../doc/models/claim-type-enum.md) | Required | what to claim<br>**Default**: `"all"` |

## Example (as JSON)

```json
{
  "code": "code4",
  "device_type": "ap",
  "type": "all"
}
```

