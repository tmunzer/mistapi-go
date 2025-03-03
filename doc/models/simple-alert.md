
# Simple Alert

Set of heuristic rules will be enabled when marvis subscription is not available. It triggers when, in a Z minute window, there are more than Y distinct client encountering over X failures

*This model accepts additional fields of type interface{}.*

## Structure

`SimpleAlert`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ArpFailure` | [`*models.SimpleAlertArpFailure`](../../doc/models/simple-alert-arp-failure.md) | Optional | - |
| `DhcpFailure` | [`*models.SimpleAlertDhcpFailure`](../../doc/models/simple-alert-dhcp-failure.md) | Optional | - |
| `DnsFailure` | [`*models.SimpleAlertDnsFailure`](../../doc/models/simple-alert-dns-failure.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "arp_failure": {
    "client_count": 26,
    "duration": 130,
    "incident_count": 226,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "dhcp_failure": {
    "client_count": 246,
    "duration": 94,
    "incident_count": 6,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "dns_failure": {
    "client_count": 252,
    "duration": 100,
    "incident_count": 0,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

