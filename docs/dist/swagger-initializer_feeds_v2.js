window.onload = function () {
  window.ui = SwaggerUIBundle({
    url: 'https://raw.githubusercontent.com/GetStream/protocol/main/openapi/v2/feeds-serverside-api.yaml',
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
    plugins: [SwaggerUIBundle.plugins.DownloadUrl, FeedsV3NotePlugin],
    layout: 'StandaloneLayout',
  });
};
