
# Wlan Mist Nac

Mist NAC RADIUS settings for a WLAN

## Structure

`WlanMistNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctInterimInterval` | `*int` | Optional | How frequently should interim accounting be reported, 60-65535. default is 0 (use one specified in Access-Accept request from Server). Very frequent messages can affect the performance of the RADIUS server, 600 and up is recommended when enabled.<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 65535` |
| `AuthServersRetries` | `*int` | Optional | RADIUS auth session retries. Following fast timers are set if `fast_dot1x_timers` knob is enabled. "retries" are set to value of `auth_servers_timeout`. "max-requests" is also set when setting `auth_servers_retries` is set to default value to 3.<br><br>**Default**: `2`<br><br>**Constraints**: `>= 1`, `<= 10` |
| `AuthServersTimeout` | `*int` | Optional | RADIUS auth session timeout. Following fast timers are set if `fast_dot1x_timers` knob is enabled. "quite-period" and "transmit-period" are set to half the value of `auth_servers_timeout`. "supplicant-timeout" is also set when setting `auth_servers_timeout` is set to default value of 10.<br><br>**Default**: `5`<br><br>**Constraints**: `>= 1`, `<= 30` |
| `CoaEnabled` | `*bool` | Optional | Allows a RADIUS server to dynamically modify the authorization status of a user session.<br><br>**Default**: `false` |
| `CoaPort` | `*int` | Optional | the communication port used for “Change of Authorization” (CoA) messages<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `Enabled` | `*bool` | Optional | When enabled:<br><br>* `auth_servers` is ignored<br>* `acct_servers` is ignored<br>* `auth_servers_*` are ignored<br>* `coa_servers` is ignored<br>* `radsec` is ignored<br>* `coa_enabled` is assumed<br><br>**Default**: `false` |
| `FastDot1xTimers` | `*bool` | Optional | If set to true, sets default fast-timers with values calculated from `auth_servers_timeout` and `auth_server_retries`.<br><br>**Default**: `false` |
| `Network` | `models.Optional[string]` | Optional | Which network the mist nac server resides in |
| `SourceIp` | `models.Optional[string]` | Optional | In case there is a static IP for this network, we can specify it using source ip |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanMistNac := models.WlanMistNac{
        AcctInterimInterval:  models.ToPointer(60),
        AuthServersRetries:   models.ToPointer(3),
        AuthServersTimeout:   models.ToPointer(5),
        CoaEnabled:           models.ToPointer(false),
        CoaPort:              models.ToPointer(3799),
        Enabled:              models.ToPointer(false),
        FastDot1xTimers:      models.ToPointer(false),
        Network:              models.NewOptional(models.ToPointer("default")),
        SourceIp:             models.NewOptional(models.ToPointer("1.2.3.4")),
    }

}
```

