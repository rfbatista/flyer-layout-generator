-- name: GetComponentByID :one
select pc.* from design_components pc
where pc.id = $1 LIMIT 1;

-- name: HaveElementsIn :many
select pc.* from design_components pc
inner join design_element as pe on pe.component_id = pc.id 
where pc.id = $1;

-- name: ListComponentByFileId :many
select pc.* from design_components pc
where pc.design_id = $1;

-- name: UpdateManydesignElement :many
UPDATE design_element
SET 
-- You can use sqlc.arg() and @ to identify named parameters
    component_id = CASE WHEN @component_id_do_update::boolean
        THEN @component_id::int ELSE component_id END,

    name = CASE WHEN @name_do_update::boolean
        THEN sqlc.narg(name) ELSE name END
WHERE
    id = ANY (sqlc.slice(ids)) and design_id = @design_id
RETURNING *;


-- name: RemoveComponentFromElements :many
UPDATE design_element
SET 
    component_id = NULL
WHERE
    id = ANY (sqlc.slice(ids)) and design_id = @design_id
RETURNING *;

-- name: CreateComponent :one
INSERT INTO design_components (
  design_id,
  width,
  height,
  xi,
  xii,
  yi,
  yii,
  type,
  color
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;


