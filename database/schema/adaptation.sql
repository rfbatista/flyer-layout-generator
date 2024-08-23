CREATE TYPE ADAPTATION_BATCH_STATUS AS ENUM (
  'pending',
  'started',
  'finished',
  'error',
  'closed',
  'unknown'
);

CREATE TABLE adaptation_batch (
  id   BIGSERIAL PRIMARY KEY,
  layout_id INT,
  design_id INT,
  request_id INT,
  user_id INT,
  template_id INT,
  status ADAPTATION_BATCH_STATUS,
  started_at TIMESTAMP,
  finished_at TIMESTAMP,
  error_at TIMESTAMP,
  stopped_at TIMESTAMP,
  updated_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  config TEXT,
  log TEXT,

  FOREIGN KEY (template_id) REFERENCES templates (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (request_id) REFERENCES layout_requests (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (design_id) REFERENCES design (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE adaptation_batch_job (
  id   BIGSERIAL PRIMARY KEY,
  layout_id INT,
  design_id INT,
  request_id INT,
  template_id INT,
  status TEXT,
  image_url TEXT,
  started_at TIMESTAMP,
  finished_at TIMESTAMP,
  error_at TIMESTAMP,
  stopped_at TIMESTAMP,
  updated_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  config TEXT,
  log TEXT,

  FOREIGN KEY (template_id) REFERENCES templates (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (request_id) REFERENCES layout_requests (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (design_id) REFERENCES design (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);
