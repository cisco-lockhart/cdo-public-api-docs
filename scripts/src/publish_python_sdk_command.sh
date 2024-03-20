# Retrieve the secret value from AWS Secrets Manager and update .pypirc
secret_value=$(aws secretsmanager get-secret-value --secret-id arn:aws:secretsmanager:us-west-2:107042026245:secret:jenkins-pypi-credentials-mD4NdK --query SecretString --output text)
sed -i "s/\"CHANGE_ME\"/$secret_value/" .pypirc

cd cdo-sdk/python
pip install -r requirements.txt
pip install wheel
pip install twine

# Prepare to publish & Publish to PyPI
python3 setup.py sdist bdist_wheel
twine upload dist/*