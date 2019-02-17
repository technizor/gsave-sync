const os = require('os');
const path = require('path');
const fs = require('fs-extra');

const platform = os.platform();
const homedir = os.homedir();

async function gameSync(config, dryRun = false) {
    let timestamp = new Date().toISOString();
    for (let gameName in config.games) {
        await syncOne(config.settings, config.games[gameName], timestamp, dryRun);
    }
}
async function syncOne(settings, gameSettings, timestamp, dryRun = false) {
    let groupName = gameSettings.shortName || gameSettings.name;
    let itemName = settings.prefix;
    if (settings.timestamp) {
        let timestampPath = timestamp.substring(0, timestamp.indexOf('.')).replace(/:/g, '-');
        itemName = `${settings.prefix}-${timestampPath}`;
    }
    let sourceLocation = path.resolve(gameSettings.path);
    let saveLocation = path.resolve(settings.savePath, groupName, itemName);
    console.log(sourceLocation, '->', saveLocation);
    if (await fs.exists(sourceLocation)) {
        if (!dryRun) {
            await fs.mkdirp(saveLocation);
            await fs.copy(sourceLocation, saveLocation);
        }
    }
}

module.exports = gameSync;