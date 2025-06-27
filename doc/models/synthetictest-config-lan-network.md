
# Synthetictest Config Lan Network

configure minis probes to be tested on lan networks of gateways

*This model accepts additional fields of type interface{}.*

## Structure

`SynthetictestConfigLanNetwork`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Networks` | `[]string` | Optional | List of networks to be used for synthetic tests |
| `Probes` | `[]string` | Optional | app name comes from `custom_probes` above or /const/synthetic_test_probes |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "networks": [
    "pos-stations",
    "pos-machines"
  ],
  "probes": [
    "probes7"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

