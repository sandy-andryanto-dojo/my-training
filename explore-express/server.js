const express = require("express");
// const bodyParser = require("body-parser"); /* deprecated */
const cors = require("cors");

const bodyParser = require('body-parser');

const jwt = require('./app/helpers/jwt');
const errorHandler = require('./app/helpers/error-handler');
const auditLogMiddleware = require('@ozawa/express-audit-log-middleware');

const app = express();

var corsOptions = {
  origin: "http://localhost:8081"
};

app.use(auditLogMiddleware({
  routes: [
    {
      methods: [
        'POST',
        'PUT',
        'DELETE',
      ],
      uris: [
        '**',
      ],
    },
  ],
  extractorFunctions: {
    actor        : (req) => req.auth,
    ipAddress    : (req) => req.ip,
    correlationID: (req) => req.headers['x-correlation-id'] || undefined,
  },
  filename: `./audit_log.json`,
}));

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());
app.use(cors());

app.use(jwt());

app.use(cors(corsOptions));

// parse requests of content-type - application/json
app.use(express.json());  /* bodyParser.json() is deprecated */

// parse requests of content-type - application/x-www-form-urlencoded
app.use(express.urlencoded({ extended: true }));   /* bodyParser.urlencoded() is deprecated */

const db = require("./app/models");

db.sequelize.sync();
// // drop the table if it already exists
// db.sequelize.sync({ force: true }).then(() => {
//   console.log("Drop and re-sync db.");
// });

// simple route
app.get("/", (req, res) => {
  res.sendStatus(404);
});

require("./app/routes/auth.routes")(app);
require("./app/routes/turorial.routes")(app);


// set port, listen for requests
const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}.`);
});
