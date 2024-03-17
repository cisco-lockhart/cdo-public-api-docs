# MspAddTenantInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_token** | **str** | The Tenant API Access Token. | 

## Example

```python
from cdo-python-sdk.models.msp_add_tenant_input import MspAddTenantInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspAddTenantInput from a JSON string
msp_add_tenant_input_instance = MspAddTenantInput.from_json(json)
# print the JSON string representation of the object
print(MspAddTenantInput.to_json())

# convert the object into a dict
msp_add_tenant_input_dict = msp_add_tenant_input_instance.to_dict()
# create an instance of MspAddTenantInput from a dict
msp_add_tenant_input_form_dict = msp_add_tenant_input.from_dict(msp_add_tenant_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


