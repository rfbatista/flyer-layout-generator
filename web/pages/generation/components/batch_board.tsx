import { useState, useEffect } from "react";
import Masonry from "react-masonry-css";
import { LayoutCard } from "../../../components/layoutcard/layout_card";
import { useDesignsStore } from "../../../domain/design/store";
import { getLayoutByIdAPI } from "../../../domain/layout/api/get_layout_by_id";
import { Layout } from "../../../domain/layout/entities/layout";
import { useLayoutRequestStore } from "../../../domain/layout/layout_request_store";
import { useTemplatesStore } from "../../../domain/template/store";
import { apiClient } from "../../../infrastructure/api";
import { useGenerationStore } from "./generation_store";
import './batch-board.css'

export function BatchBoard() {
  const { activeDesign } = useDesignsStore();
  const { getJobs, jobs, request } = useLayoutRequestStore();
  const { listTemplatesByProjectID, templates } = useTemplatesStore();
  const { priorities } = useGenerationStore();
  const [isLoading, setLoading] = useState(false);
  const [layouts, setLayouts] = useState<Layout[]>([]);

  useEffect(() => {
    const init = async () => {
      if (!activeDesign) return;
      try {
        setLoading(true);
        await getJobs(activeDesign.id);
        await listTemplatesByProjectID(activeDesign.projectId);
        setLoading(false);
      } catch (e) {
        console.error(e);
        setLoading(false);
      }
    };
    init();
  }, [activeDesign]);

  useEffect(() => {
    if (!activeDesign) return;
    const source = new EventSource("/sse");
    source.onmessage = async function (event) {
      console.log(event);
      if (event.data == "JOB_BATCH_UPDATE") {
        await getJobs(activeDesign.id);
      }
    };
    return () => source.close();
  }, [activeDesign]);

  useEffect(() => {
    const init = async () => {
      try {
        if (!activeDesign) return;
        const layouts: Layout[] = [];
        for (const j of jobs) {
          if (j.layoutID) {
            const l = await getLayoutByIdAPI(j.layoutID);
            layouts.push(l);
          }
        }
        setLayouts(layouts);
      } catch (e) {
        console.error(e);
      }
    };
    init();
  }, [jobs]);

  const startGeneration = async () => {
    if (!activeDesign || !activeDesign.layout) return;
    try {
      setLoading(true);
      await apiClient.post(
        `/v1/project/design/${activeDesign.id}/layout/${activeDesign.layoutID}/generate`,
        {
          design_id: activeDesign.id,
          layout_id: activeDesign.layoutID,
          minimium_component_size: 50,
          minimium_text_size: 20,
          templates: templates.map((t) => t.id),
          padding: 10,
          priority: priorities.map((p) => p.text),
        },
      );
      await getJobs(activeDesign.id);
      setLoading(false);
    } catch (e) {
      console.error(e);
      setLoading(false);
    }
  };

  return (
    <>
      <div className="box">
        <div className="cluster">
          <div className="p-1">
            <button
              className="my-2"
              data-state={isLoading || request?.isDoing ? "loading" : ""}
              onClick={startGeneration}
            >
              <div
                className="ld ld-loader box-loader"
                data-src="/web/assets/box.gif"
              ></div>
              Start Batch
            </button>
            <h1>{request && `${request.done}/${request.total}`}</h1>
          </div>
        </div>
      </div>
      <Masonry
        breakpointCols={3}
        className="my-masonry-grid"
        columnClassName="my-masonry-grid_column"
      >
        {layouts.map((l) => {
          return <LayoutCard layout={l} />;
        })}
      </Masonry>
    </>
  );
}
