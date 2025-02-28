window.onload = function () {
  window.ui = SwaggerUIBundle({
    urls: [
      {
        name: 'Chat',
        url: 'https://raw.githubusercontent.com/GetStream/protocol/main/openapi/chat-openapi.yaml',
      },
      {
        name: 'Video',
        url: 'https://raw.githubusercontent.com/GetStream/protocol/main/openapi/video-openapi.yaml',
      },
      {
        name: 'Moderation',
        url: 'https://raw.githubusercontent.com/GetStream/protocol/main/openapi/moderation-serverside-api.yaml',
      },
    ],
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
    plugins: [SwaggerUIBundle.plugins.DownloadUrl],
    layout: 'StandaloneLayout',
  });
};
