# Authentication

All requests to CDO APIs must be authenticated using the Bearer token method for the tenant you want to make requests for. You can generate a long-lived API token by logging into CDO and going to the Settings page. We recommend creating an [API-only user](https://docs.defenseorchestrator.com/c-secure-device-connector-sdc.html#!t-create-api-only-users.html) so that your scripts are not tied to a single implementation.

![img](../images/cdo-api-token.png)
