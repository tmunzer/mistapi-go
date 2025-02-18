
# Alarm

Additional information per alarm type

*This model accepts additional fields of type interface{}.*

## Structure

`Alarm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AckAdminId` | `*uuid.UUID` | Optional | UUID of the admin who acked the alarm |
| `AckAdminName` | `*string` | Optional | Name & Email ID of the admin who acked the alarm |
| `Acked` | `*bool` | Optional | Whether the alarm is acked or not |
| `AckedTime` | `*int` | Optional | Epoch (seconds) when the alarm was acked |
| `Aps` | `[]string` | Optional | additional information: List of MACs of the APs |
| `Bssids` | `[]string` | Optional | List of BSSIDs |
| `Count` | `int` | Required | Number of incident within an alarm window |
| `Gateways` | `[]string` | Optional | additional information: List of MACs of the gateways |
| `Group` | `string` | Required | Group of the alarm |
| `Hostnames` | `[]string` | Optional | additional information: List of Hostnames of the devices (AP/Switch/Gateway) |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organnization |
| `LastSeen` | `float64` | Required | Epoch (seconds) of the last incident/alarm within an alarm window |
| `Note` | `*string` | Optional | Text describing the alarm |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Severity` | `string` | Required | Severity of the alarm |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssids` | `[]string` | Optional | List of SSIDs |
| `Switches` | `[]string` | Optional | additional information: List of MACs of the switches |
| `Timestamp` | `int` | Required | Epoch (seconds) of the first incident/alarm |
| `Type` | `string` | Required | Key-name of the alarm type |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ack_admin_id": "456b7016-a916-a4b1-78dd-72b947c152b7",
  "ack_admin_name": "Joe",
  "acked": true,
  "acked_time": 1711031352,
  "aps": [
    "ffeeddccbbaa",
    "ffeeddccbbab"
  ],
  "count": 2,
  "gateways": [
    "ffeeddccbbaa",
    "ffeeddccbbab"
  ],
  "group": "security",
  "hostnames": [
    "MC_DavidL",
    "MCM_AP_33_Nishant"
  ],
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "last_seen": 1711031774.0,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "severity": "critical",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "switches": [
    "ffeeddccbbaa",
    "ffeeddccbbab"
  ],
  "timestamp": 1711031774,
  "type": "rogue_client",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

