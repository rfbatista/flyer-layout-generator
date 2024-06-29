import { apiClient } from "../../../infrastructure/api";
import { Project } from "../entities/project";

type Response = {
  data: {
    id: number;
    client: {
      id: number;
      name: string;
      created_at: string;
      updated_at: string;
      deleteed_at: string;
    };
    advertiser: {
      id: number;
      name: string;
      created_at: string;
      updated_at: string;
      deleteed_at: string;
    };
    name: string;
    use_ai: boolean;
    briefing: string;
    created_at: string;
    updated_at: string;
    deleteed_at: string;
  };
};

export function updateProjectByIdAPI(id: number, form: FormData): Promise<{ data: Project }> {
  return apiClient
    .patch<Response>(`/v1/project/${id}`, form)
    .then((res) => {
      return { data: Project.create(res.data.data) };
    })
    .catch((e) => {
      throw e;
    });
}
