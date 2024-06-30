import { apiClient } from "../../../infrastructure/api";
import { LayoutProps } from "../../layout/entities/layout";
import { Design } from "../entities/design";

type Response = {
  status: string;
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
    layout?: LayoutProps;
  };
};

export function getDesignByIDApi(id: number): Promise<{ data: Design }> {
  return apiClient
    .get<Response>(`/v1/design/${id}`)
    .then((res) => {
      return { data: Design.create(res.data.data) };
    })
    .catch((e) => {
      throw e;
    });
}
