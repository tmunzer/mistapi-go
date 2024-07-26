
# Switch Vrf Instances

Property key is the network name

## Structure

`SwitchVrfInstances`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Networks` | `[]string` | Optional | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "guest": {
    "extra_routes": {
      "0.0.0.0/0": {
        "via": "192.168.31.1"
      }
    },
    "networks": [
      "guest"
    ]
  },
  "networks": [
    "networks6",
    "networks7"
  ]
}
```

