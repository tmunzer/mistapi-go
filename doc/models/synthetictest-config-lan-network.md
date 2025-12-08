
# Synthetictest Config Lan Network

configure minis probes to be tested on lan networks of gateways

## Structure

`SynthetictestConfigLanNetwork`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Networks` | `[]string` | Optional | List of networks to be used for synthetic tests |
| `Probes` | `[]string` | Optional | app name comes from `custom_probes` above or /const/synthetic_test_probes |

## Example (as JSON)

```json
{
  "networks": [
    "pos-stations",
    "pos-machines"
  ],
  "probes": [
    "probes7"
  ]
}
```

