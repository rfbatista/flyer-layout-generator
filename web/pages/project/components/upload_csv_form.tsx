import { useState } from "react";
import { useTemplatesStore } from "../../../domain/template/store";
import { useModal } from "../../../components/modal/store";
import { uploadTemplateCSVAPI } from "../../../domain/template/api/upload_template_csv";
import { useProjectsStore } from "../../../domain/projects/store";

export function UploadCSVForm() {
  const [isLoading, setLoading] = useState(false);
  const { close } = useModal();
  const { listTemplatesByProjectID } = useTemplatesStore();
  const { activeProject } = useProjectsStore();
  const [file, setFile] = useState<any>();
  const onSubmit = (e: any) => {
    e.preventDefault();
    console.log(e.target.value, file);
    const data = new FormData();
    data.set("file", file);
    activeProject && data.set("project_id", String(activeProject?.id));
    setLoading(true);
    uploadTemplateCSVAPI(data)
      .then(() => {
        setLoading(false);
        activeProject && listTemplatesByProjectID(activeProject?.id);
      })
      .catch((e) => {
        setLoading(false);
        console.error(e);
      });
  };
  return (
    <form className="stack" onSubmit={onSubmit}>
      <div className="cluster">
        <div>
          <fieldset>
            <label htmlFor="file">File</label>
            <input
              data-size="md"
              name="file"
              type="file"
              onChange={(event) => {
                event.target.files && setFile(event.target.files[0]);
              }}
            />
          </fieldset>
          <div>
            <button type="submit" data-state={isLoading && "loading"}>
              <div className="ld ld-ring ld-spin"></div>
              Create
            </button>
          </div>
          <div>
            <button data-type="outline" type="button" onClick={close}>
              Close
            </button>
          </div>
        </div>
      </div>
    </form>
  );
}
