import { apiClient } from "../../../infrastructure/api";
import { Design } from "../entities/design";

type Response = {
  data: {
    id: number;
    project_id: number;
    name: string;
    filepath: string;
    layout_id: number;
    image_path: string;
    image_url: string;
    width: number;
    height: number;
    created_at: string;
    is_processed: boolean;
  };
};

export function getDesignByIDApi(id: number): Promise<{ data: Design }> {
  return apiClient
    .get<Response>(`/v1/project/${id}`)
    .then((res) => {
      return { data: Design.create(res.data.data) };
    })
    .catch((e) => {
      throw e;
    });
}
