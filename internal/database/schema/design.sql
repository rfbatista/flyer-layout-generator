
CREATE TABLE request (
                         id   SERIAL PRIMARY KEY,
                         name text      NOT NULL,
                         started_at TIMESTAMP,
                         finished_at TIMESTAMP,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP,
                         deleted_at TIMESTAMP
);

CREATE TABLE request_steps (
                               id   SERIAL PRIMARY KEY,
                               name text NOT NULL,
                               request_id  INT      NOT NULL,
                               started_at TIMESTAMP,
                               finished_at TIMESTAMP,
                               error_at TIMESTAMP,
                               log text,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP
);



CREATE TABLE design
(
    id              SERIAL PRIMARY KEY,
    name            TEXT NOT NULL,
    request_id INT,
    image_url       text,
    image_extension text,
    file_url        text,
    file_extension  text,
    width           int,
    height          int,
    is_proccessed   bool DEFAULT false,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP,

    FOREIGN KEY (request_id) REFERENCES request (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TYPE COMPONENT_TYPE AS ENUM (
    'background',
    'logotipo_marca',
    'logotipo_produto',
    'texto_cta'
    );

CREATE TABLE design_components
(
    id         SERIAL PRIMARY KEY,
    design_id  INT NOT NULL,
    width      INT,
    height     INT,
    color      TEXT,
    type       COMPONENT_TYPE,

    xi         INT,
    xii        INT,
    yi         INT,
    yii        INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (design_id) REFERENCES design (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE design_element
(
    id              SERIAL PRIMARY KEY,
    design_id       INT NOT NULL,
    name            TEXT,
    layer_id        TEXT,
    text            TEXT,
    xi              INT,
    xii             INT,
    yi              INT,
    yii             INT,
    width           INT,
    height          INT,
    is_group        BOOL,
    group_id        INT,
    level           INT,
    kind            TEXT,
    component_id    INT,
    image_url       TEXT,
    image_extension text,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP,
    CONSTRAINT fk_design_element_design_id FOREIGN KEY (design_id) REFERENCES design (id) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_design_element_component_id FOREIGN KEY (component_id) REFERENCES design_components (id)
);


