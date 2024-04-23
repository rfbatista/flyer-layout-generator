CREATE TABLE template (
  id   SERIAL PRIMARY KEY,
  name text      NOT NULL,
  width INT,
  height INT,
  slots_x INT,
  slots_y INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);
