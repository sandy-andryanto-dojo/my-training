const config = require('./config.json');
var bcrypt = require('bcryptjs');
var mysql = require('mysql2');
var faker = require('faker');
var conn = mysql.createConnection({
    host: config["DB_HOST"],
    user: config["DB_USERNAME"],
    password: config["DB_PASSWORD"],
    database: config["DB_DATABASE"],
    multipleStatements: true
});

function formatDate(date) {
    var d = new Date(date),
        month = '' + (d.getMonth() + 1),
        day = '' + d.getDate(),
        year = d.getFullYear();

    if (month.length < 2)
        month = '0' + month;
    if (day.length < 2)
        day = '0' + day;

    return [year, month, day].join('-');
}

conn.connect((err) => {

    if (err) throw err;

    // to do here
    var password = bcrypt.hashSync("password", 10);
    for(i = 1; i < 100; i++){
        let user = {
            username: faker.internet.userName().toLowerCase(),
            email: faker.internet.email().toLowerCase(),
            phone: faker.phone.phoneNumber().toLowerCase(),
            password: password,
            is_banned: 0,
            createdAt: formatDate(new Date().toLocaleString()),
            updatedAt: formatDate(new Date().toLocaleString()),
        };
        let sql = "INSERT INTO users SET ?"
        conn.query(sql, user, (err, results) => {
            console.log(user);
        });
        console.log('Press any key to exit');
    }



    process.stdin.setRawMode(true);
    process.stdin.resume();
    process.stdin.on('data', process.exit.bind(process, 0));

});