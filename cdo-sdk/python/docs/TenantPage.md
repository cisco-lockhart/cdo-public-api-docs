# TenantPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
**items** | [**List[Tenant]**](Tenant.md) | The list of items retrieved. | [optional] 

## Example

```python
from openapi_client.models.tenant_page import TenantPage

# TODO update the JSON string below
json = "{}"
# create an instance of TenantPage from a JSON string
tenant_page_instance = TenantPage.from_json(json)
# print the JSON string representation of the object
print(TenantPage.to_json())

# convert the object into a dict
tenant_page_dict = tenant_page_instance.to_dict()
# create an instance of TenantPage from a dict
tenant_page_form_dict = tenant_page.from_dict(tenant_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


