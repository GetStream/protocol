import React from 'react';
import OpenApiModels from './OpenApiModels';
import apiJson from '../video-client-openapi.json';

const filter = (apiJson) =>
  Object.keys(apiJson.components.schemas).filter(
    (key) =>
      apiJson.components.schemas[key]['x-stream-event-call-type'] === true,
  );

const events = filter(apiJson).map((key) => {
  const type = apiJson.components.schemas[key].properties.type.default || '-';
  const description = apiJson.components.schemas[key].description || '-';

  return { key, type, description };
});

events.sort((e1, e2) => (e1.type < e2.type ? -1 : e1.type > e2.type ? 1 : 0));

const CallEventModels = () => {
  return (
    <React.Fragment>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Description</th>
          </tr>
        </thead>
        <tbody>
          {events.map((event) => (
            <tr key={event.type}>
              <td>
                <a href={'#' + event.key}>
                  <code>{event.type}</code>
                </a>
              </td>
              <td>{event.description}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <OpenApiModels
        modelFilter={filter}
        recursive={true}
        apiJson={apiJson}
      ></OpenApiModels>
    </React.Fragment>
  );
};

export default CallEventModels;
