-- migrate:up
CREATE TABLE IF NOT EXISTS requests_counters (
  hash TEXT PRIMARY KEY,
  int1 INTEGER,
  int2 INTEGER,
  lim INTEGER,
  str1 TEXT,
  str2 TEXT,
  counter INTEGER
);

CREATE INDEX requests_counters_encoded_params_idx ON requests_counters (counter DESC);


-- migrate:down
DROP TABLE requests_counters;
