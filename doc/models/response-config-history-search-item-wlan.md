
# Response Config History Search Item Wlan

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseConfigHistorySearchItemWlan`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Auth` | `string` | Required | - |
| `Bands` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `Ssid` | `string` | Required | - |
| `VlanIds` | `[]string` | Optional | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "auth": "auth0",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "ssid": "ssid8",
  "bands": [
    "bands4"
  ],
  "vlan_ids": [
    "vlan_ids6"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

