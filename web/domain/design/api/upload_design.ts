import { apiClient } from "../../../infrastructure/api";

type Response = {};

export function uploadDesignAPI(form: FormData) {
  return apiClient
    .post<Response>("/v1/design", form)
    .then((s) => {
      return s.data;
    })
    .catch((e) => {
      console.error(e);
    });
}
