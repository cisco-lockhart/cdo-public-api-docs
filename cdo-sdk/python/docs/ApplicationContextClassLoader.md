# ApplicationContextClassLoader


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | [optional] 
**registered_as_parallel_capable** | **bool** |  | [optional] 
**parent** | [**ApplicationContextClassLoaderParent**](ApplicationContextClassLoaderParent.md) |  | [optional] 
**unnamed_module** | [**ApplicationContextClassLoaderParentUnnamedModule**](ApplicationContextClassLoaderParentUnnamedModule.md) |  | [optional] 
**defined_packages** | [**List[ApplicationContextClassLoaderParentUnnamedModuleClassLoaderDefinedPackagesInner]**](ApplicationContextClassLoaderParentUnnamedModuleClassLoaderDefinedPackagesInner.md) |  | [optional] 
**default_assertion_status** | **bool** |  | [optional] 

## Example

```python
from cdo_sdk_python.models.application_context_class_loader import ApplicationContextClassLoader

# TODO update the JSON string below
json = "{}"
# create an instance of ApplicationContextClassLoader from a JSON string
application_context_class_loader_instance = ApplicationContextClassLoader.from_json(json)
# print the JSON string representation of the object
print(ApplicationContextClassLoader.to_json())

# convert the object into a dict
application_context_class_loader_dict = application_context_class_loader_instance.to_dict()
# create an instance of ApplicationContextClassLoader from a dict
application_context_class_loader_form_dict = application_context_class_loader.from_dict(application_context_class_loader_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


