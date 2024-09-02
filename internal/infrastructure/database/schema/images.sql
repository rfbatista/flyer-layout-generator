CREATE TABLE images (
  id   BIGSERIAL PRIMARY KEY,
  url text      NOT NULL,
  photoshop_id INT,
  template_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
