# AccessGroupPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
**items** | [**List[AccessGroup]**](AccessGroup.md) | The list of items retrieved. | [optional] 

## Example

```python
from cdo_sdk_python.models.access_group_page import AccessGroupPage

# TODO update the JSON string below
json = "{}"
# create an instance of AccessGroupPage from a JSON string
access_group_page_instance = AccessGroupPage.from_json(json)
# print the JSON string representation of the object
print(AccessGroupPage.to_json())

# convert the object into a dict
access_group_page_dict = access_group_page_instance.to_dict()
# create an instance of AccessGroupPage from a dict
access_group_page_form_dict = access_group_page.from_dict(access_group_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


