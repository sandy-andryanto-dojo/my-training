const db = require("../models");
const User = db.users;
const Op = db.Sequelize.Op;
const config = require('../../config.json');
const jwt = require('jsonwebtoken');

var bcrypt = require('bcryptjs');


module.exports = {
  login,
  getAll
};

async function login(req, res) {
 
  if (!req.body.email) {
    res.status(400).send({
      message: "Email can not be empty!"
    });
    return;
  }

  if (!req.body.password) {
      res.status(400).send({
        message: "Password can not be empty!"
      });
      return;
  }

  let email = req.body.email;
  let password = req.body.password;

  User.findOne({ where: { email: email } })
    .then(data => {
      if (data) {
        let token = jwt.sign({ sub: data.id }, config.secret, { expiresIn: '7d' });
        let password_saved = data.get('password');
        if(!bcrypt.compareSync(password, password_saved)){
          res.status(404).send({
            message: "These credentials do not match our records."
          });
        }else{
          res.send({
            ...omitPassword(data),
            token
          }); 
        }
      } else {
        res.status(404).send({
          message: "These credentials do not match our records."
        });
      }
    })
    .catch(err => {
      console.log(err);
      res.status(404).send({
        message: "These credentials do not match our records."
      });
    });

}

async function getAll() {
  let auth_user = users.map(u => omitPassword(u));
}

// helper functions

function omitPassword(user) {
  const { password, ...userWithoutPassword } = user;
  let result = userWithoutPassword;
}