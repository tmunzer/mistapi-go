
# Const Alarm Definition

Alarm type definition returned by the constants API

## Structure

`ConstAlarmDefinition`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `string` | Required | Description of the alarm type |
| `Example` | `*interface{}` | Optional | Sample alarm payload returned for this alarm type |
| `Fields` | `[]string` | Required | List of fields available in an alarm details payload (in REST APIs & Webhooks); e.g. `aps`, `switches`, `gateways`, `hostnames`, `ssids`, `bssids` |
| `Group` | `string` | Required | Alarm group to which this definition belongs |
| `Key` | `string` | Required | Alarm type key used in event and alarm payloads |
| `MarvisSuggestionCategory` | `*string` | Optional | Marvis defined category to which the alarm belongs |
| `Severity` | `string` | Required | Alarm severity level for this definition |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constAlarmDefinition := models.ConstAlarmDefinition{
        Display:                  "Device offline",
        Example:                  models.ToPointer(interface{}("[aps, System.Object[]][count, 1][group, infrastructure][hostnames, System.Object[]][id, f70c308f-7007-4866-9ecd-0d01842979ea][last_seen, 1629753888][org_id, 09dac91f-6e73-4100-89f7-698e0fafbb1b][severity, warn][site_id, dcfb31a1-d615-4361-8c95-b9dde05aa704][timestamp, 1629753888][type, device_down]")),
        Fields:                   []string{
            "aps",
            "hostnames",
        },
        Group:                    "infrastructure",
        Key:                      "device_down",
        MarvisSuggestionCategory: models.ToPointer("marvis_suggestion_category0"),
        Severity:                 "warn",
    }

}
```

