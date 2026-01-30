# RedirectView


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**application_context** | [**ApplicationContext**](ApplicationContext.md) |  | [optional] 
**attributes** | **Dict[str, str]** |  | [optional] 
**attributes_csv** | **str** |  | [optional] 
**attributes_map** | **Dict[str, object]** |  | [optional] 
**bean_name** | **str** |  | [optional] 
**content_type** | **str** |  | [optional] 
**context_relative** | **bool** |  | [optional] 
**encoding_scheme** | **str** |  | [optional] 
**expand_uri_template_variables** | **bool** |  | [optional] 
**expose_context_beans_as_attributes** | **bool** |  | [optional] 
**expose_model_attributes** | **bool** |  | [optional] 
**expose_path_variables** | **bool** |  | [optional] 
**exposed_context_bean_names** | **List[str]** |  | [optional] 
**hosts** | **List[str]** |  | [optional] 
**http10_compatible** | **bool** |  | [optional] 
**propagate_query_params** | **bool** |  | [optional] 
**propagate_query_properties** | **bool** |  | [optional] 
**redirect_view** | **bool** |  | [optional] 
**request_context_attribute** | **str** |  | [optional] 
**servlet_context** | [**ServletContext**](ServletContext.md) |  | [optional] 
**static_attributes** | **Dict[str, object]** |  | [optional] 
**status_code** | [**HttpStatusCode**](HttpStatusCode.md) |  | [optional] 
**url** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.redirect_view import RedirectView

# TODO update the JSON string below
json = "{}"
# create an instance of RedirectView from a JSON string
redirect_view_instance = RedirectView.from_json(json)
# print the JSON string representation of the object
print(RedirectView.to_json())

# convert the object into a dict
redirect_view_dict = redirect_view_instance.to_dict()
# create an instance of RedirectView from a dict
redirect_view_from_dict = RedirectView.from_dict(redirect_view_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


