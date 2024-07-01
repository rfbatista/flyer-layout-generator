import { apiClient } from "../../../infrastructure/api";
import { Layout } from "../entities/layout";
import { LayoutComponentProps } from "../entities/layout_component";
import { LayoutElementProps } from "../entities/layout_element";

type Response = {
  Layout: {
    id: number;
    image_url: string;
    design_id: number;
    width: number;
    height: number;
    total?: number;
    done?: number;
    components: Array<LayoutComponentProps>;
    elements: Array<LayoutElementProps>;
    template: {
      distortion: {};
      created_at: string;
    };
    grid: {
      regions: any;
      SlotsX: number;
      SlotsY: number;
    };
    stages: Array<string>;
    config: {
      grid: {
        regions: any;
        SlotsX: number;
        SlotsY: number;
      };
    };
  };
};

export function getLayoutByIdAPI(designId: number, layoutId: number) {
  return apiClient
    .get<Response>(`/v1/project/design/${designId}/layout/${layoutId}`)
    .then((r) => {
      return Layout.create(r.data.Layout);
    });
}
