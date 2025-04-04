
# User Mac

*This model accepts additional fields of type interface{}.*

## Structure

`UserMac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Labels` | `[]string` | Optional | - |
| `Mac` | `string` | Required | Only non-local-admin MAC is accepted |
| `Name` | `*string` | Optional | - |
| `Notes` | `*string` | Optional | - |
| `RadiusGroup` | `*string` | Optional | - |
| `Vlan` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  "vlan": "30",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

