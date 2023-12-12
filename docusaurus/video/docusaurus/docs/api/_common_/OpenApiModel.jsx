import React from 'react';
import apiJson from '../video-openapi.json';

const OpenApiModel = ({ modelName, recursive = true }) => {

  const models = React.useMemo(() => {
    const modelsResult = [{name: modelName, properties: []}];
  
    for (let i = 0; i < modelsResult.length; i++) {
      const model = modelsResult[i];
  
      const schemaJson = apiJson.components.schemas[model.name];
      const propertiesJson = schemaJson.properties;

      model.properties = Object.keys(propertiesJson).map(key => {
        const property = propertiesJson[key];
        
        // Parse type info
        let type;
        let typeHref;
        let displayType;
        let isArray = property.type === 'array';
        if (property.$ref || isArray && property.items?.$ref) {
          const ref = isArray ? property.items?.$ref : property.$ref;
          // Example $ref: #/components/schemas/EdgeResponse
          type = ref?.split('/')?.pop() || '';
          typeHref = `#${type}`
          if (recursive && apiJson.components.schemas[type] && !modelsResult.find(r => r.name === type)) {
            modelsResult.push({name: type, properties: []});
          }
        } else {
          type = (isArray ? property.items?.type : property.type) || ''
        }
        displayType = isArray ? `${type}[]` : type;
        if (property.enum) {
          displayType += ` (${property.enum.join(', ')})`
        }

        // Parse title + description
        let description = property.title;
        if (property.description) {
          description += ` - ${property.description}`;
        }

        // Parse constraints
        const constraints = [];
        if (schemaJson.required?.includes(key)) {
          constraints.push('Required');
        } 
        if (property.minimum !== undefined) {
          constraints.push(`Minimum: ${property.minimum}`)
        }
        if (property.maximum !== undefined) {
          constraints.push(`Maximum: ${property.maximum}`)
        }

        return {
          name: key,
          type,
          displayType,
          typeHref,
          constraints,
          description
        }
      });
    }

    return modelsResult;
  }, [modelName]);

  return (
    <div>
      {models.map((model) => (
        <React.Fragment>
          <h2 id={model.name}>{model.name}</h2>
          <table>
            <thead>
              <th>Name</th>
              <th>Type</th>
              <th>Description</th>
              <th>Constraints</th>
            </thead>
            {model.properties.map(p => {
              return (
                <React.Fragment>
                  <tr>
                    <td>
                      <code>{p.name}</code>
                    </td>
                    <td>
                      {p.typeHref ? <a href={p.typeHref}><code>{p.displayType}</code></a> : <code>{p.displayType}</code>}
                    </td>
                    <td>{p.description || '-'}</td>
                    <td>{p.constraints.join(', ') || '-'}</td>
                  </tr>
                </React.Fragment>
              );
            })}
          </table>
        </React.Fragment>
      ))}
    </div>
  );
};

export default OpenApiModel;
