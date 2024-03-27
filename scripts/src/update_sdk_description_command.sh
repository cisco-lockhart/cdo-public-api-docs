SETUP_PY_PATH="cdo-sdk/python/setup.py"
TEMP_FILE_PATH="cdo-sdk/python/setup.tmp.py"

# Replace the long_description block in setup.py
awk '
  /long_description=/{print "    long_description=open('\''../../SDK-DESC.md'\'', '\''r'\'', encoding='\''utf-8'\'').read(),"; inDescBlock=1; next}
  /^    [a-z_]+=/ {inDescBlock=0}
  inDescBlock {next}
  {print}
' "$SETUP_PY_PATH" > "$TEMP_FILE_PATH" && mv "$TEMP_FILE_PATH" "$SETUP_PY_PATH"