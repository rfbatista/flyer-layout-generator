import { useEffect } from "react";
import { Modal } from "../../components/modal/modal";
import { TopBar } from "../../components/topbar/topbar";
import { useProjectsStore } from "../../domain/projects/store";
import "./App.css";
import { LeftMenu } from "./components/left_menu";
import { GenerationBoard } from "./components/generation_board";
import { useDesignsStore } from "../../domain/design/store";
import { useTemplatesStore } from "../../domain/template/store";

export default function App() {
  const { setActiveProject, activeProject } = useProjectsStore();
  const { setActiveDesign, activeDesign } = useDesignsStore();
  const { listTemplatesByProjectID } = useTemplatesStore();
  useEffect(() => {
    let params = new URLSearchParams(document.location.search);
    let id = params.get("project");
    setActiveProject(Number(id));
    listTemplatesByProjectID(Number(id));
    let desingId = params.get("design");
    setActiveDesign(Number(desingId));
  }, []);
  return (
    <>
      <Modal title="New design"></Modal>
      <TopBar
        items={[
          { title: "projects", link: "/" },
          {
            title: activeProject ? activeProject.name : "",
            link: `/project?id=${activeProject?.id}`,
          },
          {
            title: activeDesign ? activeDesign.name : "",
            link: `/project?id=${activeProject?.id}`,
          },
          {
            title: "generate",
            link: "",
          },
        ]}
      />
      <div className="with-sidebar">
        <div>
          <div className="box">
            <LeftMenu />
          </div>
          <div className="box">
            <div>
              <div className="stack">
                <div>
                  <div className="box">
                    <GenerationBoard />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
