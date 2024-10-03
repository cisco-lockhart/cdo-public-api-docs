const fs = require('fs');
const yaml = require('yaml');

function deleteSchema(openApiFile, schemaToDelete, referenceFile) {
  // Read the OpenAPI YAML file
  const fileContent = fs.readFileSync(openApiFile, 'utf8');
  const openApiDoc = yaml.parse(fileContent);

  // Remove the schema from components.schemas
  if (openApiDoc.components && openApiDoc.components.schemas) {
    delete openApiDoc.components.schemas[schemaToDelete];
  }

  // Update references in paths
  if (openApiDoc.paths) {
    for (const path in openApiDoc.paths) {
      const pathItem = openApiDoc.paths[path];
      for (const method in pathItem) {
        const operation = pathItem[method];
        if (operation.responses) {
          for (const response in operation.responses) {
            const responseObject = operation.responses[response];
            if (responseObject.content) {
              for (const contentType in responseObject.content) {
                const contentObject = responseObject.content[contentType];
                if (contentObject.schema && contentObject.schema.$ref === `#/components/schemas/${schemaToDelete}`) {
                  contentObject.schema.$ref = `./${referenceFile}#/components/schemas/${schemaToDelete}`;
                }
              }
            }
          }
        }
        if (operation.requestBody && operation.requestBody.content) {
          for (const contentType in operation.requestBody.content) {
            const contentObject = operation.requestBody.content[contentType];
            if (contentObject.schema && contentObject.schema.$ref === `#/components/schemas/${schemaToDelete}`) {
              contentObject.schema.$ref = `./${referenceFile}#/components/schemas/${schemaToDelete}`;
            }
          }
        }
      }
    }
  }

  if (openApiDoc.components && openApiDoc.components.schemas) {
    for (const schema in openApiDoc.components.schemas) {
      const schemaObject = openApiDoc.components.schemas[schema];
      if (schemaObject.properties) {
        for (const property in schemaObject.properties) {
          const propertyObject = schemaObject.properties[property];
          if (propertyObject.$ref === `#/components/schemas/${schemaToDelete}`) {
            propertyObject.$ref = `./${referenceFile}#/components/schemas/${schemaToDelete}`;
          }
        }
      }
    }
  }

  // Update references in schemas too


  // Write the updated YAML back to the file
  const updatedYaml = yaml.stringify(openApiDoc);
  fs.writeFileSync(openApiFile, updatedYaml, 'utf8');

  console.log(`${schemaToDelete} schema removed from ${openApiFile} and references updated.`);
}

// Usage example
const args = process.argv.slice(2);
if (args.length !== 3) {
  console.error('Usage: node delete_schema.js <openapi-file.yaml> <schema-to-delete> <reference-file>');
  process.exit(1);
}

const [openApiFile, schemaToDelete, referenceFile] = args;
deleteSchema(openApiFile, schemaToDelete, referenceFile);