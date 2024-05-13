import { beforeEach, describe, expect, it } from 'vitest';
import {
  parsePropertyConstraints,
  parsePropertyDescription,
  parsePropertyType,
} from '../../docusaurus/docs/api/_common_/open-api-model-parser';

describe('Parse model properties', () => {
  let property;

  beforeEach(() => {
    property = {
      description: 'Response HTTP status code',
      format: 'int32',
      title: 'Status code',
      type: 'integer',
      'x-stream-index': '004',
    };
  });

  describe('type', () => {
    it('primitive type', () => {
      property.type = 'string';

      expect(parsePropertyType(property)).toEqual({
        name: 'string',
        formattedName: 'string',
        definitionLink: undefined,
        isArray: false
      });
    });

    it('enum', () => {
      property.type = 'string';
      property.enum = [
        "speaker",
        "earpiece"
      ],

      expect(parsePropertyType(property)).toEqual({
        name: 'string',
        formattedName: 'string (speaker, earpiece)',
        definitionLink: undefined,
        isArray: false
      });
    });

    it('array of enums', () => {
      property.type = 'array';
      property.items = {
        type: 'string',
        enum: [
          "speaker",
          "earpiece"
        ]
      };

      expect(parsePropertyType(property)).toEqual({
        name: 'string',
        formattedName: 'string (speaker, earpiece)[]',
        definitionLink: undefined,
        isArray: true
      });
    });

    it('array of primitive types', () => {
      property.type = 'array';
      property.items = {
        "type": "string"
      };

      expect(parsePropertyType(property)).toEqual({
        name: 'string',
        formattedName: 'string[]',
        definitionLink: undefined,
        isArray: true
      });
    });

    it('reference type', () => {
      property.type = undefined;
      property.$ref = '#/components/schemas/CallResponse';

      expect(parsePropertyType(property)).toEqual({
        name: 'CallResponse',
        formattedName: 'CallResponse',
        definitionLink: '#CallResponse',
        isArray: false
      });
    });

    it('array of reference types', () => {
      property.type = 'array';
      property.items = {
        $ref: "#/components/schemas/MemberResponse"
      };

      expect(parsePropertyType(property)).toEqual({
        name: 'MemberResponse',
        formattedName: 'MemberResponse[]',
        definitionLink: '#MemberResponse',
        isArray: true
      });
    });
  })

  it('description', () => {
    property.description = undefined;

    expect(parsePropertyDescription(property)).toBe(undefined);

    property.description = 'Response HTTP status code';

    expect(parsePropertyDescription(property)).toBe(
      'Response HTTP status code',
    );
  });

  describe('constraints', () => {
    it('no constraints', () => {
      expect(parsePropertyConstraints(property)).toEqual([]);
    });

    it('required', () => {
      expect(parsePropertyConstraints(property, true)).toEqual(['Required'])
    });

    it('maximum', () => {
      property.maximum = 5;

      expect(parsePropertyConstraints(property)).toEqual(['Maximum: 5'])
    });

    it('minimum', () => {
      property.minimum = 1;

      expect(parsePropertyConstraints(property)).toEqual(['Minimum: 1'])
    });

    it('multiple constraints', () => {
      property.minimum = 1;

      expect(parsePropertyConstraints(property, true)).toEqual(['Required', 'Minimum: 1'])
    });
  })
});
