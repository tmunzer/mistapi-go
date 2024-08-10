
# Response Config History Search Item Wlan

## Structure

`ResponseConfigHistorySearchItemWlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Auth` | `string` | Required | - |
| `Bands` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Id` | `uuid.UUID` | Required | - |
| `Ssid` | `string` | Required | - |
| `VlanIds` | `[]string` | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "auth": "auth0",
  "bands": [
    "bands4"
  ],
  "id": "00000130-0000-0000-0000-000000000000",
  "ssid": "ssid8",
  "vlan_ids": [
    "vlan_ids6"
  ]
}
```

