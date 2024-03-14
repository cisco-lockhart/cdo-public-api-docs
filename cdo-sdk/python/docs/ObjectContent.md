# ObjectContent

The content value of the override. This overrides the object's default content.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**protocol** | **str** | The service object protocol | [optional] 
**service_value** | [**ServiceObjectValueContent**](ServiceObjectValueContent.md) |  | [optional] 
**url** | **str** | The URL literal | 
**literals** | [**List[SingleContent]**](SingleContent.md) | List of content literals | [optional] 
**referenced_object_uids** | **List[str]** | Set of UIDs of the group&#39;s referenced objects | [optional] 
**literal** | **str** | The literal content of the network object | 

## Example

```python
from openapi_client.models.object_content import ObjectContent

# TODO update the JSON string below
json = "{}"
# create an instance of ObjectContent from a JSON string
object_content_instance = ObjectContent.from_json(json)
# print the JSON string representation of the object
print(ObjectContent.to_json())

# convert the object into a dict
object_content_dict = object_content_instance.to_dict()
# create an instance of ObjectContent from a dict
object_content_form_dict = object_content.from_dict(object_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


