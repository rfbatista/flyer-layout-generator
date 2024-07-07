import { useEffect } from "react";
import { useProjectsStore } from "../../../domain/projects/store";
import "./project_table.css";
import "../../../components/link/link.css";

export function ProjectsTable() {
  const { projects, listProjects } = useProjectsStore();
  useEffect(() => {
    listProjects();
  }, []);
  return (
    <div className="table-container">
      <table className="project-table">
        <thead>
          <tr>
            <th scope="col" className="project-table__header__item">
              Name
            </th>
            <th scope="col" className="project-table__header__item">
              Client
            </th>
            <th scope="col" className="project-table__header__item">
              Advertiser
            </th>
            <th scope="col" className="project-table__header__item">
              Created at
            </th>
            <th scope="col" className="project-table__header__item"></th>
          </tr>
        </thead>
        <tbody>
          {projects.map((p) => {
            return (
              <tr>
                <td className="project-table__body__item">{p.name}</td>
                <td className="project-table__body__item">{p.clientName}</td>
                <td className="project-table__body__item">
                  {p.advertiserName}
                </td>
                <td className="project-table__body__item">{p.createdAtText}</td>
                <td className="project-table__body__item">
                  <div className="cluster">
                    <div className="p-2">
                      <a href={`/project?id=${p.id}`} data-type="button">
                        Open
                      </a>
                      <button data-type="outline" data-color="danger">
                        Delete
                      </button>
                    </div>
                  </div>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}
