import React from 'react';
import { parseModel } from './open-api-model-parser';

const OpenApiModels = ({ modelName, modelFilter, recursive = true }) => {

  const models = React.useMemo(() => {
    if (!modelName && !modelFilter) {
      return [];
    }
    return parseModel({modelName, modelFilter, recursive});
  }, [modelName, modelFilter]);

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
                      {p.type.definitionLink ? <a href={p.type.definitionLink}><code>{p.type.formattedName}</code></a> : <code>{p.type.formattedName}</code>}
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

export default OpenApiModels;
