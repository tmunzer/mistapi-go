
# Utils Rrm Optimize

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsRrmOptimize`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bands` | `[]string` | Required | List of bands |
| `Macs` | `[]string` | Optional | Targeting AP (neighbor APs may get changed, too), default is empty for ALL APs |
| `TxpowerOnly` | `*bool` | Optional | Only changing TX Power (will not disconnect clients)<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "bands": [
    "bands2"
  ],
  "txpower_only": false,
  "macs": [
    "macs7",
    "macs6",
    "macs5"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

