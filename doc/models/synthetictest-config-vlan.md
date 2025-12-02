
# Synthetictest Config Vlan

## Structure

`SynthetictestConfigVlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomTestUrls` | `[]string` | Optional | - |
| `Disabled` | `*bool` | Optional | For some vlans where we don't want this to run<br><br>**Default**: `false` |
| `Probes` | `[]string` | Optional | app name comes from `custom_probes` above or /const/synthetic_test_probes |
| `VlanIds` | [`[]models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | - |

## Example (as JSON)

```json
{
  "custom_test_urls": [
    "https://www.abc.com/",
    "https://10.3.5.1:8080/about"
  ],
  "disabled": false,
  "vlan_ids": [
    10,
    20,
    "{{vlan}}"
  ],
  "probes": [
    "probes9",
    "probes8"
  ]
}
```

