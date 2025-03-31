window.onload = function () {
  window.ui = SwaggerUIBundle({
    url: 'https://raw.githubusercontent.com/GetStream/protocol/main/openapi/v2/moderation-serverside-api.yaml',
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
    plugins: [SwaggerUIBundle.plugins.DownloadUrl],
    layout: 'StandaloneLayout',
  });
};
