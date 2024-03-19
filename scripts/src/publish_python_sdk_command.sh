cd cdo-sdk/python
pip install -r requirements.txt
pip install wheel
pip install twine
python3 setup.py sdist bdist_wheel

# need to set pypi creds
twine upload dist/*