
# Synthetictest Config

## Structure

`SynthetictestConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Aggressiveness` | [`*models.SynthetictestConfigAggressivenessEnum`](../../doc/models/synthetictest-config-aggressiveness-enum.md) | Optional | enum: `auto`, `high`, `low`<br><br>**Default**: `"auto"` |
| `CustomProbes` | [`map[string]models.SynthetictestConfigCustomProbe`](../../doc/models/synthetictest-config-custom-probe.md) | Optional | Custom probes to be used for synthetic tests |
| `Disabled` | `*bool` | Optional | **Default**: `false` |
| `LanNetworks` | [`[]models.SynthetictestConfigLanNetwork`](../../doc/models/synthetictest-config-lan-network.md) | Optional | List of networks to be used for synthetic tests |
| `Vlans` | [`[]models.SynthetictestConfigVlan`](../../doc/models/synthetictest-config-vlan.md) | Optional | - |
| `WanSpeedtest` | [`*models.SynthetictestConfigWanSpeedtest`](../../doc/models/synthetictest-config-wan-speedtest.md) | Optional | - |

## Example (as JSON)

```json
{
  "aggressiveness": "auto",
  "disabled": false,
  "custom_probes": {
    "key0": {
      "aggressiveness": "med",
      "host": "host0",
      "port": 242,
      "threshold": 178,
      "type": "icmp"
    },
    "key1": {
      "aggressiveness": "med",
      "host": "host0",
      "port": 242,
      "threshold": 178,
      "type": "icmp"
    }
  },
  "lan_networks": [
    {
      "networks": [
        "networks4"
      ],
      "probes": [
        "probes5",
        "probes4"
      ]
    }
  ],
  "vlans": [
    {
      "custom_test_urls": [
        "custom_test_urls9"
      ],
      "disabled": false,
      "probes": [
        "probes7",
        "probes8",
        "probes9"
      ],
      "vlan_ids": [
        "String7",
        "String8"
      ]
    },
    {
      "custom_test_urls": [
        "custom_test_urls9"
      ],
      "disabled": false,
      "probes": [
        "probes7",
        "probes8",
        "probes9"
      ],
      "vlan_ids": [
        "String7",
        "String8"
      ]
    }
  ]
}
```

