CREATE TABLE IF NOT EXISTS url (
  shorturl      VARCHAR(6) NOT NULL,
  longurl       VARCHAR(255) NOT NULL,
  redirects     INT NOT NULL,
  -- use natural key instead of surrogate key
  PRIMARY KEY (shorturl)
);