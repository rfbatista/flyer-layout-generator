CREATE TABLE photoshop (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  image_url text,
  image_extension text,
  file_url text,
  file_extension text,
  width int,
  height int,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TYPE COMPONENT_TYPE AS ENUM (
  'background',
  'logotipo_marca',
  'logotipo_produto',
  'texto_cta'
);

CREATE TABLE photoshop_element (
  id SERIAL PRIMARY KEY,
  photoshop_id INT NOT NULL,
  name TEXT,
  layer_id TEXT,
  text TEXT,
  xi INT,
  xii INT,
  yi INT,
  yii INT,
  width INT,
  height INT,
  is_group BOOL,
  group_id INT,
  level INT,
  kind TEXT,
  component_id INT,
  image_url TEXT,
  image_extension text,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,

  CONSTRAINT fk_photoshop_element_photoshop_id FOREIGN KEY (photoshop_id) REFERENCES photoshop (id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT fk_photoshop_element_component_id FOREIGN KEY (component_id) REFERENCES photoshop (id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE TABLE photoshop_components (
  id SERIAL PRIMARY KEY,
  photoshop_id INT NOT NULL,
  width INT,
  height INT,
  color TEXT,
  type COMPONENT_TYPE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

