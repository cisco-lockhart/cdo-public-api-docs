import cdo_python_sdk, os
from cdo_python_sdk.rest import ApiException
from pprint import pprint

# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_python_sdk.Configuration(
    host="https://edge.staging.cdo.cisco.com/api/rest"
)

# Configure Bearer authorization (JWT): bearerAuth
configuration.access_token = os.environ["BEARER_TOKEN"]


# Enter a context with an instance of the API client
with cdo_python_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_python_sdk.TenantManagementApi(api_client)

    try:
        api_response = api_instance.list_tenants()
        print("The response of TenantManagementApi->list_tenants:\n")
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling TenantManagementApi->list_tenants: %s\n" % e)