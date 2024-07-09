
# User Mac

## Structure

`UserMac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | - |
| `Labels` | `[]string` | Optional | - |
| `Mac` | `*string` | Optional | only non-local-admin MAC is accepted |
| `Notes` | `*string` | Optional | - |
| `Vlan` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "labels": [
    "byod",
    "flr1"
  ],
  "mac": "5684dae9ac8b",
  "notes": "mac address refers to Canon printers",
  "vlan": 30,
  "id": "00000968-0000-0000-0000-000000000000"
}
```

