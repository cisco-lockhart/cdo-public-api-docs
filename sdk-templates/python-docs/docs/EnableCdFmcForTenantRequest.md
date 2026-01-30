# EnableCdFmcForTenantRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**dedicated_cd_fmc_instance** | **bool** |  | 

## Example

```python
from scc_firewall_manager_sdk.models.enable_cd_fmc_for_tenant_request import EnableCdFmcForTenantRequest

# TODO update the JSON string below
json = "{}"
# create an instance of EnableCdFmcForTenantRequest from a JSON string
enable_cd_fmc_for_tenant_request_instance = EnableCdFmcForTenantRequest.from_json(json)
# print the JSON string representation of the object
print(EnableCdFmcForTenantRequest.to_json())

# convert the object into a dict
enable_cd_fmc_for_tenant_request_dict = enable_cd_fmc_for_tenant_request_instance.to_dict()
# create an instance of EnableCdFmcForTenantRequest from a dict
enable_cd_fmc_for_tenant_request_form_dict = enable_cd_fmc_for_tenant_request.from_dict(enable_cd_fmc_for_tenant_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


