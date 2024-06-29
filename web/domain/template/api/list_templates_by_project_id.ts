import { apiClient } from "../../../infrastructure/api";
import { Template } from "../entities/template";

export type Response = {
  page: number;
  limit: number;
  total: number;
  data: Array<{
    id: number;
    name: string;
    width: number;
    height: number;
    distortion: {};
    created_at: string;
  }>;
};

export function listTemplatesByProjectID(
  id: number,
): Promise<{ data: Template[] }> {
  return apiClient
    .get<Response>(`/v1/project/${id}/templates`)
    .then((res) => {
      const templates: Template[] = [];
      for (const t of res.data.data) {
        templates.push(Template.create(t));
      }
      return { data: templates };
    })
    .catch((e) => {
      throw e;
    });
}
