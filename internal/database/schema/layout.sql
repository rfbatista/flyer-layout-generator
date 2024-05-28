CREATE TABLE layout (
  id   BIGSERIAL PRIMARY KEY,
  width INT,
  height INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP
);

CREATE TABLE layout_components (
  id   BIGSERIAL PRIMARY KEY,
  design_id  INT NOT NULL,
  layout_id  INT NOT NULL,
  width      INT,
  height     INT,
  color      TEXT,
  type       TEXT,
  xi         INT,
  xii        INT,
  yi         INT,
  yii        INT,
  bbox_xi         INT,
  bbox_xii        INT,
  bbox_yi         INT,
  bbox_yii        INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (design_id) REFERENCES design (id) ON DELETE CASCADE ON UPDATE CASCADE,
  FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE layout_template (
  id   BIGSERIAL PRIMARY KEY,
  layout_id  INT NOT NULL,
  type TEXT,
  width INT,
  height INT,
  slots_x INT,
  slots_y INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE layout_region (
  id   BIGSERIAL PRIMARY KEY,
  layout_id  INT NOT NULL,
  xi         INT,
  xii        INT,
  yi         INT,
  yii        INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (layout_id) REFERENCES layout (id) ON DELETE CASCADE ON UPDATE CASCADE
);

