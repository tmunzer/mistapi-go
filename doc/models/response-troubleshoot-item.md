
# Response Troubleshoot Item

## Structure

`ResponseTroubleshootItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Category` | `*string` | Optional | - |
| `Reason` | `*string` | Optional | - |
| `Recommendation` | `*string` | Optional | - |
| `Text` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "category": "client",
  "reason": "slow association",
  "recommendation": "Ensure the IP helper-address is configured on the VLAN interface.",
  "text": "Clients of the AP had slow association 8% of the time on Bhavabhi and 5 GHz. ..."
}
```

