-- name: CreateAPICredential :exec
INSERT INTO companies_api_credentials (name, api_key, company_id)
VALUES ($1, $2, $3);

-- name: ListAPICredetialsByCompany :many
SELECT * FROM companies_api_credentials WHERE deleted_at IS NULL AND company_id = $1;

-- name: DeleteAPICredetial :exec 
DELETE FROM companies_api_credentials
WHERE id = $1;

