const fs = require('fs');
const puppeteer = require('puppeteer');

(async () => {
  const source_url = process.argv[2];
  const target_file = process.argv[3];

  const browser = await puppeteer.launch({
    defaultViewport: {
      height: 1080,
      width: 1920,
    },
    devtools: true,
    headless: false,
  });

  const page = await browser.newPage();

  await page.goto(source_url, {
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
  const companies = await page.evaluate(() => {
    return Array.from(document.querySelectorAll('.startup')).map(
      (value, key) => {
        const record = {
          company: {
            url: '',
            name: '',
            pitch: '',
          },
          joined: '',
          location: [],
          market: [],
          website: '',
          company_size: '',
          stage: '',
          raised: 0.0,
        };
