import { test as setup, expect } from '@playwright/test';
import dotenv from 'dotenv';

const authFile = 'playwright/.auth/user.json';

dotenv.config();

setup('authenticate', async ({ page }) => {
    await page.goto('http://127.0.0.1:8080/')

    const user = process.env.TEST_USER ?? ''
    const password = process.env.TEST_PASSWORD ?? ''

    console.log(user, password)

    await page.getByLabel('Email address').fill(user)
    await page.getByLabel('Password').fill(password)
    await page.getByRole('button', { name: 'Continue', exact: true }).click()

    await page.waitForURL('http://127.0.0.1:8080/');

    await expect(page.getByText(/Bienvenido/)).toBeVisible();

    await page.context().storageState({ path: authFile })
});


