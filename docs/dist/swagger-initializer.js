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
        name: 'Chat v2',
        url: 'https://raw.githubusercontent.com/GetStream/protocol/main/openapi/v2/chat-serverside-api.yaml',
      },
      {
        name: 'Video v2',
        url: 'https://raw.githubusercontent.com/GetStream/protocol/main/openapi/v2/video-serverside-api.yaml',
      },
      {
        name: 'Moderation v2',
        url: 'https://raw.githubusercontent.com/GetStream/protocol/main/openapi/v2/moderation-serverside-api.yaml',
      },
    ],
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
    plugins: [SwaggerUIBundle.plugins.DownloadUrl],
    layout: 'StandaloneLayout',
  });
};
