import { useEffect } from "react";
import "./App.css";
import { Editor } from "./components/editor";
import { useDesignsStore } from "../../domain/design/store";
import { EditorPanel } from "./components/editor_panel";
import { TopBar } from "../../components/topbar/topbar";
import { useProjectsStore } from "../../domain/projects/store";

export default function App() {
  const { setActiveProject, activeProject } = useProjectsStore();
  const { setActiveDesign, activeDesign } = useDesignsStore();
  useEffect(() => {
    let params = new URLSearchParams(document.location.search);
    let id = params.get("design");
    let projectId = params.get("project");
    setActiveProject(Number(projectId));
    setActiveDesign(Number(id));
  }, []);
  return (
    <>
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
            title: "edit",
            link: "",
          },
        ]}
      />
      <main className="right-sidebar">
        <div>
          <div className="stack">
            <Editor />
          </div>
          <div className="box">
            <EditorPanel />
          </div>
        </div>
      </main>
    </>
  );
}
