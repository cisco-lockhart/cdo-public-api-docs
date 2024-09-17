# ApplicationContextClassLoaderParentUnnamedModuleClassLoader


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** |  | [optional] 
**registered_as_parallel_capable** | **bool** |  | [optional] 
**defined_packages** | [**List[ApplicationContextClassLoaderParentUnnamedModuleClassLoaderDefinedPackagesInner]**](ApplicationContextClassLoaderParentUnnamedModuleClassLoaderDefinedPackagesInner.md) |  | [optional] 
**default_assertion_status** | **bool** |  | [optional] 

## Example

```python
from cdo_sdk_python.models.application_context_class_loader_parent_unnamed_module_class_loader import ApplicationContextClassLoaderParentUnnamedModuleClassLoader

# TODO update the JSON string below
json = "{}"
# create an instance of ApplicationContextClassLoaderParentUnnamedModuleClassLoader from a JSON string
application_context_class_loader_parent_unnamed_module_class_loader_instance = ApplicationContextClassLoaderParentUnnamedModuleClassLoader.from_json(json)
# print the JSON string representation of the object
print(ApplicationContextClassLoaderParentUnnamedModuleClassLoader.to_json())

# convert the object into a dict
application_context_class_loader_parent_unnamed_module_class_loader_dict = application_context_class_loader_parent_unnamed_module_class_loader_instance.to_dict()
# create an instance of ApplicationContextClassLoaderParentUnnamedModuleClassLoader from a dict
application_context_class_loader_parent_unnamed_module_class_loader_form_dict = application_context_class_loader_parent_unnamed_module_class_loader.from_dict(application_context_class_loader_parent_unnamed_module_class_loader_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


