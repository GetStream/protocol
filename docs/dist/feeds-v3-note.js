// The Feeds v3 spec is generated into openapi/v2/ because v2 is the version of
// the spec generator, not the Feeds product. Show a note above the spec title
// so readers aren't misled by the path.
window.FeedsV3NotePlugin = function () {
  return {
    wrapComponents: {
      info: (Original, system) => (props) => {
        const React = system.React;
        const original = React.createElement(Original, props);
        if (!/\/openapi\/v2\/feeds-/.test(system.specSelectors.url())) {
          return original;
        }
        return React.createElement(
          'div',
          null,
          React.createElement(
            'div',
            { className: 'feeds-v3-note' },
            React.createElement('strong', null, 'Feeds v3 API. '),
            'This spec lives under ',
            React.createElement('code', null, 'openapi/v2/'),
            ' because v2 is the version of the OpenAPI spec, not the Feeds ' +
              'product version.',
          ),
          original,
        );
      },
    },
  };
};
