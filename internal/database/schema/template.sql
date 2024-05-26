CREATE TYPE TEMPLATE_TYPE AS ENUM (
  'slots',
  'distortion'
);


CREATE TABLE templates (
  id   SERIAL PRIMARY KEY,
  name text      NOT NULL,
  type TEMPLATE_TYPE,
  request_id TEXT,
  width INT,
  height INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);


CREATE TABLE templates_slots (
  id   SERIAL PRIMARY KEY,
  xi INT,
  yi INT,
  width INT,
  height INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  template_id INT NOT NULL,

  FOREIGN KEY (template_id) REFERENCES templates (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE templates_distortions (
  id   SERIAL PRIMARY KEY,
  x INT,
  y INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  template_id INT NOT NULL,

  FOREIGN KEY (template_id) REFERENCES templates (id) ON DELETE CASCADE ON UPDATE CASCADE
);
