# Getting Started

The following sections introduce you to the resources available in the Security Cloud Control Firewall Manager API and provide instructions for your first API request. Sample requests are shown using [curl](https://curl.se/), a CLI tool. 

## Supported RESTful Operations

The resources in the Security Cloud Control API support one or more `GET`, `POST`, `PATCH`,
`PUT`, and `DELETE` operations.

> **Note**:
> Not every resource supports all these operations.

- **GET**: There are two types of `GET` operations in Security Cloud Control: 
  - List `GET` operations, which return a paginated list of resources. List endpoints enable searching and sorting using a subset of the resource fields.
  - Individual GET operations, which return a single resource by its UID.
- **POST**: These operations are used to create resources or to trigger asynchronous operations. All POST operations that use an action verb initiate asynchronous operations, which can be tracked using the Security Cloud Control transaction API. 
- **PATCH**: These operations are used to modify existing resources.
- **PUT**: hese operations are used by the cloud-delivered Firewall Management Center (cdFMC) API endpoints to modify resources. 
- **DELETE**: hese operations are used to remove resources. All DELETE operations in the Security Cloud Control API are idempotent.

All the responses from the Security Cloud Control Firewall Manager API are in JSON format. Additionally, all the request payloads sent to a Security Cloud Control API must be in JSON format. 

## API Resources

At a high level, the Security Cloud Control Firewall Manager APIs support operations in these resources:

- **Inventory Management** -Manage the devices, device managers, cloud services, and templates in your Security Cloud Control tenant.
- **Connector Management** - Add, remove, and view information about the connectors used to communicate with devices in your Security Cloud Control tenant.
- **Cloud-delivered FMC** - Manage the cloud-delivered Firewall Management Center (cdFMC) in your Security Cloud Control tenant (if present).
- **Object Management** - Manage the firewall policy objects in your Security Cloud Control tenant.
- **User Management** - Manage the users in your Security Cloud Control tenant.
- **Tenant Management** - Manage your Security Cloud Control tenant or Managed Service Provider (MSP) Portal.
- **Remote Access Monitoring** - View and manage the Remote Access Virtual Private Network (RA VPN) sessions and Multi-factor Authentication (MFA) events in your Security Cloud Control tenant.
- **Search** - Perform searches across all the resources in your Security Cloud Control tenant.
- **Changelogs** - View a detailed history of all the changes made to your Security Cloud Control tenant.
- **Change Requests** - Manage change requests, which can be used to associate changes made to your
  security policy on Security Cloud Control with external change-management systems.
- **Transaction Management** - Track the progress of asynchronous operations triggered via Security Cloud Control APIs. For example, use transaction management to monitor the deployment progress of a configuration to an ASA.
- **Meta endpoints** - Information about Security Cloud Control Firewall Manager itself. This includes information such as details on available regions, cloud connector IPs and JSON Web Key Set.

## API User Prerequisites

To use the Security Cloud Control API, you must have an API token. We recommend creating
an [API-only user](https://docs.defenseorchestrator.com/c-secure-device-connector-sdc.html#!t-create-api-only-users.html)
in your Security Cloud Control tenant, and generating a token for it.

To use the Security Cloud Control API, you must have an API token. We recommend that you create an API-only user in your Security Cloud Control tenant and generate a token for that user. 

To perform the operations described in this page, the API-only user must have at least **Edit-only** privileges. However, some endpoints in the Security Cloud Control API may still require Admin or Super Admin privileges. 

## Base URL

Cisco Security Cloud Control is deployed across multiple regions, and the API is available in all of
them. The server URLs for each region is as follows:

- US: https://api.us.security.cisco.com/firewall
- EU: https://api.eu.security.cisco.com/firewall
- APJ: https://api.apj.security.cisco.com/firewall
- Australia: https://api.au.security.cisco.com/firewall
- India: https://api.in.security.cisco.com/firewall

> **Note**:
> If you are still using the old base URLs (`https://edge.<region>.cdo.cisco.com`), they will continue to
> work. But please switch your scripts to use the new URLs.

## 1. Authorization

In addition to the base URL, an Authorization header must be added to every API request with the
following format.

```
Authorization: Bearer $API_TOKEN
```

where `$API_TOKEN` is an API token for your Security Cloud Control tenant.

## 2. Get a list of devices

### Request

Use this `curl` command to make a request to retrieve the list of devices from the
`/v1/inventory/devices` endpoint.

```curl
curl -X GET \                                                                                                             08:48:01
--url https://api.us.security.cisco.com/firewall/v1/inventory/devices \
--header "Authorization: Bearer $API_TOKEN"
```

Replace `us` with the region you are using.

The response returned is a JSON object that has some metadata on the request itself, including the:

* **count**: The number of items in the collection. If this number is greater than `limit`, the
  number of items returned will be less than the total number of items in the list.
* **limit**: The maximum number of items returned by this request.
* **offset**: The offset from which the list was retrieved.
* **items:** The items in the response.

> **Note**:
> The `limit` and `offset` values can be set as query parameters to the request, and can be used to
> page through a list. By default, `limit` is set to `50`, and `offset` to `0`.

## 3. Get the UID and domain UUID of the cdFMC on a tenant

### Request

Use this `curl` command to request the retrieval of the UUID of the cloud-delivered FMC (cdFMC)  on your
tenant from the `/v1/inventory/managers` endpoint with the query parameter `q` set to
`deviceType:CDFMC`.

```curl
curl -X GET \
--url https://api.us.security.cisco.com/firewall/v1/inventory/managers?q=deviceType:CDFMC \
--header "Authorization: Bearer $API_TOKEN"
```

### Response

If your tenant has a cdFMC, the response will be a list of size 1. If not, the response will
be an empty list. A non-empty response should look something like this:

```json
{
  "count": 1,
  "limit": 50,
  "offset": 0,
  "items": [
    {
      "uid": "1ba692de-8166-4cee-90b5-6f817c3c2008",
      "name": "FMC",
      "deviceType": "CDFMC",
      "connectorType": "CDG",
      "address": "tf-managed-starks-2.app.staging.cdo.cisco.com:443",
      "deviceRoles": [],
      "softwareVersion": "7.7.10-build 1248",
      "connectivityState": "ONLINE",
      "configState": "SYNCED",
      "conflictDetectionState": "NO_CONFLICTS",
      "sseDeviceData": {
        "sseDeviceId": "78ce29e6-7579-41f7-b68b-1bde222d7d28",
        "sseProtocolVersion": "ONE"
      },
      "state": "DONE",
      "stateMachineDetails": {
        "identifier": "fmceOobStateMachine"
      },
      "fmcDomainUid": "e276abec-e0f2-11e3-8169-6d9ed49b625f",
      "optedInToAsaHealthMetrics": false
    }
  ]
}
```

#### Get the UID from the response

The response field `items[0].uid` contains the unique identifier of the cdFMC on your tenant. This
is a UUIDv4. If you are using [jq](https://jqlang.org/) to parse the response, extract the UID with
this command:

```bash
jq -r '.items[0].uid'
```

#### Get the domain UUID from the response

The response field `items[0].fmcDomainUid` contains the unique identifier of the cdFMC on your
tenant. This
is a UUIDv4. If using [jq](https://jqlang.org/) to parse the response, you can extract the domain
UUID from the `fmcDomainUid` field with this command:

```bash
jq -r '.items[0].fmcDomainUid'
```

## 4. Get all access policies on a cdFMC

> **Note**:
> All the cdFMC API endpoints are prefixed by `/v1/cdfmc`. These endpoints
> are
> compatible with the endpoints used to communicate with an on-premise FMC. So, for example, if you make an API call to
> the FMC to fetch all access policies using the endpoint
>
`https://<your-fmc-ip>/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies`,

> the same call to your cdFMC will be
>
`https://api.us.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies`.

### Request

Use this `curl` command to request the retrieval of the UUID of the cdFMC on your
tenant from the `/v1/cdfmc/api/fmc_config/v1/domain/<domainUid>/policy/accesspolicies` endpoint.

```curl
curl -X GET \
--url https://api.us.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies \
--header "Authorization: Bearer $API_TOKEN"
```

> **Notes**:
> - Replace `e276abec-e0f2-11e3-8169-6d9ed49b625f` with the domain UUID you retrieved in the
    previous step.
> - Replace `us` with your region

### Responses

The response will be a JSON object that is similar to this example. Note that the response will vary depending on the
access policies configured in your cdFMC.

```json
{
  "links": {
    "self": "https://api.int.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies?offset=0&limit=25"
  },
  "items": [
    {
      "type": "AccessPolicy",
      "links": {
        "self": "https://api.int.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0248CDE2-456F-0ed3-0000-004294967323"
      },
      "name": "Default Access Control Policy",
      "id": "0248CDE2-456F-0ed3-0000-004294967323"
    }
  ],
  "paging": {
    "offset": 0,
    "limit": 25,
    "count": 1,
    "pages": 1
  }
}
```

## 5. Get an access policy on a cdFMC

To retrieve the details of an access policy in cdFMC, you need the `id` (represented as a
universally unique UUIDv4 identifier) of the access policy, which you can retrieve using the API
call described in the previous example.

### Request

Use this `curl` command to make a request to retrieve the UUID of the Cloud-delivered FMC on your
tenant from the `/v1/cdfmc/api/fmc_config/v1/domain/<domainUid>/policy/accesspolicies/<id>`
endpoint.

```curl
curl -X GET \                                                                       22:52:21
--url https://api.us.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0248CDE2-456F-0ed3-0000-004294967323 \
--header "Authorization: Bearer $API_TOKEN"
```

> **Notes**:
> - Replace `e276abec-e0f2-11e3-8169-6d9ed49b625f` with the domain UUID you retrieved in the
    previous step.
> - Replace `0248CDE2-456F-0ed3-0000-004294967323` with the `id` of the access policy you want to
    retrieve.
> - Replace `us` with your region

### Response

The response will be a JSON object that is similar to this example. Note that the response will vary depending on the
access policy you selected.

```json
{
  "metadata": {
    "inherit": false,
    "lockingStatus": {
      "status": "UNLOCKED"
    },
    "timestamp": 1755053526892,
    "lastUser": {
      "name": "Firepower System"
    },
    "domain": {
      "name": "Global",
      "id": "e276abec-e0f2-11e3-8169-6d9ed49b625f",
      "type": "Domain"
    }
  },
  "type": "AccessPolicy",
  "links": {
    "self": "https://api.int.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0248CDE2-456F-0ed3-0000-004294967323"
  },
  "rules": {
    "refType": "list",
    "links": {
      "self": "https://api.int.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0248CDE2-456F-0ed3-0000-004294967323/accessrules"
    },
    "type": "AccessRule"
  },
  "securityIntelligence": {
    "id": "0248CDE2-456F-0ed3-0000-004294967371",
    "type": "SecurityIntelligencePolicy",
    "links": {
      "self": "https://api.int.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0248CDE2-456F-0ed3-0000-004294967323/securityintelligencepolicies/0248CDE2-456F-0ed3-0000-004294967371"
    }
  },
  "prefilterPolicySetting": {
    "type": "PrefilterPolicy",
    "name": "Default Prefilter Policy",
    "id": "4897c8f4-e211-4661-b0a4-25b0826cded9"
  },
  "defaultAction": {
    "sendEventsToFMC": false,
    "enableSyslog": false,
    "logBegin": false,
    "logEnd": false,
    "action": "BLOCK",
    "type": "AccessPolicyDefaultAction",
    "id": "0248CDE2-456F-0ed3-0000-000268434432"
  },
  "name": "Default Access Control Policy",
  "description": "Default Access Control Policy with default action block",
  "id": "0248CDE2-456F-0ed3-0000-004294967323"
}
```

## 6. Get access rules associated with an access policy on a cdFMC

To retrieve the access rules associated with an access policy in cdFMC, you need the `id` of the
access policy, which you can retrieve using the API call described in example *5*.

### Request

Use this `curl` command to to retrieve the access rules associated with a cdFMC access policy on your
tenant from the
`/v1/cdfmc/api/fmc_config/v1/domain/<domainUid>/policy/accesspolicies/<id>/accessrules` endpoint.

```curl
curl --location 'https://api.eu.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0248CDE2-456F-0ed3-0000-004294967323/accessrules' \
--header 'Accept: application/json' \
--header 'Authorization: $API_TOKEN'
```

> **Notes**:
> - Replace `e276abec-e0f2-11e3-8169-6d9ed49b625f` with the domain UUID you retrieved in the
    previous step.
> - Replace `0248CDE2-456F-0ed3-0000-004294967323` with the `id` of the access policy you want to
    retrieve the rules for.
> - Replace `us` with your region
> - Set the query parameter `expanded` to `true` (`?expanded=true`) to expand the
    response to include all the fields in the access rules. If you do not set this query
    parameter, the response will only include the `id`, `name`, and `links` fields of each access
    rule.

### Response

The response provided will be a JSON object that is similar to this example. Note that the actual response will vary depending
on the access rules configured on your cdFMC.

```json
{
  "links": {
    "self": "https://api.eu.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0AB208F0-1FF7-0ed3-0000-004294967299/accessrules?offset=0&limit=25"
  },
  "items": [
    {
      "links": {
        "self": "https://api.eu.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0AB208F0-1FF7-0ed3-0000-004294967299/accessrules/0AB208F0-1FF7-0ed3-0000-000268435465"
      },
      "type": "AccessRule",
      "id": "0AB208F0-1FF7-0ed3-0000-000268435465",
      "name": "webex-trust"
    },
    {
      "links": {
        "self": "https://api.eu.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0AB208F0-1FF7-0ed3-0000-004294967299/accessrules/0AB208F0-1FF7-0ed3-0000-000268435464"
      },
      "type": "AccessRule",
      "id": "0AB208F0-1FF7-0ed3-0000-000268435464",
      "name": "inside-inside-trust"
    },
    {
      "links": {
        "self": "https://api.eu.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0AB208F0-1FF7-0ed3-0000-004294967299/accessrules/0AB208F0-1FF7-0ed3-0000-000268435462"
      },
      "type": "AccessRule",
      "id": "0AB208F0-1FF7-0ed3-0000-000268435462",
      "name": "block-adult"
    },
    {
      "links": {
        "self": "https://api.eu.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0AB208F0-1FF7-0ed3-0000-004294967299/accessrules/0AB208F0-1FF7-0ed3-0000-000268435463"
      },
      "type": "AccessRule",
      "id": "0AB208F0-1FF7-0ed3-0000-000268435463",
      "name": "geo-ip-block"
    },
    {
      "links": {
        "self": "https://api.eu.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0AB208F0-1FF7-0ed3-0000-004294967299/accessrules/0AB208F0-1FF7-0ed3-0000-000268435461"
      },
      "type": "AccessRule",
      "id": "0AB208F0-1FF7-0ed3-0000-000268435461",
      "name": "allow-inside->outside"
    }
  ],
  "paging": {
    "offset": 0,
    "limit": 25,
    "count": 5,
    "pages": 1
  }
}
```

## 7. Add an access rule to an access policy on a cdFMC

To add an access rule to an access policy in cdFMC, you need the `id` of the access policy, which
you can retrieve using the API call described in example *5*.

### Request

Use this `curl` command to make a request to add an access rule to an access policy in cdFMC to
block unsuitable content. The assumption is that you have defined `inside` and `outside` security zones
on your cdFMC (
See [the documentation on creating security zones](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/create-multiple-security-zone-object/)
for more information).

Multiple options are available when creating an access rule. This example is a simple
access rule that blocks traffic from the `inside` security zone to the `outside` security zone
to URLs in the `adult` and `pornography` categories.

```curl
curl --location 'https://api.eu.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0AB208F0-1FF7-0ed3-0000-004294967299/accessrules?bulk=false' \
--header 'Content-Type: application/json' \
--header 'Authorization: $API_TOKEN' \
--data '{
    "enabled": true,
    "sourceZones": {
        "objects": [
            {
                "name": "inside",
                "id": "c2bdfa72-11e6-11f0-8c56-4a5ba7afc6cc",
                "type": "SecurityZone",
                "subType": "Routed"
            }
        ]
    },
    "destinationZones": {
        "objects": [
            {
                "name": "outside",
                "id": "ca032564-11e6-11f0-8c56-4a5ba7afc6cc",
                "type": "SecurityZone",
                "subType": "Routed"
            }
        ]
    },
    "logBegin": true,
    "logFiles": false,
    "sourceDynamicObjects": {},
    "urls": {
        "urlCategoriesWithReputation": [
            {
                "reputation": "ANY_AND_UNKNOWN",
                "type": "UrlCategoryAndReputation",
                "category": {
                    "name": "Adult",
                    "id": "abba9b63-bb10-4729-b901-2e2aa0f02006",
                    "type": "URLCategory"
                }
            },
            {
                "reputation": "ANY_AND_UNKNOWN",
                "type": "UrlCategoryAndReputation",
                "category": {
                    "name": "Pornography",
                    "id": "abba9b63-bb10-4729-b901-2e2aa0f02054",
                    "type": "URLCategory"
                }
            }
        ]
    },
    "sendEventsToFMC": true,
    "action": "BLOCK",
    "type": "AccessRule",
    "name": "block-unsuitable-content"
}'
```

> **Notes**:
> - Replace `e276abec-e0f2-11e3-8169-6d9ed49b625f` with the domain UUID you retrieved in the
    previous step.
> - Replace `0248CDE2-456F-0ed3-0000-004294967323` with the `id` of the access policy to which you want to
    add the access rule.
> - Replace `us` with your region

### Response

The response will be an `HTTP 201` with a JSON object that is similar to this example. Note that the response will vary
depending on the request you make.

```json
{
  "metadata": {
    "ruleIndex": 6,
    "section": "Default",
    "category": "--Undefined--",
    "accessPolicy": {
      "name": "Default Access Control Policy",
      "id": "0AB208F0-1FF7-0ed3-0000-004294967299",
      "type": "AccessPolicy"
    }
  },
  "links": {
    "self": "https://api.eu.security.cisco.com/firewall/v1/cdfmc/api/fmc_config/v1/domain/e276abec-e0f2-11e3-8169-6d9ed49b625f/policy/accesspolicies/0AB208F0-1FF7-0ed3-0000-004294967299/accessrules/0AB208F0-1FF7-0ed3-0000-000268435466"
  },
  "enabled": true,
  "sourceZones": {
    "objects": [
      {
        "name": "inside",
        "id": "c2bdfa72-11e6-11f0-8c56-4a5ba7afc6cc",
        "type": "SecurityZone",
        "subType": "Routed"
      }
    ]
  },
  "destinationZones": {
    "objects": [
      {
        "name": "outside",
        "id": "ca032564-11e6-11f0-8c56-4a5ba7afc6cc",
        "type": "SecurityZone",
        "subType": "Routed"
      }
    ]
  },
  "logBegin": true,
  "logFiles": false,
  "sourceDynamicObjects": {},
  "urls": {
    "urlCategoriesWithReputation": [
      {
        "reputation": "ANY_AND_UNKNOWN",
        "type": "UrlCategoryAndReputation",
        "category": {
          "name": "Adult",
          "id": "abba9b63-bb10-4729-b901-2e2aa0f02006",
          "type": "URLCategory"
        }
      },
      {
        "reputation": "ANY_AND_UNKNOWN",
        "type": "UrlCategoryAndReputation",
        "category": {
          "name": "Pornography",
          "id": "abba9b63-bb10-4729-b901-2e2aa0f02054",
          "type": "URLCategory"
        }
      }
    ]
  },
  "variableSet": {
    "name": "Default-Set",
    "id": "76fa83ea-c972-11e2-8be8-8e45bb1343c0",
    "type": "VariableSet"
  },
  "sendEventsToFMC": true,
  "enableSyslog": false,
  "vlanTags": {},
  "logEnd": false,
  "destinationDynamicObjects": {},
  "action": "BLOCK",
  "type": "AccessRule",
  "id": "0AB208F0-1FF7-0ed3-0000-000268435466",
  "name": "block-unsuitable-content"
}
```

## Try it out with Postman

Prefer to use Postman rather than code or the command line? Check out
our [Postman Collection](/docs/cisco-security-cloud-control-firewall-manager/postman-collections/).