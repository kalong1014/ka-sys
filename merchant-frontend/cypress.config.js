const { defineConfig } = require('cypress')

module.exports = defineConfig({
  e2e: {
    setupNodeEvents(on, config) {
      // 配置测试环境
    },
    baseUrl: 'http://localhost:3000',
    specPattern: 'cypress/e2e/**/*.spec.{js,jsx,ts,tsx}',
    supportFile: 'cypress/support/e2e.js'
  }
})