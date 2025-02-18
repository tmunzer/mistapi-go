
# Wlan Mist Nac

*This model accepts additional fields of type interface{}.*

## Structure

`WlanMistNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | When enabled:<br><br>* `auth_servers` is ignored<br>* `acct_servers` is ignored<br>* `auth_servers_*` are ignored<br>* `coa_servers` is ignored<br>* `radsec` is ignored<br>* `coa_enabled` is assumed'<br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

