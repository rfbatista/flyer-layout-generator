import { apiClient } from "../../../infrastructure/api";

type ResponseV2 = {
  data: {
    image_url: string;
    elements: Array<{
      element_id: number;
      image_url: string;
    }>;
  };
};

type Response = {
  image_url: string;
};

export function generateImageAPI(
  designID: number,
  layoutID: number,
  templateID: number,
  payload?: { priorities: string[] },
) {
  return apiClient
    .post<ResponseV2>(
      `/v1/design/${designID}/layout/${layoutID}/template/${templateID}/generate`,
      {
        slots_x: 8,
        slots_y: 8,
        ...payload,
      },
    )
    .then((r) => {
      return r.data;
    });
}

export function generateImageAPIV2(
  layoutID: number,
  templateID: number,
  payload?: { priorities?: string[]; padding?: number },
) {
  return apiClient
    .post<Response>(`/v2/layout/${layoutID}/template/${templateID}/generate`, {
      slots_x: 8,
      slots_y: 8,
      ...payload,
    })
    .then((r) => {
      return r.data;
    });
}
