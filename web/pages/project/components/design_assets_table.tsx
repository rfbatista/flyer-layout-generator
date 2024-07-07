import { useEffect, useState } from "react";
import { DesignAssetType } from "../../../domain/design/entities/design_asset";
import { useDesignsStore } from "../../../domain/design/store";
import { useProjectsStore } from "../../../domain/projects/store";
import "./design_assets_table.css";
import { useModal } from "../../../components/modal/store";
import { addNewPropertyToAssetAPI } from "../../../domain/design/api/add_new_asset";

function CreatePropModal() {
  const [isLoading, setLoading] = useState(false);
  const { activeAsset } = useDesignsStore();
  const { close } = useModal();
  const [text, setText] = useState("");
  const onSubmit = (e: any) => {
    e.preventDefault();
    if (!activeAsset) return;
    setLoading(true);
    addNewPropertyToAssetAPI(activeAsset, text)
      .then(() => {
        setLoading(false);
      })
      .catch((e) => {
        setLoading(false);
        console.error(e);
      });
  };
  return (
    <form className="stack" onSubmit={onSubmit}>
      <fieldset>
        <label htmlFor="file">Text</label>
        <input
          data-size="md"
          name="filename"
          value={text}
          onChange={(e) => setText(e.target.value)}
        />
      </fieldset>
      <div className="cluster">
        <div>
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

export function DesignAssetsTable() {
  const { designAssets } = useProjectsStore();
  const { getDesignById, setActiveAsset } = useDesignsStore();
  const { setCh, open, setTitle } = useModal();
  const [tableData, setTableData] = useState<
    {
      id: number;
      design: string;
      content: string;
      fonts: string[];
      fontSize: string;
    }[]
  >([]);
  useEffect(() => {
    designAssets
      .filter((d) => d.type === DesignAssetType.TEXT)
      .map(async (p) => {
        const design = await getDesignById(p.desinID);
        setTableData([
          ...tableData,
          {
            id: p.id,
            design: design.name,
            content: p.text || "",
            fonts: p.fonts,
            fontSize: p.fontSize || "",
          },
        ]);
      });
  }, [designAssets]);

  const openModal = (id: number) => {
    setActiveAsset(id);
    setCh(<CreatePropModal />);
    setTitle("Add a copy");
    open();
  };
  return (
    <>
      <div className="table-container">
        <table className="design-table">
          <thead>
            <tr>
              <th scope="col" className="design-table__header__item">
                Design
              </th>
              <th scope="col" className="design-table__header__item">
                Content
              </th>
              <th scope="col" className="design-table__header__item">
                Fonts
              </th>
              <th scope="col" className="design-table__header__item">
                Size
              </th>
              <th scope="col" className="design-table__header__item"></th>
            </tr>
          </thead>
          <tbody>
            {tableData.map((p) => {
              return (
                <tr>
                  <td className="design-table__body__item">{p.design}</td>
                  <td className="design-table__body__item">{p.content}</td>
                  <td className="design-table__body__item">
                    <div className="cluster">
                      {p.fonts.map((f) => {
                        return (
                          <div>
                            <span data-type="badge">{f}</span>
                          </div>
                        );
                      })}
                    </div>
                  </td>
                  <td className="design-table__body__item">{p.fontSize}</td>
                  <td className="design-table__body__item">
                    <button
                      onClick={() => openModal(p.id)}
                      type="button"
                      data-type="icon"
                      data-icon="magic"
                    ></button>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    </>
  );
}
