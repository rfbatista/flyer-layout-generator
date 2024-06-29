import { apiClient } from "../../../infrastructure/api";

type Response = {};

export function uploadTemplateCSVAPI(f: FormData) {
  return apiClient
    .post<Response>(`/v1/project/${f.get("project_id")}/templates`, f)
    .then(() => {
      return;
    })
    .catch((e) => {
      console.error(e);
      throw e;
    });
}
