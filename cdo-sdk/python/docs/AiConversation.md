# AiConversation


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of the conversation. | 
**created_date** | **datetime** | The time (UTC; represented using the RFC-3339 standard) at which the conversation was created. | [optional] 
**last_interaction_date** | **datetime** | The time (UTC; represented using the RFC-3339 standard) at which the user interacted with this conversation. | [optional] 
**last_answer_date** | **datetime** | The time (UTC; represented using the RFC-3339 standard) at which an answer was received. | [optional] 
**title** | **str** | The title of the conversation. This is set to the first question asked as part of the conversation. | [optional] 

## Example

```python
from cdo_sdk_python.models.ai_conversation import AiConversation

# TODO update the JSON string below
json = "{}"
# create an instance of AiConversation from a JSON string
ai_conversation_instance = AiConversation.from_json(json)
# print the JSON string representation of the object
print(AiConversation.to_json())

# convert the object into a dict
ai_conversation_dict = ai_conversation_instance.to_dict()
# create an instance of AiConversation from a dict
ai_conversation_form_dict = ai_conversation.from_dict(ai_conversation_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


