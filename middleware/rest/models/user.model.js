const mongoose = require("mongoose");

const User = mongoose.model(
    "User",
    new mongoose.Schema({
        // С.А. Кантор
        member: String,
        // Administration
        group: String,
        // using during auth
        email: String,
        password: String
    })
);

module.exports = User;