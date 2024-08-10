
# User Mac

## Structure

`UserMac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | - |
| `Labels` | `[]string` | Optional | - |
| `Mac` | `*string` | Optional | only non-local-admin MAC is accepted |
| `Name` | `*string` | Optional | - |
| `Notes` | `*string` | Optional | - |
| `RadiusGroup` | `*string` | Optional | - |
| `Vlan` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "labels": [
    "byod",
    "flr1"
  ],
  "mac": "5684dae9ac8b",
  "name": "Printer2",
  "notes": "mac address refers to Canon printers",
  "radius_group": "VIP",
  "vlan": 30,
  "id": "000003ce-0000-0000-0000-000000000000"
}
```

