
# Response Troubleshoot Item

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseTroubleshootItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Category` | `*string` | Optional | - |
| `Reason` | `*string` | Optional | - |
| `Recommendation` | `*string` | Optional | - |
| `Text` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "category": "client",
  "reason": "slow association",
  "recommendation": "Ensure the IP helper-address is configured on the VLAN interface.",
  "text": "Clients of the AP had slow association 8% of the time on Bhavabhi and 5 GHz. ...",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

