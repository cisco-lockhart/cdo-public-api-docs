# AiAssistantConversationPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
**items** | [**List[AiConversation]**](AiConversation.md) | The list of items retrieved. | [optional] 

## Example

```python
from cdo_sdk_python.models.ai_assistant_conversation_page import AiAssistantConversationPage

# TODO update the JSON string below
json = "{}"
# create an instance of AiAssistantConversationPage from a JSON string
ai_assistant_conversation_page_instance = AiAssistantConversationPage.from_json(json)
# print the JSON string representation of the object
print(AiAssistantConversationPage.to_json())

# convert the object into a dict
ai_assistant_conversation_page_dict = ai_assistant_conversation_page_instance.to_dict()
# create an instance of AiAssistantConversationPage from a dict
ai_assistant_conversation_page_form_dict = ai_assistant_conversation_page.from_dict(ai_assistant_conversation_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


