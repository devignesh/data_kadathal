const fs = require('fs');
const puppeteer = require('puppeteer');

(async () => {
  const source = process.argv[2];
  const target = process.argv[3];

  const browser = await puppeteer.launch({
    defaultViewport: {
      height: 1080,
      width: 1920,
    },
    devtools: true,
    headless: false,
  });

  const page = await browser.newPage();

  await page.goto(source, {
    waitUntil: ['networkidle0'],
  });

  try {
    for (p = 2; p <= 20; p++) {
      await page.click(`[data-page="${p}"]`);
      await page.waitForSelector(`[data-page="${p + 1}"]`, {
        timeout: 30000,
        visible: true,
      });
    }
  } catch (exception) {
    console.log('for (...) {...}');
    console.log(exception);
  }
  try {
    for (p = 2; p <= 20; p++) {
      await page.click(`[data-page="${p}"]`);
      await page.waitForSelector(`[data-page="${p + 1}"]`, {
        timeout: 30000,
        visible: true,
      });
    }
  } catch (exception) {
    console.log('for (...) {...}');
    console.log(exception);
  }

