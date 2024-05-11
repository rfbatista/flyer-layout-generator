export const appConfig = {
  api: {
    baeURL: import.meta.env.VITE_API_BASE_URL,
  },
  paths: {
    gerador: "/",
    templates: "/templates",
    createTemplates: "/templates/create",
    generateParts: "/generate",
    designs: "/designs",
  },
};
