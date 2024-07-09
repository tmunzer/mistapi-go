
# Utils Rrm Optimize

## Structure

`UtilsRrmOptimize`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bands` | `[]string` | Required | list of bands |
| `Macs` | `[]string` | Optional | targeting AP (neighbor APs may get changed, too), default is empty for ALL APs |
| `TxpowerOnly` | `*bool` | Optional | only changng TX Power (will not disconnect clients)<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "bands": [
    "bands4",
    "bands5",
    "bands6"
  ],
  "txpower_only": false,
  "macs": [
    "macs9"
  ]
}
```

