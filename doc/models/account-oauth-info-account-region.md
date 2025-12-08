
# Account Oauth Info Account Region

## Structure

`AccountOauthInfoAccountRegion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AggregateRegion` | `*string` | Optional | Bandwidth Aggregate region for this region |
| `AllocatedBandwidth` | `*int` | Optional | Allocated bandwidth for the region, in Mbps |
| `Name` | `*string` | Optional | Display name for this region |

## Example (as JSON)

```json
{
  "aggregate_region": "us-southwest",
  "allocated_bandwidth": 1000,
  "name": "US West"
}
```

