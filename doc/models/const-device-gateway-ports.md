
# Const Device Gateway Ports

Object Key is the interface name (e.g. "ge-0/0/1", ...)

*This model accepts additional fields of type interface{}.*

## Structure

`ConstDeviceGatewayPorts`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `*string` | Optional | - |
| `PciAddress` | `*string` | Optional | - |
| `Speed` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "display": "display8",
  "pci_address": "pci_address8",
  "speed": 146,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

