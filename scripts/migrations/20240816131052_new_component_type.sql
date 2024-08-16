-- Add value to enum type: "component_type"
ALTER TYPE "component_type" ADD VALUE 'icone' AFTER 'grafismo';
-- Add value to enum type: "component_type"
ALTER TYPE "component_type" ADD VALUE 'contorno' AFTER 'icone';
-- Add value to enum type: "component_type"
ALTER TYPE "component_type" ADD VALUE 'titulo' AFTER 'contorno';
-- Add value to enum type: "component_type"
ALTER TYPE "component_type" ADD VALUE 'preco' AFTER 'titulo';
-- Add value to enum type: "component_type"
ALTER TYPE "component_type" ADD VALUE 'botao' AFTER 'preco';
