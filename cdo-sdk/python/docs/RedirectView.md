# RedirectView


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**application_context** | [**ApplicationContext**](ApplicationContext.md) |  | [optional] 
**servlet_context** | [**RedirectViewServletContext**](RedirectViewServletContext.md) |  | [optional] 
**content_type** | **str** |  | [optional] 
**request_context_attribute** | **str** |  | [optional] 
**static_attributes** | **Dict[str, object]** |  | [optional] 
**expose_path_variables** | **bool** |  | [optional] 
**expose_context_beans_as_attributes** | **bool** |  | [optional] 
**exposed_context_bean_names** | **List[str]** |  | [optional] 
**bean_name** | **str** |  | [optional] 
**url** | **str** |  | [optional] 
**context_relative** | **bool** |  | [optional] 
**http10_compatible** | **bool** |  | [optional] 
**expose_model_attributes** | **bool** |  | [optional] 
**encoding_scheme** | **str** |  | [optional] 
**status_code** | **str** |  | [optional] 
**expand_uri_template_variables** | **bool** |  | [optional] 
**propagate_query_params** | **bool** |  | [optional] 
**hosts** | **List[str]** |  | [optional] 
**redirect_view** | **bool** |  | [optional] 
**propagate_query_properties** | **bool** |  | [optional] 
**attributes_map** | **Dict[str, object]** |  | [optional] 
**attributes_csv** | **str** |  | [optional] 
**attributes** | **Dict[str, str]** |  | [optional] 

## Example

```python
from cdo_sdk_python.models.redirect_view import RedirectView

# TODO update the JSON string below
json = "{}"
# create an instance of RedirectView from a JSON string
redirect_view_instance = RedirectView.from_json(json)
# print the JSON string representation of the object
print(RedirectView.to_json())

# convert the object into a dict
redirect_view_dict = redirect_view_instance.to_dict()
# create an instance of RedirectView from a dict
redirect_view_form_dict = redirect_view.from_dict(redirect_view_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


