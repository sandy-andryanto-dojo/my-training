const config = require('../../config.json');

module.exports = {
    HOST: config["DB_HOST"],
    USER: config["DB_USERNAME"],
    PASSWORD: config["DB_PASSWORD"],
    DB: config["DB_DATABASE"],
    dialect: config["DB_CONNECTION"],
    pool: {
      max: 5,
      min: 0,
      acquire: 30000,
      idle: 10000
    }
  };