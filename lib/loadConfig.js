const os = require('os');
const path = require('path');
const yaml = require('js-yaml');
const fs = require('fs-extra');

function resolveHome(filePath) {
    if (filePath[0] === '~') {
        return path.join(os.homedir(), filePath.slice(1));
    }
}
async function loadConfig(configFile) {
    try {
        let defaultConfig = yaml.safeLoad(fs.readFileSync(path.resolve(__dirname, 'default.yml')));
        let config = yaml.safeLoad(fs.readFileSync(configFile, 'utf8'));

        config.settings = Object.assign({}, defaultConfig.settings, config.settings);

        for (let gameName in config.games) {
            let origPath = config.games[gameName].path;
            let resPath = resolveHome(origPath);
            if (resPath !== origPath) {
                console.log(`${gameName}: ${origPath} -> ${resPath}`);
                config.games[gameName].path = resPath;
            }
        }
        return config;
    }
    catch (e) {
        console.error(e);
    }
}

module.exports = loadConfig;
// Available:
// - 'aix'
// - 'darwin'
// - 'freebsd'
// - 'linux'
// - 'openbsd'
// - 'sunos'
// - 'win32'