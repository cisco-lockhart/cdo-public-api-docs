# ConversationMessage

The list of items retrieved.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of the Message. | 
**type** | **str** | The type of the message, indicating whether it is a request or a response. | [optional] 
**in_reply_to** | **str** | The unique identifier of the message to which this message is replying. This field is populated only for messages of type RESPONSE. | [optional] 
**content** | **str** | The content of the message. | [optional] 
**created_date** | **datetime** | The time (UTC; represented using the RFC-3339 standard) at which the message was sent. | [optional] 

## Example

```python
from cdo_sdk_python.models.conversation_message import ConversationMessage

# TODO update the JSON string below
json = "{}"
# create an instance of ConversationMessage from a JSON string
conversation_message_instance = ConversationMessage.from_json(json)
# print the JSON string representation of the object
print(ConversationMessage.to_json())

# convert the object into a dict
conversation_message_dict = conversation_message_instance.to_dict()
# create an instance of ConversationMessage from a dict
conversation_message_form_dict = conversation_message.from_dict(conversation_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


