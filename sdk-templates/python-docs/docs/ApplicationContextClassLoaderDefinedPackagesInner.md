# ApplicationContextClassLoaderDefinedPackagesInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**annotations** | **List[object]** |  | [optional] 
**declared_annotations** | **List[object]** |  | [optional] 
**implementation_title** | **str** |  | [optional] 
**implementation_vendor** | **str** |  | [optional] 
**implementation_version** | **str** |  | [optional] 
**name** | **str** |  | [optional] 
**sealed** | **bool** |  | [optional] 
**specification_title** | **str** |  | [optional] 
**specification_vendor** | **str** |  | [optional] 
**specification_version** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.application_context_class_loader_defined_packages_inner import ApplicationContextClassLoaderDefinedPackagesInner

# TODO update the JSON string below
json = "{}"
# create an instance of ApplicationContextClassLoaderDefinedPackagesInner from a JSON string
application_context_class_loader_defined_packages_inner_instance = ApplicationContextClassLoaderDefinedPackagesInner.from_json(json)
# print the JSON string representation of the object
print(ApplicationContextClassLoaderDefinedPackagesInner.to_json())

# convert the object into a dict
application_context_class_loader_defined_packages_inner_dict = application_context_class_loader_defined_packages_inner_instance.to_dict()
# create an instance of ApplicationContextClassLoaderDefinedPackagesInner from a dict
application_context_class_loader_defined_packages_inner_from_dict = ApplicationContextClassLoaderDefinedPackagesInner.from_dict(application_context_class_loader_defined_packages_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


