# coding: utf-8

"""
    Cisco Defense Orchestrator API

    Use the interactive documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 0.0.1
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.models.labels import Labels

class TestLabels(unittest.TestCase):
    """Labels unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> Labels:
        """Test Labels
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `Labels`
        """
        model = Labels()
        if include_optional:
            return Labels(
                grouped_labels = {"group1":["label-1","label-2"],"group2":["label-1","label-2"]},
                ungrouped_labels = ["label-a","label-b","label-c"]
            )
        else:
            return Labels(
        )
        """

    def testLabels(self):
        """Test Labels"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()