
# Synthetictest Properties

*This model accepts additional fields of type interface{}.*

## Structure

`SynthetictestProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomTestUrls` | `[]string` | Optional | - |
| `Disabled` | `*bool` | Optional | For some vlans where we don't want this to run<br>**Default**: `false` |
| `VlanIds` | [`[]models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "custom_test_urls": [
    "http://www.abc.com/",
    "https://10.3.5.1:8080/about"
  ],
  "disabled": false,
  "vlan_ids": [
    10,
    20,
    "{{vlan}}"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

