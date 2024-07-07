import { useState } from "react";
import { useDesignsStore } from "../../../domain/design/store";
import {
  generateImageAPI,
  generateImageAPIV2,
} from "../../../domain/layout/api/generate_image";
import { useTemplatesStore } from "../../../domain/template/store";
import { Image } from "../../../components/image/image";
import { useGenerationStore } from "./generation_store";

type Props = {};

export function TryBoard() {
  const { activeDesign } = useDesignsStore();
  const { templates } = useTemplatesStore();
  const { priorities } = useGenerationStore();
  const [templateId, setTemplate] = useState(0);
  const [imageURL, setImagURL] = useState("");
  const createImage = (e: any) => {
    e.preventDefault();
    if (!activeDesign) return;

    generateImageAPI(activeDesign.id, activeDesign.layoutID, templateId, {
      priorities: priorities.map((p) => p.text),
    })
      .then((r) => {
        setImagURL(r.data.image_url);
      })
      .catch((e) => {
        console.error(e);
      });

    // generateImageAPIV2(activeDesign.layoutID, templateId, {
    //   priorities: priorities.map((p) => p.text),
    //   padding: 15,
    // })
    //   .then((r) => {
    //     setImagURL(r.image_url);
    //   })
    //   .catch((e) => {
    //     console.error(e);
    //   });
  };
  return (
    <>
      <div className="box">
        <form onSubmit={createImage}>
          <div className="stack">
            <div>
              <fieldset data-type="select" className="w-10">
                <label htmlFor="os">Client</label>
                <span className="arrow" />
                <select
                  id="os"
                  name="client_id"
                  value={templateId}
                  onChange={(e) => setTemplate(Number(e.target.value))}
                >
                  <option selected></option>
                  {templates.map((c) => {
                    return <option value={c.id}>{c.name}</option>;
                  })}
                </select>
              </fieldset>
            </div>
            <div>
              <button className="my-2" type="submit">
                <div
                  className="ld ld-loader box-loader"
                  data-src="/web/assets/box.gif"
                ></div>
                Generate
              </button>
            </div>
          </div>
        </form>
      </div>
      <div className="box">
        <div className="center">
          {imageURL !== "" && <Image source={imageURL} />}
        </div>
      </div>
    </>
  );
}
