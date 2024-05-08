import { test, expect } from '@playwright/test';
import { randomString } from './utils'

test('creating a new product', async ({ page }) => {
    await page.goto('http://127.0.0.1:8080/admin/products');

    await expect(page).toHaveTitle(/Inventarios - Bastriguez/);

    const name = randomString(10)
    const presentation = "Kilogramos"
    const amountOfPresentations = await page.locator("td", { hasText: presentation }).count()

    await page.getByRole('button', { name: "Crear" }).click()
    await page.getByLabel('Nombre:').fill(name)
    await page.getByLabel('Presentación:').selectOption({ label: presentation })

    await page.getByRole('button', { name: "Guardar" }).click()

    await expect(page.getByText(name)).toBeVisible()
    await expect(page.getByText(presentation)).toHaveCount(amountOfPresentations + 1)
});


test('creating a new client', async ({page}) => {
    await page.goto('http://127.0.0.1:8080/admin/clients');

    const docType = "NIT"
    const docNumber = randomString(10)
    const name = randomString(10)

    await page.getByRole('button', {name: "Crear"}).click()
    await page.getByLabel('Tipo de documento:').selectOption(docType)
    await page.getByLabel('Número de identificación:').fill(docNumber)
    await page.getByLabel("Nombre:").fill(name)

    await page.getByRole("button", {name:'Guardar'}).click()

    await expect(page.getByText("Cliente creado.")).toBeVisible()
})
