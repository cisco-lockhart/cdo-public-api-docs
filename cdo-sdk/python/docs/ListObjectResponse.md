# ListObjectResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
**items** | [**List[UnifiedObjectListView]**](UnifiedObjectListView.md) | The list of objects retrieved. | [optional] 

## Example

```python
from cdo_sdk_python.models.list_object_response import ListObjectResponse

# TODO update the JSON string below
json = "{}"
# create an instance of ListObjectResponse from a JSON string
list_object_response_instance = ListObjectResponse.from_json(json)
# print the JSON string representation of the object
print(ListObjectResponse.to_json())

# convert the object into a dict
list_object_response_dict = list_object_response_instance.to_dict()
# create an instance of ListObjectResponse from a dict
list_object_response_form_dict = list_object_response.from_dict(list_object_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


