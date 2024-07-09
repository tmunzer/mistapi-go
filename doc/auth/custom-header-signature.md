
# Custom Header Signature



Documentation for accessing and setting credentials for apiToken.

## Auth Credentials

| Name | Type | Description | Setter | Getter |
|  --- | --- | --- | --- | --- |
| Authorization | `string` | Like many other API providers, it’s also possible to generate API Tokens to be used (in HTTP Header) for authentication. An API token ties to a Admin with equal or less privileges.<br><br>**Format**:<br>API Token value format is `Token {apitoken}`<br><br>**Notes**:<br><br>* an API token generated for a specific admin has the same privilege as the user<br>* an API token will be automatically removed if not used for > 90 days<br>* SSO admins cannot generate these API tokens. Refer Org level API tokens which can have privileges of a specific Org/Site for more information. | `WithAuthorization` | `Authorization` |



**Note:** Required auth credentials can be set using `WithApiTokenCredentials()` by providing a credentials instance with `NewApiTokenCredentials()` in the configuration initialization and accessed using the `ApiTokenCredentials()` method in the configuration instance.

## Usage Example

### Client Initialization

You must provide credentials in the client as shown in the following code snippet.

```go
client := mistapigo.NewClient(
    mistapigo.CreateConfiguration(
        mistapigo.WithApiTokenCredentials(
            mistapigo.NewApiTokenCredentials("Authorization"),
        ),
    ),
)
```


