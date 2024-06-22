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
  slots_x INT,
  slots_y INT,
  max_slots_x INT,
  max_slots_y INT,
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


CREATE TABLE design
(
    id              SERIAL PRIMARY KEY,
    name            TEXT NOT NULL,
    image_url       text,
    image_extension text,
    file_url        text,
    file_extension  text,
    width           int,
    height          int,
    is_proccessed   bool DEFAULT false,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP
);

CREATE TYPE COMPONENT_TYPE AS ENUM (
    'background',
    'logotipo_marca',
    'logotipo_produto',
    'packshot',
    'celebridade',
    'modelo',
    'ilustracao',
    'oferta',
    'texto_legal',
    'grafismo',
    'texto_cta'
    );


CREATE TABLE layout (
  id   BIGSERIAL PRIMARY KEY,
  design_id INT,
  width INT,
  height INT,
  data TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  CONSTRAINT fk_layout_design_id FOREIGN KEY (design_id) REFERENCES design (id) ON DELETE CASCADE ON UPDATE CASCADE
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
    FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE layout_requests (
  id   BIGSERIAL PRIMARY KEY,
  design_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  log TEXT,
  config TEXT,
  deleted_at TIMESTAMP,
  FOREIGN KEY (design_id) REFERENCES design (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE layout_requests_jobs (
  id   BIGSERIAL PRIMARY KEY,
  layout_id INT,
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
  FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE images (
  id   BIGSERIAL PRIMARY KEY,
  url text      NOT NULL,
  photoshop_id INT,
  template_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
