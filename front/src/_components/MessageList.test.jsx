import { render, screen } from '@testing-library/react';
import { vi } from 'vitest';

import MessageList from './MessageList.jsx';

describe('MessageList', () => {
  test('renders empty state when no messages', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => [],
    });
    vi.stubGlobal('fetch', fetchMock);

    render(<MessageList />);

    expect(screen.getByText(/chargement/i)).toBeInTheDocument();
    expect(await screen.findByText(/aucun message/i)).toBeInTheDocument();

    vi.unstubAllGlobals();
  });

  test('renders error when fetch fails', async () => {
    const fetchMock = vi.fn().mockRejectedValue(new Error('Failed to fetch'));
    vi.stubGlobal('fetch', fetchMock);

    render(<MessageList />);

    expect(
      await screen.findByText(/erreur lors du chargement/i),
    ).toBeInTheDocument();

    vi.unstubAllGlobals();
  });
});
