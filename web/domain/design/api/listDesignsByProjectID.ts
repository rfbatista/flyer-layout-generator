import { apiClient } from "../../../infrastructure/api";
import { Design } from "../entities/design";

type Response = {
  data: Array<{
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
  }>;
};

export function listDesignsByProjectID(
  id: number,
): Promise<{ data: Design[] }> {
  return apiClient
    .get<Response>(`/v1/designs/project/${id}`)
    .then((res) => {
      const designs: Design[] = [];
      for (const d of res.data.data) {
        designs.push(Design.create(d));
      }
      return { data: designs };
    })
    .catch((e) => {
      throw e;
    });
}
