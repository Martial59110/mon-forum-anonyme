import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { vi } from 'vitest';

import MessageForm from './MessageForm.jsx';

describe('MessageForm', () => {
  test('shows validation error when pseudonym is missing', async () => {
    const user = userEvent.setup();
    render(<MessageForm onMessagePosted={() => {}} />);

    await user.click(screen.getByRole('button', { name: /poster/i }));

    expect(
      await screen.findByText(/pseudonyme est requis/i),
    ).toBeInTheDocument();
  });

  test('posts message and calls onMessagePosted', async () => {
    const user = userEvent.setup();
    const onMessagePosted = vi.fn();

    const fetchMock = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({
        id: 1,
        pseudonym: 'Alice',
        content: 'Hello',
        avatar: '',
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString(),
      }),
    });
    vi.stubGlobal('fetch', fetchMock);

    render(<MessageForm onMessagePosted={onMessagePosted} />);

    await user.type(screen.getByLabelText(/pseudonyme/i), 'Alice');
    await user.type(screen.getByLabelText(/message/i), 'Hello');

    await user.click(screen.getByRole('button', { name: /poster/i }));

    expect(fetchMock).toHaveBeenCalledTimes(1);
    expect(onMessagePosted).toHaveBeenCalledTimes(1);

    vi.unstubAllGlobals();
  });
});
