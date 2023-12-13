import { describe, expect, it } from 'vitest';
import { parseModel } from '../../docusaurus/docs/api/_common_/open-api-model-parser';
import mockApiJson from './mock-open-api.json';

describe('Parse model properties', () => {

  it('display a single model', () => {
    const result = parseModel({modelName: 'GetOrCreateCallRequest', recursive: false, apiJson: mockApiJson});

    expect(result.length).toBe(1);

    const model = result[0];

    expect(model.name).toBe('GetOrCreateCallRequest');
    expect(model.properties.length).toBe(4);

    const property = model.properties[0];

    expect(property.name).toBe('data');
    expect(property.type).toBeDefined();
    expect(property.type.formattedName).toBe('CallRequest');
    expect(property.description).toBe('Configuration options for the call');
    expect(property.constraints).toEqual(['Required']);
  });

  it('display a single model recursively', () => {
    const result = parseModel({modelName: 'GetOrCreateCallRequest', recursive: true, apiJson: mockApiJson});

    expect(result.length).toBe(3);
  });

  it('display multiple models', () => {
    const result = parseModel({modelFilter: (apiJson) => Object.keys(apiJson.components.schemas).filter(modelName => modelName.toLowerCase().includes('call')), recursive: false, apiJson: mockApiJson});

    expect(result.length).toBe(3);
  });

  it('display multiple models recursively', () => {
    const result = parseModel({modelFilter: (apiJson) => Object.keys(apiJson.components.schemas).filter(modelName => modelName.toLowerCase().includes('call')), recursive: true, apiJson: mockApiJson});

    expect(result.length).toBe(4);
  });

  it(`shouldn't display property definition links for non-recursive models`, () => {
    const result = parseModel({modelName: 'GetOrCreateCallRequest', recursive: false, apiJson: mockApiJson});

    const dataProperty = result[0].properties.find(p => p.name === 'data');

    expect(dataProperty.type.name).toBe('CallRequest');
    expect(dataProperty.type.definitionLink).toBeUndefined();
  });

  it(`shouldn't display property definition links for top-level enums`, () => {
    const result = parseModel({modelName: 'UpdateCallResponse', recursive: false, apiJson: mockApiJson});

    const ownCapabilitiesProperty = result[0].properties.find(p => p.name === 'own_capabilities');

    expect(ownCapabilitiesProperty.type.definitionLink).toBeUndefined();
  });

  it('should set required flag properly', () => {
    const apiJson = {
        components: {
            schemas: {
                MyModel: {
                    properties: {
                        a: {},
                        b: {}
                    },
                    required: ['a']
                }
            }
        }
    }

    const result = parseModel({modelName: 'MyModel', apiJson});

    const aProperty = result[0].properties.find(p => p.name === 'a');
    const bProperty = result[0].properties.find(p => p.name === 'b');
    expect(aProperty.constraints).toEqual(['Required']);
    expect(bProperty.constraints).toEqual([]);
  })
});
