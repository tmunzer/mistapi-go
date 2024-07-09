
# Alarm Search Result

## Structure

`AlarmSearchResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `int` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Page` | `*int` | Optional | - |
| `Results` | [`[]models.Alarm`](../../doc/models/alarm.md) | Required | - |
| `Start` | `int` | Required | - |
| `Total` | `int` | Required | - |

## Example (as JSON)

```json
{
  "end": 1711035686,
  "limit": 10,
  "next": "/api/v1/orgs/b3b9f5e6-67b1-4112-9b4c-6824c565eaeb/alarms/search?end=1711035686&limit=10&search_after=%5B1711031354000%2C+%2256bfa7af-b2db-43ee-a4c8-9b820bbba0e1%22%5D&start=1710949286",
  "page": 1,
  "results": [
    {
      "ack_admin_id": "6603c94c-eaf9-4014-9edf-b9f8eed6b183",
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
      "id": "483f6eca-6276-4993-bfeb-56cbbbba7f08",
      "last_seen": 1711031774,
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "severity": "critical",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "switches": [
        "ffeeddccbbaa",
        "ffeeddccbbab"
      ],
      "timestamp": 1711031774,
      "type": "rogue_client"
    }
  ],
  "start": 1710949286,
  "total": 232
}
```
