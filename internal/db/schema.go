package db

var schema string = `
CREATE TABLE IF NOT EXISTS transaction (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    date DATE NOT NULL,
    description TEXT NOT NULL,
    amount NUMERIC(10,2) NOT NULL,
    category TEXT NOT NULL,
    account TEXT NOT NULL
);
`
