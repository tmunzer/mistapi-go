
# Wlan Mist Nac

## Structure

`WlanMistNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | when enabled:<br><br>* `auth_servers` is ignored<br>* `acct_servers` is ignored<br>* `auth_servers_*` are ignored<br>* `coa_servers` is ignored<br>* `radsec` is ignored<br>* `coa_enabled` is assumed'<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "enabled": false
}
```

