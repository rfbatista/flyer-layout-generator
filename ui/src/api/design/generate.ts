import { api } from "../../infra/api";

export type GenerateDesignAPIResult = {
  status: string;
  data: {
    id: number;
    image_url: any;
    image_type: any;
  };
};

export async function generateDesignAPI(input: FormData): Promise<any> {
  return api.post<GenerateDesignAPIResult>(`/api/v1/design`, input, {
    headers: { "Content-Type": "multipart/form-data" },
  });
}
