import React from 'react';
import { parseModel } from './open-api-model-parser';

const OpenApiModels = ({ modelName, modelFilter, recursive = true, apiJson }) => {

  const models = React.useMemo(() => {
    if (!modelName && !modelFilter) {
      return [];
    }
    return parseModel({modelName, modelFilter, recursive, apiJson});
  }, [modelName, modelFilter]);

  return (
    <div>
      {models.map((model) => (
        <React.Fragment key={model.name}>
          <h2 id={model.name}>{model.name}</h2>
          <table data-testid={model.name + '-table'}>
            <thead>
              <tr>
                <th>Name</th>
                <th>Type</th>
                <th>Description</th>
                <th>Constraints</th>
              </tr>
            </thead>
            <tbody>
              {model.properties.map(p => {
                return (
                  <tr key={model.name + p.name}>
                    <td data-testid={model.name + '-' + p.name + '-name'}>
                      <code>{p.name}</code>
                    </td>
                    <td data-testid={model.name + '-' + p.name + '-type'}>
                      {p.type.definitionLink ? <a data-testid={model.name + '-' + p.name + '-typelink'} href={p.type.definitionLink}><code>{p.type.formattedName}</code></a> : <code>{p.type.formattedName}</code>}
                    </td>
                    <td data-testid={model.name + '-' + p.name + '-description'}>{p.description || '-'}</td>
                    <td data-testid={model.name + '-' + p.name + '-constraints'}>{p.constraints.join(', ') || '-'}</td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </React.Fragment>
      ))}
    </div>
  );
};

export default OpenApiModels;
