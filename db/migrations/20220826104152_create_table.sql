-- migrate:up
CREATE TABLE IF NOT EXISTS requests_counters (
  int1 INTEGER,
  int2 INTEGER,
  lim INTEGER,
  str1 TEXT,
  str2 TEXT,
  counter INTEGER,
  PRIMARY KEY(int1,int2,lim,str1,str2)
);

CREATE INDEX requests_counters_encoded_params_idx ON requests_counters (counter DESC);


-- migrate:down
DROP TABLE requests_counters;
