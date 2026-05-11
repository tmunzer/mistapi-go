
# Installer Device Ble Stat

BLE statistics for the device

## Structure

`InstallerDeviceBleStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Major` | `*int` | Optional | - |
| `Minors` | `[]int` | Optional | - |
| `Uuid` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "major": 12345,
  "uuid": "ada72f8f-1643-e5c6-94db-f2a5636f1a64",
  "minors": [
    169,
    170,
    171
  ]
}
```

