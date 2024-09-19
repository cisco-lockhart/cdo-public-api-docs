# ApplicationContext


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**parent** | [**ApplicationContext**](ApplicationContext.md) |  | [optional] 
**id** | **str** |  | [optional] 
**display_name** | **str** |  | [optional] 
**autowire_capable_bean_factory** | **object** |  | [optional] 
**application_name** | **str** |  | [optional] 
**startup_date** | **int** |  | [optional] 
**environment** | [**Environment**](Environment.md) |  | [optional] 
**bean_definition_names** | **List[str]** |  | [optional] 
**bean_definition_count** | **int** |  | [optional] 
**parent_bean_factory** | **object** |  | [optional] 
**class_loader** | [**ApplicationContextClassLoader**](ApplicationContextClassLoader.md) |  | [optional] 

## Example

```python
from cdo_sdk_python.models.application_context import ApplicationContext

# TODO update the JSON string below
json = "{}"
# create an instance of ApplicationContext from a JSON string
application_context_instance = ApplicationContext.from_json(json)
# print the JSON string representation of the object
print(ApplicationContext.to_json())

# convert the object into a dict
application_context_dict = application_context_instance.to_dict()
# create an instance of ApplicationContext from a dict
application_context_form_dict = application_context.from_dict(application_context_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


