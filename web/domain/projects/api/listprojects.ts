import { apiClient } from "../../../infrastructure/api";
import { Project } from "../entities/project";

type Response = {
  page: number;
  limit: number;
  data: Array<{
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
    use_ai: boolean;
    briefing: string;
    name: string;
    created_at: string;
    updated_at: string;
    deleteed_at: string;
  }>;
};

export function listProjectsAPI(
  page?: number,
  limit?: number,
): Promise<{ data: Project[] }> {
  return apiClient
    .get<Response>(`/v1/projects?page=${page}&limit=${limit}`)
    .then((res) => {
      const projects: Project[] = [];
      for (const p of res.data.data) {
        projects.push(Project.create(p));
      }
      return { data: projects };
    })
    .catch((e) => {
      throw e;
    });
}
