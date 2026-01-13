import { defineConfig } from '@playwright/test';

export default defineConfig({
  testDir: './tests',
  timeout: 60_000,
  expect: { timeout: 10_000 },
  retries: process.env.CI ? 1 : 0,
  use: {
    baseURL: process.env.E2E_BASE_URL || 'http://localhost',
    headless: true,
    trace: 'retain-on-failure',
  },
  reporter: process.env.CI ? [['github'], ['list']] : [['list']],
});
