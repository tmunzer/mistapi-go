
# Const Alarm Definition

## Structure

`ConstAlarmDefinition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `string` | Required | Description of the alarm type |
| `Example` | `*interface{}` | Optional | - |
| `Fields` | `[]string` | Required | List of fields available in an alarm details payload (in REST APIs & Webhooks); e.g. `aps`, `switches`, `gateways`, `hostnames`, `ssids`, `bssids` |
| `Group` | `string` | Required | Group to which the alarm belongs |
| `Key` | `string` | Required | Key name of the alarm type |
| `MarvisSuggestionCategory` | `*string` | Optional | Marvis defined category to which the alarm belongs |
| `Severity` | `string` | Required | Severity of the alarm |

## Example (as JSON)

```json
{
  "display": "Device offline",
  "example": {
    "aps": [
      "d420b02000fa"
    ],
    "count": 1,
    "group": "infrastructure",
    "hostnames": [
      "Vendor_AP2"
    ],
    "id": "f70c308f-7007-4866-9ecd-0d01842979ea",
    "last_seen": 1629753888,
    "org_id": "09dac91f-6e73-4100-89f7-698e0fafbb1b",
    "severity": "warn",
    "site_id": "dcfb31a1-d615-4361-8c95-b9dde05aa704",
    "timestamp": 1629753888,
    "type": "device_down"
  },
  "fields": [
    "aps",
    "hostnames"
  ],
  "group": "infrastructure",
  "key": "device_down",
  "severity": "warn",
  "marvis_suggestion_category": "marvis_suggestion_category4"
}
```

