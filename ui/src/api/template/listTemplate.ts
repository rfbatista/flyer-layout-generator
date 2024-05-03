import { Template } from "../../entities/template";
import { api } from "../../infra/api";

export type ListTemplateAPIResult = {
  status: string;
  data: {
    id: number;
    name: string;
    image_url: any;
    file_url: string;
    width?: number;
    height?: number;
    created_at: string;
    updated_at: any;
  };
};

export function listTemplateAPI(limit = 10, skip = 0): Promise<Template[]> {
  return api
    .get<ListTemplateAPIResult>(`/api/v1/template?limit=${limit}&skip=${skip}`)
    .then((res) => {
      return Template.fromApiList(res);
    });
}
