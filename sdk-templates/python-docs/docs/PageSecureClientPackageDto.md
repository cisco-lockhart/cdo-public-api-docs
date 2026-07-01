# PageSecureClientPackageDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**items** | [**List[SecureClientPackageDto]**](SecureClientPackageDto.md) | The list of items retrieved. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The Security Cloud Control API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.page_secure_client_package_dto import PageSecureClientPackageDto

# TODO update the JSON string below
json = "{}"
# create an instance of PageSecureClientPackageDto from a JSON string
page_secure_client_package_dto_instance = PageSecureClientPackageDto.from_json(json)
# print the JSON string representation of the object
print(PageSecureClientPackageDto.to_json())

# convert the object into a dict
page_secure_client_package_dto_dict = page_secure_client_package_dto_instance.to_dict()
# create an instance of PageSecureClientPackageDto from a dict
page_secure_client_package_dto_form_dict = page_secure_client_package_dto.from_dict(page_secure_client_package_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


