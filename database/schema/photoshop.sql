CREATE TABLE photoshop (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  image_url text,
  file_url text,
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
  component_id TEXT,
  component_type COMPONENT_TYPE,
  image_url TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,

  FOREIGN KEY (photoshop_id) REFERENCES photoshop (id)
);
