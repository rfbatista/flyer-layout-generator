import { PhotoshopFile } from "../../entities/photoshop";
import { api } from "../../infra/api";

export type PhotoshopListAPIResult = {
  status: string;
  data: {
    id: number;
    name: string;
    image_url: any;
    file_url: string;
    width?: number;
    height?: number;
    created_at: string;
    updated_at: any;
  }[];
};

export function getPhotoshopListAPI(): Promise<PhotoshopFile[]> {
  return api.get<PhotoshopListAPIResult>("/api/v1/photoshop").then((res) => {
    const data = PhotoshopFile.from_api_list(res.data);
    return data;
  });
}
