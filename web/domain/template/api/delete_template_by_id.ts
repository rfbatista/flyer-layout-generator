import { apiClient } from "../../../infrastructure/api";

type Response = {};

export function deleteTemplateByID(projectID: number, templateId: number) {
  return apiClient
    .delete<Response>(`/v1/project/${projectID}/templates/${templateId}`)
    .then((s) => {
      return s.data;
    });
}
