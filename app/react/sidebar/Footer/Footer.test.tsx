import { render, screen } from '@testing-library/react';
import { vi } from 'vitest';

import { withTestQueryProvider } from '@/react/test-utils/withTestQuery';
import { withUserProvider } from '@/react/test-utils/withUserProvider';

import { SidebarProvider } from '../useSidebarState';

import { Footer } from './Footer';

vi.mock('./UpdateNotifications', () => ({
  UpdateNotification: () => (
    <div data-cy="update-notification">Update Notification</div>
  ),
}));

vi.mock('./BuildInfoModal', () => ({
  BuildInfoModalButton: () => (
    <button data-cy="build-info-modal-button" type="button">
      Build Info
    </button>
  ),
}));

describe('Footer', () => {
  afterEach(() => {
    vi.restoreAllMocks();
  });

  test('should render footer with copyright symbol', () => {
    renderComponent();

    expect(screen.getByText('©')).toBeInTheDocument();
  });

  test('should render OpenDocking text', () => {
    renderComponent();

    expect(screen.getByText('OpenDocking')).toBeInTheDocument();
  });

  test('should render UpdateNotification component', () => {
    renderComponent();

    expect(screen.getByTestId('update-notification')).toBeInTheDocument();
  });

  test('should render BuildInfoModalButton component', () => {
    renderComponent();

    expect(screen.getByTestId('build-info-modal-button')).toBeInTheDocument();
  });

  describe('FooterContent', () => {
    test('should render all child elements in correct order', () => {
      renderComponent();

      const copyrightSymbol = screen.getByText('©');
      const editionText = screen.getByText('OpenDocking');
      const buildInfoButton = screen.getByTestId('build-info-modal-button');

      expect(copyrightSymbol).toBeInTheDocument();
      expect(editionText).toBeInTheDocument();
      expect(buildInfoButton).toBeInTheDocument();
    });
  });
});

function renderComponent() {
  const Wrapped = withTestQueryProvider(
    withUserProvider(SidebarProvider(Footer))
  );

  return render(<Wrapped />);
}
