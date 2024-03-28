# Cisco Defense Orchestrator (CDO) Python SDK
The CDO Python SDK facilitates automated interaction with the Cisco Defense Orchestrator (CDO) API, enabling developers to integrate Cisco's cloud-based security policy and device management into Python applications.

For more information and detailed documentation, visit the [CDO Python SDK documentation](https://github.com/cisco-lockhart/cdo-public-api-docs/tree/main/cdo-sdk/python).

Read more about CDO [here](https://docs.defenseorchestrator.com/).

## Installation
Install the CDO Python SDK package with the following command:
```
pip install cdo-sdk-python
```

## Getting Started
After installation, import the SDK into your Python project and configure your API credentials.

## Usage Example
```python
import cdo_sdk_python, os
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# See configuration.py for a list of all supported configuration parameters
configuration = cdo_sdk_python.Configuration(
    host="https://edge.staging.cdo.cisco.com/api/rest"
)

# Configure Bearer authorization (JWT)
configuration.access_token = os.environ["BEARER_TOKEN"]

with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.TenantManagementApi(api_client)

    try:
        api_response = api_instance.get_tenants()
        print("The response of TenantManagementApi->get_tenants:\n")
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling TenantManagementApi->get_tenants: %s\n" % e)
```

## Developer Support
If you need technical support, choose from the options below.

### Consulting/Best Practices
Use the DevNet Community for general best practices, help, tips, or examples using Cisco APIs. Free to any DevNet member, log in and post your questions in the [Network Security forum](https://community.cisco.com/t5/network-security/bd-p/disc-network-security) using the `Cisco Defense Orchestrator (CDO)` label.

### Issues
If you identify issues with the CDO Public API, you can contact the [Cisco Technical Assistance Center (TAC)](https://www.cisco.com/c/en/us/support/web/tsd-cisco-worldwide-contacts.html) for support. First time? [Start here](https://www.cisco.com/c/dam/en/us/services/collateral/acquisitions/cjp-tac-support-guide.pdf).

- Internet: https://mycase.cloudapps.cisco.com/
- Worldwide Support Contacts: https://www.cisco.com/c/en/us/support/web/tsd-cisco-worldwide-contacts.html
- Webex: tac.connect@webex.bot

### Enhancement Requests
Please file enhancement requests in the [Network Security forum](https://community.cisco.com/t5/network-security/bd-p/disc-network-security) using the `Cisco Defense Orchestrator (CDO)` label.

### One-on-one Consulting
For one-on-one consulting, contact a [Cisco Developer Partner](https://www.cisco.com/c/en/us/partners/connect-with-a-partner.html) or your Cisco Sales Team for professional service options.
