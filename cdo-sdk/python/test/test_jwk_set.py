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

from cdo_sdk_python.models.jwk_set import JwkSet

class TestJwkSet(unittest.TestCase):
    """JwkSet unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> JwkSet:
        """Test JwkSet
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `JwkSet`
        """
        model = JwkSet()
        if include_optional:
            return JwkSet(
                keys = [
                    cdo_sdk_python.models.json_web_key.JsonWebKey(
                        kty = '', 
                        e = '', 
                        use = '', 
                        kid = '', 
                        alg = '', 
                        n = '', )
                    ]
            )
        else:
            return JwkSet(
        )
        """

    def testJwkSet(self):
        """Test JwkSet"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
