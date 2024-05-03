import { api } from "../../infra/api";

export type CreateTemplateAPIResult = {
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

export function createTemplateAPI(data: any): Promise<any> {
  return api
    .post<CreateTemplateAPIResult>(`/api/v1/template`, data)
    .then((res) => {
      return;
    });
}
