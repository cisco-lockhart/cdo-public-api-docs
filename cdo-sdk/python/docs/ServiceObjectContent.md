# ServiceObjectContent


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**protocol** | **str** | The service object protocol | [optional] 
**service_value** | [**ServiceObjectValueContent**](ServiceObjectValueContent.md) |  | [optional] 

## Example

```python
from openapi_client.models.service_object_content import ServiceObjectContent

# TODO update the JSON string below
json = "{}"
# create an instance of ServiceObjectContent from a JSON string
service_object_content_instance = ServiceObjectContent.from_json(json)
# print the JSON string representation of the object
print(ServiceObjectContent.to_json())

# convert the object into a dict
service_object_content_dict = service_object_content_instance.to_dict()
# create an instance of ServiceObjectContent from a dict
service_object_content_form_dict = service_object_content.from_dict(service_object_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


