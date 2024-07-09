
# Marvis Client

## Structure

`MarvisClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Diabled` | `*bool` | Optional | **Default**: `false` |
| `Id` | `*uuid.UUID` | Optional | - |
| `Name` | `*string` | Optional | - |
| `ProvisionUrl` | `*string` | Optional | in MDM, add `--provision_url <provision_url>` to the instlal command |

## Example (as JSON)

```json
{
  "diabled": false,
  "id": "3a14098f-b995-7552-b0a4-b8ee39b337a6",
  "name": "Handhelds",
  "provision_url": "https://api.mist.com/path/to/url"
}
```

