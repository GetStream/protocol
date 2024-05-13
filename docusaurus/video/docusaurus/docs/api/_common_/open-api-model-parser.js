import videoApiJson from '../video-openapi.json';

export const parseModel = ({ modelName, modelFilter, recursive = true, apiJson = videoApiJson }) => {
    const models = modelName ? [{name: modelName, properties: []}] : modelFilter(apiJson).map(_modelName => ({name: _modelName, properties: []}));
  
    for (let i = 0; i < models.length; i++) {
      const model = models[i];
  
      const schemaJson = apiJson.components.schemas[model.name];
      const propertiesJson = schemaJson.properties;
  
      model.properties = Object.keys(propertiesJson).map(propertyName => {
        const property = propertiesJson[propertyName];
        
        const type = parsePropertyType(property);

        // enums are not yet parsed
        const isEnum = apiJson.components.schemas[type.name] && !apiJson.components.schemas[type.name].properties;
        const shouldDisplayTypeDefLink = recursive && !isEnum;

        if (recursive && !isEnum && apiJson.components.schemas[type.name] && !models.find(r => r.name === type.name)) {
            models.push({name: type.name, properties: []});
        }
  
        const description = parsePropertyDescription(property);
  
        const isRequired = schemaJson.required?.includes(propertyName);
        const constraints = parsePropertyConstraints(property, isRequired);
        
  
        return {
          name: propertyName,
          type: {...type, definitionLink: shouldDisplayTypeDefLink ? type.definitionLink : undefined},
          constraints,
          description
        }
      });
    }
  
    return models;
  }

  export const parsePropertyType = (property) => {
    let name;
    let definitionLink;
    let formattedName;
    let isArray = property.type === 'array';
    if (property.$ref || isArray && property.items?.$ref) {
      const ref = isArray ? property.items?.$ref : property.$ref;
      // Example $ref: #/components/schemas/EdgeResponse
      name = ref?.split('/')?.pop() || '';
      definitionLink = `#${name}`;
    } else {
        name = (isArray ? property.items?.type : property.type) || ''
    }

    formattedName = name;

    if (property.enum || isArray && property.items.enum) {
        const enumValues = isArray ? property.items.enum : property.enum;
        formattedName += ` (${enumValues.join(', ')})`
    }

    formattedName = isArray ? `${formattedName}[]` : formattedName;

    return {
        name, definitionLink, formattedName, isArray
    }
  }

  export const parsePropertyDescription = (property) => {
    return property.description;
  }

  export const parsePropertyConstraints = (property, required) => {
    const constraints = [];
    if (required) {
        constraints.push('Required');
    } 
    if (property.minimum !== undefined) {
        constraints.push(`Minimum: ${property.minimum}`)
    }
    if (property.maximum !== undefined) {
        constraints.push(`Maximum: ${property.maximum}`)
    }

    return constraints;
  }