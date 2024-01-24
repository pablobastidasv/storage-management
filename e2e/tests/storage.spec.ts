import { test, expect } from '@playwright/test';

test('Adding an items to a storage',
    async ({ page }) => {
        await page.goto('http://127.0.0.1:8080/')

        const info = await page.getByText(/Bronce - \d+ Kilogramos/).textContent()
        const actualAmount = info?.split(' ')[2] ?? '0'
        const expectedAmount = parseInt(actualAmount) + 2
        await page.getByRole('button', { name: 'Ingresar Producto' }).click()
        await page.getByLabel('Producto').selectOption({ label: 'Bronce (Kilogramos)' })
        await page.getByLabel('Cantidad:').fill('2')

        await page.getByRole('button', { name: 'Ingresar', exact: true }).click()

        await expect(page.getByText(`Bronce - ${expectedAmount} Kilogramos`)).toBeVisible()


    });


