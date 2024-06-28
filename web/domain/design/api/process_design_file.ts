import { apiClient } from "../../../infrastructure/api";

type Response = {};

export function processDesignFileAPI(id: number) {
  return apiClient
    .post<Response>(`/v1/design/${id}/process`)
    .then((r) => {
      console.log(r.data);
    })
    .catch(console.error);
}
