import { test, expect } from '@playwright/test';
import { randomString } from './utils'

test('creating a new product', async ({ page }) => {
    await page.goto('http://127.0.0.1:8080/admin');

    await expect(page).toHaveTitle(/Inventarios - Bastriguez/);

    const name = randomString(10)
    const presentation = "Kilogramos"

    await page.getByRole('button', { name: "Create" }).click()
    await page.getByLabel('Name:').fill(name)
    await page.getByLabel('Presentación:').selectOption({ label: presentation })

    await page.getByRole('button', { name: "Guardar" }).click()

    await expect(page.getByText(`${name} - ${presentation}`)).toBeVisible()

});

