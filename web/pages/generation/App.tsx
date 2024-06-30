import { useEffect } from "react";
import { Modal } from "../../components/modal/modal";
import { TopBar } from "../../components/topbar/topbar";
import { useProjectsStore } from "../../domain/projects/store";
import "./App.css";
import { Breadcrump } from "../../components/breadcrump/breadcrump";

export default function App() {
  const { setActiveProject, activeProject } = useProjectsStore();
  useEffect(() => {
    let params = new URLSearchParams(document.location.search);
    let id = params.get("project");
    setActiveProject(Number(id));
  }, []);
  return (
    <>
      <Modal title="New design"></Modal>
      <TopBar />
      <div>
        <div className="container">
          <div>
            <div className="stack">
              <div>
                <Breadcrump
                  items={[
                    { title: "projects", link: "/" },
                    {
                      title: activeProject ? activeProject.name : "",
                      link: `/project?id=${activeProject?.id}`,
                    },
                    {
                      title: "generate",
                      link: "",
                    },
                  ]}
                />
                <div className="box">
                  <button> Start </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
