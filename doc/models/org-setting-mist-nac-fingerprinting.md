
# Org Setting Mist Nac Fingerprinting

Allows customer to enable client fingerprinting for policy enforcement

## Structure

`OrgSettingMistNacFingerprinting`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | enable/disable writes to NAC DDB fingerprint table<br><br>**Default**: `false` |
| `GenerateCoa` | `*bool` | Optional | enable/disable CoA triggers on fingerprint change for wired clients, always port-bounce<br><br>**Default**: `false` |
| `GenerateWirelessCoa` | `*bool` | Optional | enable/disable CoA triggers on fingerprint change for wireless clients<br><br>**Default**: `false` |
| `WirelessCoaType` | [`*models.OrgSettingMistNacFingerprintingWirelessCoaEnum`](../../doc/models/org-setting-mist-nac-fingerprinting-wireless-coa-enum.md) | Optional | enum: `reauth`, `disconnect` |

## Example (as JSON)

```json
{
  "enabled": false,
  "generate_coa": false,
  "generate_wireless_coa": false,
  "wireless_coa_type": "reauth"
}
```

