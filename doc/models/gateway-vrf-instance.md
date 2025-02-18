
# Gateway Vrf Instance

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayVrfInstance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Networks` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "networks": [
    "CORP_NET",
    "MGMT_NET"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

