# ApplicationContext


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**application_name** | **str** |  | [optional] 
**autowire_capable_bean_factory** | **object** |  | [optional] 
**bean_definition_count** | **int** |  | [optional] 
**bean_definition_names** | **List[str]** |  | [optional] 
**class_loader** | [**ApplicationContextClassLoader**](ApplicationContextClassLoader.md) |  | [optional] 
**display_name** | **str** |  | [optional] 
**environment** | [**Environment**](Environment.md) |  | [optional] 
**id** | **str** |  | [optional] 
**parent** | [**ApplicationContext**](ApplicationContext.md) |  | [optional] 
**parent_bean_factory** | **object** |  | [optional] 
**startup_date** | **int** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.application_context import ApplicationContext

# TODO update the JSON string below
json = "{}"
# create an instance of ApplicationContext from a JSON string
application_context_instance = ApplicationContext.from_json(json)
# print the JSON string representation of the object
print(ApplicationContext.to_json())

# convert the object into a dict
application_context_dict = application_context_instance.to_dict()
# create an instance of ApplicationContext from a dict
application_context_from_dict = ApplicationContext.from_dict(application_context_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


