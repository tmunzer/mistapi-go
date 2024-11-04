
# Marvis Client

## Structure

`MarvisClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Diabled` | `*bool` | Optional | **Default**: `false` |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `Name` | `*string` | Optional | - |
| `ProvisionUrl` | `*string` | Optional | in MDM, add `--provision_url <provision_url>` to the instlal command |

## Example (as JSON)

```json
{
  "diabled": false,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "name": "Handhelds",
  "provision_url": "https://api.mist.com/path/to/url"
}
```

