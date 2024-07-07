import { useEffect } from "react";
import { useTemplatesStore } from "../../../domain/template/store";
import { useProjectsStore } from "../../../domain/projects/store";
import "../../../components/table/table.css";
import "../../../components/icons/icons.css";
import "../../../components/badge/badge.css";
import "../../../components/pagination/pagination.css";
import { useModal } from "../../../components/modal/store";
import { UploadCSVForm } from "./upload_csv_form";
import { Briefing } from "./briefing";

export function SideBar() {
  const { activeProject } = useProjectsStore();
  const { listTemplatesByProjectID, templates, deleteTemplate, isLoading } =
    useTemplatesStore();
  const { setCh, open, setTitle } = useModal();
  useEffect(() => {
    activeProject && listTemplatesByProjectID(activeProject?.id);
  }, [activeProject]);

  const openModal = () => {
    setCh(<UploadCSVForm />);
    setTitle("Upload CSV");
    open();
  };

  return (
    <div className="stack">
      <div>
        <div className="box">
          <div className="table-container">
            <table>
              <thead>
                <tr>
                  <th scope="col">Name</th>
                  <th scope="col">Size</th>
                  <th scope="col"></th>
                </tr>
              </thead>
              <tbody>
                {templates.map((t) => {
                  return (
                    <tr>
                      <td>
                        <span
                          style={{
                            maxWidth: "220px",
                            width: "220px",
                            textOverflow: "ellipsis",
                          }}
                        >
                          {t.name}
                        </span>
                      </td>
                      <td>
                        <span data-type="badge">
                          {t.width}x{t.height}
                        </span>
                      </td>
                      <td>
                        <button
                          data-type="icon"
                          data-state={isLoading && "loading"}
                          data-color="blue"
                          onClick={async () => {
                            await new Promise((r) => setTimeout(r, 5000));
                            activeProject &&
                              deleteTemplate(activeProject.id, t.id);
                          }}
                        >
                          <div className="ld ld-ring ld-spin"></div>
                          <i className="fa-solid fa-xmark"></i>
                        </button>
                      </td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>
          <div className="center"></div>
        </div>
        <div className="box">
          <div className="cluster">
            <div>
              <button onClick={openModal}>CSV</button>
            </div>
          </div>
        </div>
        <div className="box">
          <Briefing />
        </div>
      </div>
    </div>
  );
}
