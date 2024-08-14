CREATE TABLE companies (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  enabled bool,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE companies_api_credentials (
  id BIGSERIAL PRIMARY KEY,
  name text,
  api_key text NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  company_id int,
  FOREIGN KEY (company_id) REFERENCES companies (id) ON UPDATE CASCADE ON DELETE CASCADE
);
