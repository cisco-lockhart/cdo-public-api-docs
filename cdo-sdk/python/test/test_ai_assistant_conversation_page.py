# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.0.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.ai_assistant_conversation_page import AiAssistantConversationPage

class TestAiAssistantConversationPage(unittest.TestCase):
    """AiAssistantConversationPage unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AiAssistantConversationPage:
        """Test AiAssistantConversationPage
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AiAssistantConversationPage`
        """
        model = AiAssistantConversationPage()
        if include_optional:
            return AiAssistantConversationPage(
                count = 100,
                limit = 50,
                offset = 0,
                items = [
                    cdo_sdk_python.models.ai_assistant_conversation.AiAssistantConversation(
                        uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb', 
                        created_date = '2023-12-13T08:15:44Z', 
                        last_interaction_date = '2023-12-13T08:15:44Z', 
                        last_answer_date = '2023-12-13T08:15:44Z', 
                        title = 'Is example.com blocked on my FTD?', )
                    ]
            )
        else:
            return AiAssistantConversationPage(
        )
        """

    def testAiAssistantConversationPage(self):
        """Test AiAssistantConversationPage"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()