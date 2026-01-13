import { test, expect } from '@playwright/test';

test('poster un message puis le voir dans la liste', async ({ page }) => {
  await page.goto('/');

  const pseudonym = `E2E-${Date.now()}`;
  const content = `Hello from E2E ${Date.now()}`;

  await page.getByLabel(/pseudonyme/i).fill(pseudonym);
  await page.getByLabel(/^💬\s*message/i).fill(content);

  await page.getByRole('button', { name: /poster/i }).click();

  await expect(page.getByText(/message posté avec succès/i)).toBeVisible();

  // Le MessageList se refresh après POST, on attend l’apparition du contenu
  await expect(page.getByText(content)).toBeVisible();
});
