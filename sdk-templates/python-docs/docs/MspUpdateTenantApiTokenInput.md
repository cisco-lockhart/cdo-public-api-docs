# MspUpdateTenantApiTokenInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_token** | **str** | The new Tenant API Access Token. | 

## Example

```python
from scc_firewall_manager_sdk.models.msp_update_tenant_api_token_input import MspUpdateTenantApiTokenInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspUpdateTenantApiTokenInput from a JSON string
msp_update_tenant_api_token_input_instance = MspUpdateTenantApiTokenInput.from_json(json)
# print the JSON string representation of the object
print(MspUpdateTenantApiTokenInput.to_json())

# convert the object into a dict
msp_update_tenant_api_token_input_dict = msp_update_tenant_api_token_input_instance.to_dict()
# create an instance of MspUpdateTenantApiTokenInput from a dict
msp_update_tenant_api_token_input_form_dict = msp_update_tenant_api_token_input.from_dict(msp_update_tenant_api_token_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


