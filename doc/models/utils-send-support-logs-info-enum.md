
# Utils Send Support Logs Info Enum

Optional, enum:
* `code-dumps`: Upload all core dump files, if any found
* `full`: Upload 1 file with output of `request support information`, 1 file that concatenates all `/var/log/outbound-ssh.log*` files, all core dump files, the 5 most recent `/var/log/messages*` files, and Mist agent logs
* `messages`: Upload 1 to 10 `/var/log/messages*` files
* `outbound-ssh`: Upload 1 file that concatenates all `/var/log/outbound-ssh.log*` files
* `process`: Upload 1 file with output of show ``system processes extensive`` *``var-logs`: Upload all non-empty files in the`/var/log/` directory

## Enumeration

`UtilsSendSupportLogsInfoEnum`

## Fields

| Name |
|  --- |
| `code-dumps` |
| `full` |
| `messages` |
| `outbound-ssh` |
| `process` |
| `var-logs` |

