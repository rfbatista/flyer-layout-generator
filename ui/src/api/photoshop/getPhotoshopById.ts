import { PhotoshopFile } from "../../entities/photoshop";
import { api } from "../../infra/api";

export type GetPhotoshopByIDAPIResult = {
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
  };
};

export function getPhotoshopByIDAPI(id: string): Promise<PhotoshopFile> {
  return api.get<GetPhotoshopByIDAPIResult>(`/api/v1/photoshop/${id}`).then((res) => {
    const data = new PhotoshopFile(res.data.data);
    return data;
  });
}
