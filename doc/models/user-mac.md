
# User Mac

## Structure

`UserMac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Labels` | `[]string` | Optional | - |
| `Mac` | `string` | Required | only non-local-admin MAC is accepted |
| `Name` | `*string` | Optional | - |
| `Notes` | `*string` | Optional | - |
| `RadiusGroup` | `*string` | Optional | - |
| `Vlan` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "labels": [
    "byod",
    "flr1"
  ],
  "mac": "5684dae9ac8b",
  "name": "Printer2",
  "notes": "mac address refers to Canon printers",
  "radius_group": "VIP",
  "vlan": "30"
}
```

