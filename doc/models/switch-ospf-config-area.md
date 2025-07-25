
# Switch Ospf Config Area

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchOspfConfigArea`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NoSummary` | `*bool` | Optional | Disable OSPF summary routes for this area<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "no_summary": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

