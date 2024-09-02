-- name: CreateDesignAsset :one
INSERT INTO design_assets (project_id, alternative_to, asset_url, asset_path, design_id, type, name, width, height, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP)
RETURNING *;

-- name: CreateDesignAssetProperty :exec
INSERT INTO design_assets_properties (asset_id, key, value, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP);

-- name: GetDesignAssetByProjectID :many
SELECT *
FROM design_assets
WHERE project_id = $1;

-- name: GetDesignAssetPropertyByAssetID :many
SELECT *
FROM design_assets_properties
WHERE asset_id = $1;

-- name: GetDesignAssetByDesignID :many
SELECT *
FROM design_assets
WHERE design_id = $1;


-- name: GetDesignAssetByID :one
SELECT *
FROM design_assets
WHERE id = $1
LIMIT 1
;
