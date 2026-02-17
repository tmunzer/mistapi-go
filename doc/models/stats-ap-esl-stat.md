
# Stats Ap Esl Stat

## Structure

`StatsApEslStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `models.Optional[int]` | Optional | - |
| `Connected` | `models.Optional[bool]` | Optional | - |
| `Ip` | `models.Optional[string]` | Optional | IP address of Hanshow and SoluM dongles |
| `Mac` | `models.Optional[string]` | Optional | MAC address of Hanshow and SoluM dongles |
| `ProductId` | `models.Optional[string]` | Optional | Product ID of Hanshow and SoluM dongles |
| `Type` | `models.Optional[string]` | Optional | - |
| `Up` | `models.Optional[bool]` | Optional | - |
| `VendorId` | `models.Optional[string]` | Optional | Vendor ID of Hanshow and SoluM dongles |

## Example (as JSON)

```json
{
  "ip": "172.16.2.249",
  "mac": "98-6d-35-79-76-3b",
  "product_id": "A4A2",
  "type": "imagotag",
  "vendor_id": "0525",
  "channel": 76,
  "connected": false
}
```

