# VirtualAccountsPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**items** | [**List[VirtualAccountDto]**](VirtualAccountDto.md) | The list of items retrieved. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The Security Cloud Control API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.virtual_accounts_page import VirtualAccountsPage

# TODO update the JSON string below
json = "{}"
# create an instance of VirtualAccountsPage from a JSON string
virtual_accounts_page_instance = VirtualAccountsPage.from_json(json)
# print the JSON string representation of the object
print(VirtualAccountsPage.to_json())

# convert the object into a dict
virtual_accounts_page_dict = virtual_accounts_page_instance.to_dict()
# create an instance of VirtualAccountsPage from a dict
virtual_accounts_page_form_dict = virtual_accounts_page.from_dict(virtual_accounts_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


