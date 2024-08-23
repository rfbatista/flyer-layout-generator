CREATE TYPE TEMPLATE_TYPE AS ENUM (
  'slots',
  'distortion',
  'adaptation'
);

CREATE TABLE templates (
  id   SERIAL PRIMARY KEY,
  name text      NOT NULL,
  request_id TEXT,
  project_id INT,
  width INT,
  height INT,
  slots_x INT,
  slots_y INT,
  max_slots_x INT,
  max_slots_y INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  company_id int,

  FOREIGN KEY (company_id) REFERENCES companies (id) ON UPDATE CASCADE ON DELETE CASCADE,

  FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE ON UPDATE CASCADE
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

