# MspDeleteTenantInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tenant_name** | **str** | The name of the tenant to remove from the MSP portal. | 
**tenant_cdo_region** | **str** | The region in which the tenant to be removed resides. | 

## Example

```python
from cdo_sdk_python.models.msp_delete_tenant_input import MspDeleteTenantInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspDeleteTenantInput from a JSON string
msp_delete_tenant_input_instance = MspDeleteTenantInput.from_json(json)
# print the JSON string representation of the object
print(MspDeleteTenantInput.to_json())

# convert the object into a dict
msp_delete_tenant_input_dict = msp_delete_tenant_input_instance.to_dict()
# create an instance of MspDeleteTenantInput from a dict
msp_delete_tenant_input_form_dict = msp_delete_tenant_input.from_dict(msp_delete_tenant_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


