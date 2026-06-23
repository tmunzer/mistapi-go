
# Custom Header Signature



Documentation for accessing and setting credentials for apiToken.

## Auth Credentials

| Name | Type | Description | Setter | Getter |
|  --- | --- | --- | --- | --- |
| authorization | `string` | Preferred authentication method for automation and integrations. Send the API token in the HTTP `Authorization` header.<br><br>**Format**:<br>`Authorization: Token {apitoken}`<br><br>**Notes**:<br><br>* An API token generated for a specific admin has the same privileges as that admin<br>* An API token is automatically removed if it is not used for more than 90 days<br>* SSO admins cannot generate admin API tokens. Use organization API tokens when scoped Org/Site privileges are needed. | `WithAuthorization` | `Authorization()` |



**Note:** Required auth credentials can be set using `WithApiTokenCredentials()` by providing a credentials instance with `NewApiTokenCredentials()` in the configuration initialization and accessed using the `ApiTokenCredentials()` method in the configuration instance.

## Usage Example

### Client Initialization

You must provide credentials in the client as shown in the following code snippet.

```go
package main

import (
    "mistapi"
)

func main() {
    client := mistapi.NewClient(
    mistapi.CreateConfiguration(
            mistapi.WithApiTokenCredentials(
                mistapi.NewApiTokenCredentials("Authorization"),
            ),
        ),
    )
}
```


