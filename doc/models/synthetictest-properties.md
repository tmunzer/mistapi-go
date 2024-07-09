
# Synthetictest Properties

## Structure

`SynthetictestProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomTestUrls` | `[]string` | Optional | - |
| `Disabled` | `*bool` | Optional | for some vlans where we don't want this to run<br>**Default**: `false` |
| `VlanIds` | `[]int` | Optional | - |

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
    20
  ]
}
```

