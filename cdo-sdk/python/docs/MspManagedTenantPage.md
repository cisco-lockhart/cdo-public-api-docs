# MspManagedTenantPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
**items** | [**List[MspManagedTenant]**](MspManagedTenant.md) | The list of items retrieved. | [optional] 

## Example

```python
from cdo_sdk_python.models.msp_managed_tenant_page import MspManagedTenantPage

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedTenantPage from a JSON string
msp_managed_tenant_page_instance = MspManagedTenantPage.from_json(json)
# print the JSON string representation of the object
print(MspManagedTenantPage.to_json())

# convert the object into a dict
msp_managed_tenant_page_dict = msp_managed_tenant_page_instance.to_dict()
# create an instance of MspManagedTenantPage from a dict
msp_managed_tenant_page_form_dict = msp_managed_tenant_page.from_dict(msp_managed_tenant_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


