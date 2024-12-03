
# Org Setting Security

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingSecurity`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DisableLocalSsh` | `*bool` | Optional | whether to disable local SSH (by default, local SSH is enabled with allow_mist in Org is enabled |
| `FipsZeroizePassword` | `*string` | Optional | password required to zeroize devices (FIPS) on site level |
| `LimitSshAccess` | `*bool` | Optional | whether to allow certain SSH keys to SSH into the AP (see Site:Setting)<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "fips_zeroize_password": "NUKETHESITE",
  "limit_ssh_access": false,
  "disable_local_ssh": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

