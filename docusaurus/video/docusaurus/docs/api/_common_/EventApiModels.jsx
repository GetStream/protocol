import React from 'react';
import OpenApiModels from './OpenApiModels';
import apiJson from '../video-openapi.json';

const filter = (apiJson) =>
Object.keys(apiJson.components.schemas).filter(
  (key) =>
    apiJson.components.schemas[key]['x-stream-event-call-type'] === true,
);

const events = filter(apiJson).map(key => {
    const type = apiJson.components.schemas[key].properties.type.default || '-';
    const description = apiJson.components.schemas[key].description || '-';

    return {key, type, description}
});

events.sort((e1, e2) => e1.type < e2.type ? -1 : (e1.type > e2.type ? 1 : 0));

const EventApiModels = () => {
    return <React.Fragment>
        <table>
            <thead>
                <th>Name</th>
                <th>Description</th>
            </thead>
            {events.map(event => <tr>
                <td><a href={'#' + event.key}><code>{event.type}</code></a></td>
                <td>{event.description}</td>
            </tr>)}
        </table>
        <OpenApiModels modelFilter={filter} recursive={true}></OpenApiModels>
    </React.Fragment>
}

export default EventApiModels;