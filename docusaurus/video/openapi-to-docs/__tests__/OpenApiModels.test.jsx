import { render, screen } from './utils'
import OpenApiModels from '../../docusaurus/docs/api/_common_/OpenApiModels'
import mockApiJson from './mock-open-api.json';
import { beforeEach, describe, expect, it } from 'vitest';

describe('OpenApiModels', async () => {
  let container;

  beforeEach(() => {
    container = render(
      <OpenApiModels modelName={'GetOrCreateCallRequest'} apiJson={mockApiJson}/>,
    ).container;
  });

  it('should render each model in an HTML table', () => {
    expect(screen.getByTestId('GetOrCreateCallRequest-table')).toBeInTheDocument();
    expect(screen.getByTestId('CallRequest-table')).toBeInTheDocument();
    expect(screen.getByTestId('MemberRequest-table')).toBeInTheDocument();

    expect(container.querySelectorAll('table').length).toBe(3);
  });

  it('should render model name', () => {
    expect(screen.getByText('GetOrCreateCallRequest')).toBeInTheDocument();
  });

  it('should render properties of model - name', () => {
    expect(screen.getByTestId('GetOrCreateCallRequest-data-name').innerHTML).toContain('data');
  });

  it('should render properties of model - type and type definition links', () => {
    // ref
    expect(screen.getByTestId('GetOrCreateCallRequest-data-type').innerHTML).toContain('CallRequest');

    expect(screen.getByTestId('GetOrCreateCallRequest-data-typelink').href).toContain('#CallRequest');

    expect(container.querySelector('#CallRequest')).toBeInTheDocument();

    // primitive
    expect(screen.getByTestId('GetOrCreateCallRequest-ring-type').innerHTML).toContain('boolean');

    expect(container.querySelector('[data-testid=GetOrCreateCallRequest-ring-typelink]')).toBe(null);

    // formatted name
    expect(screen.getByTestId('CallRequest-members-type').innerHTML).toContain('MemberRequest[]');
  });

  it('should render properties of model - description', () => {
    expect(screen.getByTestId('GetOrCreateCallRequest-data-description').innerHTML).toContain('Configuration options for the call');
  })

  it('should render properties of model - constraints', () => {
    expect(screen.getByTestId('GetOrCreateCallRequest-data-constraints').innerHTML).toContain('Required');
    expect(screen.getByTestId('GetOrCreateCallRequest-notify-constraints').innerHTML).toContain('-');
    expect(screen.getByTestId('CallRequest-members-constraints').innerHTML).toContain('Required, Maximum: 100');
  });
})