# SingleContent

List of content literals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**protocol** | **str** | The service object protocol | [optional] 
**service_value** | [**ServiceObjectValueContent**](ServiceObjectValueContent.md) |  | [optional] 
**url** | **str** | The URL literal | 
**literal** | **str** | The literal content of the network object | 

## Example

```python
from cdo-python-sdk.models.single_content import SingleContent

# TODO update the JSON string below
json = "{}"
# create an instance of SingleContent from a JSON string
single_content_instance = SingleContent.from_json(json)
# print the JSON string representation of the object
print(SingleContent.to_json())

# convert the object into a dict
single_content_dict = single_content_instance.to_dict()
# create an instance of SingleContent from a dict
single_content_form_dict = single_content.from_dict(single_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


