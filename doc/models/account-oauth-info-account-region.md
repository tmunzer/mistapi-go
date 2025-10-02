
# Account Oauth Info Account Region

*This model accepts additional fields of type interface{}.*

## Structure

`AccountOauthInfoAccountRegion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AggregateRegion` | `*string` | Optional | Bandwidth Aggregate region for this region |
| `AllocatedBandwidth` | `*int` | Optional | Allocated bandwidth for the region, in Mbps |
| `Name` | `*string` | Optional | Display name for this region |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "aggregate_region": "us-southwest",
  "allocated_bandwidth": 1000,
  "name": "US West",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

