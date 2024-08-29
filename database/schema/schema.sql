CREATE TYPE ROLES AS ENUM ('admin', 'colab', 'gm');

CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  email text,
  username text,
  role ROLES default 'colab',
  company_id int,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (company_id) REFERENCES companies (id) ON UPDATE CASCADE ON DELETE CASCADE
);
CREATE TABLE advertisers (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  company_id int,

  FOREIGN KEY (company_id) REFERENCES companies (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE clients (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  company_id int,

  FOREIGN KEY (company_id) REFERENCES companies (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE projects (
  id   BIGSERIAL PRIMARY KEY,
  client_id int,
  advertiser_id int,
  briefing text,
  use_ai BOOL DEFAULT FALSE,
  name text      NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  company_id int,

  FOREIGN KEY (company_id) REFERENCES companies (id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (client_id) REFERENCES clients (id) ON UPDATE CASCADE,
  FOREIGN KEY (advertiser_id) REFERENCES advertisers (id) ON UPDATE CASCADE
);


CREATE TABLE design
(
    id              SERIAL PRIMARY KEY,
    name            TEXT NOT NULL,
    image_url       text,
    layout_id       int,
    project_id       int,
    image_extension text,
    file_url        text,
    file_extension  text,
    width           int,
    height          int,
    is_proccessed   bool DEFAULT false,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP,
    company_id int,

    FOREIGN KEY (company_id) REFERENCES companies (id) ON UPDATE CASCADE ON DELETE CASCADE,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON UPDATE CASCADE
);


CREATE TYPE DESIGN_ASSET_TYPE AS ENUM (
  'text', 
  'smartobject',
  'shape',
  'pixel',
  'group',
  'unknown'
);

CREATE TABLE design_assets
(
    id              SERIAL PRIMARY KEY,
    project_id       int,
    design_id       int,
    alternative_to int,
    name            TEXT NOT NULL,
    width           int,
    type            DESIGN_ASSET_TYPE,
    asset_url TEXT,
    asset_path TEXT,
    height          int,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP,

    FOREIGN KEY (project_id) REFERENCES projects (id) ON UPDATE CASCADE,
    FOREIGN KEY (alternative_to) REFERENCES design_assets (id) ON UPDATE CASCADE,
    FOREIGN KEY (design_id) REFERENCES design (id) ON UPDATE CASCADE
);

CREATE TABLE design_assets_properties
(
    id              SERIAL PRIMARY KEY,
    asset_id       int,
    key            TEXT NOT NULL,
    value            TEXT NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP,

    FOREIGN KEY (asset_id) REFERENCES design_assets (id) ON UPDATE CASCADE
);


CREATE TYPE TEMPLATE_TYPE AS ENUM (
  'slots', 
  'distortion', 
  'adaptation', 
  'unknown'
);


CREATE TABLE templates (
  id   SERIAL PRIMARY KEY,
  name text      NOT NULL,
  request_id TEXT,
  project_id INT,
  type TEMPLATE_TYPE,
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

CREATE TABLE layout_requests (
  id   BIGSERIAL PRIMARY KEY,
  layout_id INT,
  design_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  log TEXT,
  config TEXT,
  done INT DEFAULT 0 NOT NULL,
  total INT,
  deleted_at TIMESTAMP,
  updated_at      TIMESTAMP,
  company_id int,

  FOREIGN KEY (company_id) REFERENCES companies (id) ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY (design_id) REFERENCES design (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE layout (
  id   BIGSERIAL PRIMARY KEY,
  design_id INT,
  request_id INT,
  is_original BOOL DEFAULT FALSE,
  image_url TEXT,
  width INT,
  height INT,
  data TEXT,
  stages TEXT[],
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  company_id int,

  FOREIGN KEY (company_id) REFERENCES companies (id) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT fk_layout_design_id FOREIGN KEY (design_id) REFERENCES design (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_layout_request_id FOREIGN KEY (request_id) REFERENCES layout_requests (id) ON UPDATE CASCADE
);


CREATE TABLE layout_requests_jobs (
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


CREATE TYPE ADAPTATION_BATCH_STATUS AS ENUM (
  'pending',
  'started',
  'rendering_images',
  'finished',
  'error',
  'canceled',
  'unknown'
);

CREATE TABLE adaptation_batch (
  id   BIGSERIAL PRIMARY KEY,
  layout_id INT,
  design_id INT,
  request_id INT,
  user_id INT,
  status ADAPTATION_BATCH_STATUS,
  started_at TIMESTAMP,
  finished_at TIMESTAMP,
  error_at TIMESTAMP,
  stopped_at TIMESTAMP,
  updated_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  config TEXT,
  log TEXT,

  foreign key (user_id) references users (id) on delete cascade on update cascade,
  FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TYPE LAYOUT_JOB_STATUS AS ENUM (
  'pending',
  'started',
  'rendering_images',
  'finished',
  'error',
  'closed',
  'unknown',
  'canceled'
);

CREATE TABLE layout_jobs (
  id BIGSERIAL PRIMARY KEY,
  based_on_layout_id INT,
  created_layout_id INT,
  template_id INT,
  user_id INT,
  config TEXT,

  adaptation_batch_id INT,
  replication_batch_id INT,

  status LAYOUT_JOB_STATUS NOT NULL DEFAULT 'pending',
  attempts INT NOT NULL DEFAULT 0,
  started_at TIMESTAMP,
  finished_at TIMESTAMP,
  error_at TIMESTAMP,
  updated_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  log TEXT,
  FOREIGN KEY (adaptation_batch_id) REFERENCES adaptation_batch (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (replication_batch_id) REFERENCES templates (id) ON DELETE CASCADE ON UPDATE CASCADE,
  foreign key (user_id) references users (id) on delete cascade on update cascade,
  FOREIGN KEY (template_id) REFERENCES templates (id) ON DELETE CASCADE ON UPDATE CASCADE
);



CREATE TABLE layout_components
(
    id         SERIAL PRIMARY KEY,
    layout_id INT NOT NULL,
    design_id  INT NOT NULL,
    width      INT,
    height     INT,
    is_original BOOL,
    color      TEXT,
    type       COMPONENT_TYPE,
    xi         INT,
    xii        INT,
    yi         INT,
    yii        INT,
    bbox_xi         INT,
    bbox_xii        INT,
    bbox_yi         INT,
    bbox_yii        INT,
    priority        INT,
    inner_xi              INT,
    inner_xii             INT,
    inner_yi              INT,
    inner_yii             INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (design_id) REFERENCES design (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE layout_elements
(
    id              SERIAL PRIMARY KEY,
    design_id       INT NOT NULL,
    layout_id  INT NOT NULL,
    component_id    INT,
    asset_id INT NOT NULL,
    name            TEXT,
    layer_id        TEXT,
    text            TEXT,
    xi              INT,
    xii             INT,
    yi              INT,
    yii             INT,
    inner_xi              INT,
    inner_xii             INT,
    inner_yi              INT,
    inner_yii             INT,
    width           INT,
    height          INT,
    is_group        BOOL,
    group_id        INT,
    level           INT,
    kind            TEXT,
    image_url       TEXT,
    image_extension text,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP,
    CONSTRAINT fk_design_element_design_id FOREIGN KEY (design_id) REFERENCES design (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_design_element_component_id FOREIGN KEY (component_id) REFERENCES layout_components (id),
    CONSTRAINT fk_design_element_design_asset_id FOREIGN KEY (asset_id) REFERENCES design_assets (id),
    FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TYPE RENDERER_JOB_STATUS AS ENUM (
  'pending',
  'started',
  'rendering_images',
  'finished',
  'error',
  'unknown'
);

CREATE TABLE renderer_jobs (
  id BIGSERIAL PRIMARY KEY,
  layout_id INT,
  image_id INT,
  adaptation_id INT,
  status RENDERER_JOB_STATUS NOT NULL DEFAULT 'pending',
  attempts INT NOT NULL DEFAULT 0,
  started_at TIMESTAMP,
  finished_at TIMESTAMP,
  error_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  log TEXT,
  FOREIGN KEY (image_id) REFERENCES images (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (adaptation_id) REFERENCES adaptation_batch (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TYPE REPLICATION_BATCH_STATUS AS ENUM (
  'pending',
  'started',
  'rendering_images',
  'finished',
  'error',
  'closed',
  'unknown'
);

CREATE TABLE replication_batch (
  id   BIGSERIAL PRIMARY KEY,
  layout_id INT,
  design_id INT,
  user_id INT,
  status REPLICATION_BATCH_STATUS,
  started_at TIMESTAMP,
  finished_at TIMESTAMP,
  error_at TIMESTAMP,
  stopped_at TIMESTAMP,
  updated_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  config TEXT,
  log TEXT,

  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);


