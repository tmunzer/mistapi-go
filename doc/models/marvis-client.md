
# Marvis Client

*This model accepts additional fields of type interface{}.*

## Structure

`MarvisClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | **Default**: `false` |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `Name` | `*string` | Optional | - |
| `ProvisionUrl` | `*string` | Optional | In MDM, add `--provision_url <provision_url>` to the install command |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "disabled": false,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "Handhelds",
  "provision_url": "https://api.mist.com/path/to/url",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

