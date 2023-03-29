window.onload = function() {
  window.ui = SwaggerUIBundle({
    urls: [
      {
        name: "Chat API",
        url:"https://raw.githubusercontent.com/GetStream/protocol/main/openapi/chat-openapi.yaml"
      },
      {
        name: "Video API",
        url:"https://raw.githubusercontent.com/GetStream/protocol/main/openapi/video-openapi.yaml"
      }
    ],
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl,
    ],
    layout: "StandaloneLayout"
  });
};
