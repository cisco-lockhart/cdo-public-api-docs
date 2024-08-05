# MspCreateTenantInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tenant_name** | **str** | The name of the tenant to create. The tenant name can only contain alphabets, numbers, -, and _, and is limited to 50 characters. | 
**display_name** | **str** | A human-readable display name of the tenant to create. Use this field only if you want the display name to be different from the name of the tenant. | [optional] 

## Example

```python
from cdo_sdk_python.models.msp_create_tenant_input import MspCreateTenantInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspCreateTenantInput from a JSON string
msp_create_tenant_input_instance = MspCreateTenantInput.from_json(json)
# print the JSON string representation of the object
print(MspCreateTenantInput.to_json())

# convert the object into a dict
msp_create_tenant_input_dict = msp_create_tenant_input_instance.to_dict()
# create an instance of MspCreateTenantInput from a dict
msp_create_tenant_input_form_dict = msp_create_tenant_input.from_dict(msp_create_tenant_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


