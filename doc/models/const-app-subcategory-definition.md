
# Const App Subcategory Definition

*This model accepts additional fields of type interface{}.*

## Structure

`ConstAppSubcategoryDefinition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `string` | Required | Description of the app subcategory |
| `Key` | `string` | Required | Key name of the app subcategory |
| `TrafficType` | `string` | Required | Type of traffic (QoS) of the app subcategory |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "display": "Office Document",
  "key": "Office_Documents",
  "traffic_type": "Images",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

