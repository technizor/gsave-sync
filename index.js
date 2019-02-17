const loadConfig = require('./lib/loadConfig');
const gameSync = require('./lib/gameSync');

async function run(configPath) {
    const config = await loadConfig(configPath);
    await gameSync(config);
}

run('./example-config.yaml');