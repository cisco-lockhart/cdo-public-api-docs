# ChangeRequestPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The SCC API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
**items** | [**List[ChangeRequest]**](ChangeRequest.md) | The list of items retrieved. | [optional] 

## Example

```python
from cdo_sdk_python.models.change_request_page import ChangeRequestPage

# TODO update the JSON string below
json = "{}"
# create an instance of ChangeRequestPage from a JSON string
change_request_page_instance = ChangeRequestPage.from_json(json)
# print the JSON string representation of the object
print(ChangeRequestPage.to_json())

# convert the object into a dict
change_request_page_dict = change_request_page_instance.to_dict()
# create an instance of ChangeRequestPage from a dict
change_request_page_form_dict = change_request_page.from_dict(change_request_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


