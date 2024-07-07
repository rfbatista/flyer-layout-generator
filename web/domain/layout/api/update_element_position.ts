import { apiClient } from "../../../infrastructure/api"

type Response = {

}

export function updateElementPositionAPI(layoutId: number, elementId: number, payload: any) {
  return apiClient.patch(`/v1/layout/${layoutId}/element/${elementId}/position`, payload)
}
