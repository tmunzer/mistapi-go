
# Marvis Client

*This model accepts additional fields of type interface{}.*

## Structure

`MarvisClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | **Default**: `false` |
| `EnrollmentUrl` | `*string` | Optional | In MDM, add `--enrollment_url <enrollment_url>` to the install command |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Location` | [`*models.MarvisClientLocation`](../../doc/models/marvis-client-location.md) | Optional | - |
| `Name` | `*string` | Optional | - |
| `SyntheticTest` | [`*models.MarvisClientSyntheticTest`](../../doc/models/marvis-client-synthetic-test.md) | Optional | - |
| `Telemetry` | [`*models.MarvisClientTelemetry`](../../doc/models/marvis-client-telemetry.md) | Optional | Note: some stats are not collected when it's not connected to Mist infrastructure |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disabled": false,
  "enrollment_url": "marvisclient://api.mist.com/path/to/url",
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "Handhelds",
  "location": {
    "enabled": false
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

