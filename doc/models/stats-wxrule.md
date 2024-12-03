
# Stats Wxrule

Wxrule statistics

*This model accepts additional fields of type interface{}.*

## Structure

`StatsWxrule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`models.StatsWxruleActionEnum`](../../doc/models/stats-wxrule-action-enum.md) | Required | enum: `allow`, `block` |
| `ClientMac` | `[]string` | Required | - |
| `DstAllowWxtags` | `[]uuid.UUID` | Required | - |
| `DstDenyWxtags` | `[]uuid.UUID` | Required | - |
| `DstWxtags` | `[]uuid.UUID` | Required | - |
| `Name` | `string` | Required | - |
| `Order` | `int` | Required | - |
| `SrcWxtags` | `[]uuid.UUID` | Required | - |
| `Usage` | [`map[string]models.StatsWxruleUsageProperties`](../../doc/models/stats-wxrule-usage-properties.md) | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "action": "allow",
  "client_mac": [
    "3bbbf819bb6f",
    "bd96cbc4910f"
  ],
  "dst_allow_wxtags": [
    "fff34466-eec0-3756-6765-381c728a6037",
    "eee2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
  ],
  "dst_deny_wxtags": [
    "aaa34466-eec0-3756-6765-381c728a6037",
    "bbb2c7b0-d1d0-5a30-f349-e35fa43dc3b3"
  ],
  "dst_wxtags": [
    "d4134466-eec0-3756-6765-381c728a6037",
    "1a42c7b0-d1d0-5a30-f349-e35fa43dc3b3"
  ],
  "name": "Guest",
  "order": 1,
  "src_wxtags": [
    "8bfc2490-d726-3587-038d-cb2e71bd2330",
    "3aa8e73f-9f46-d827-8d6a-567bb7e67fc9"
  ],
  "usage": {
    "1a42c7b0-d1d0-5a30-f349-e35fa43dc3b3": {
      "num_flows": 60,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "d4134466-eec0-3756-6765-381c728a6037": {
      "num_flows": 60,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

