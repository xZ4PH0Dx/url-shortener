package pg

const Schema = `
CREATE TABLE IF NOT EXISTS urls (
  id INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  code VARCHAR(32) NOT NULL,
  original_url VARCHAR(512) NOT NULL,
  UNIQUE (code)
);
`
