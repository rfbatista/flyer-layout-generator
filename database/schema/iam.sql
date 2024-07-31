CREATE TABLE companies (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  enabled bool,
  api_key text,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TYPE ROLES AS ENUM ('admin', 'colab', 'gm');

CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  email text,
  role ROLES default 'colab',
  company_id int,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (company_id) REFERENCES companies (id) ON UPDATE CASCADE ON DELETE CASCADE
);
