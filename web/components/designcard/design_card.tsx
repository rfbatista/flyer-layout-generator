import { useState } from "react";
import { Design } from "../../domain/design/entities/design";
import { useDesignsStore } from "../../domain/design/store";
import "./design_card.css";

type Props = {
  design: Design;
};

export function DesginCard(props: Props) {
  const d = props.design;
  const { processDesignFile, listDesigns } = useDesignsStore();
  const [isLoading, setLoading] = useState(false);
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
        <div className="design-card__body">
          <h2>{d.name}</h2>
        </div>
        <div className="cluster center design-card__body">
          {d.isProcessed ? (
            <div>
              <a href={`/editor?design=${d.id}`} data-type="button">
                Edit
              </a>
              <button>Generate</button>
            </div>
          ) : (
            <div>
              <button
                data-state={isLoading && "loading"}
                onClick={() => {
                  setLoading(true);
                  processDesignFile(d.id).then(() => {
                    setLoading(false);
                    listDesigns(d.projectId);
                  });
                }}
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
