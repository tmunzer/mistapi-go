
# Sdkstats Wireless Client Zone

*This model accepts additional fields of type interface{}.*

## Structure

`SdkstatsWirelessClientZone`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `Since` | `float64` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "since": 181.4,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

