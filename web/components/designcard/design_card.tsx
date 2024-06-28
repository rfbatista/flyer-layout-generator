import { Design } from "../../domain/design/entities/design";
import { useDesignsStore } from "../../domain/design/store";
import "./design_card.css";

type Props = {
  design: Design;
};

export function DesginCard(props: Props) {
  const d = props.design;
  const { isLoading, processDesignFile, listDesigns } = useDesignsStore();
  return (
    <article className="design-card">
      <div className="stack">
        <div className="design-card__image-container">
          <img
            src={
              d.isProcessed
                ? d.imageURL
                : "https://placehold.co/240x176?text=Need+Process"
            }
            alt="view of a coastal Mediterranean village on a hillside, with small boats in the water."
          />
        </div>
        <div>
          <h2>{d.name}</h2>
        </div>
        <div className="cluster center">
          {d.isProcessed ? (
            <div>
              <button>Edit</button>
              <button>Generate</button>
            </div>
          ) : (
            <div>
              <button
                data-state={isLoading && "loading"}
                onClick={() =>
                  processDesignFile(d.id).then(() => listDesigns(d.projectId))
                }
              >
                <span data-type="spinner" />
                Process
              </button>
            </div>
          )}
        </div>
      </div>
    </article>
  );
}
