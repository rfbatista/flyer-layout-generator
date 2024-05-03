import { FileIcon, Image, Puzzle } from "lucide-react";
import { ElementTree } from "../../entities/photoshop";
import { Spinner } from "flowbite-react";
import { usePhotoshopStore } from "../../store/photoshop";

export const Item: React.FC<{ item: ElementTree }> = ({ item }) => {
  const { isSelected, selectItem, isLoading, component , components} = usePhotoshopStore(
    (d) => ({
      selectItem: d.selectElement,
      isSelected:
        item.element?.id && d.elementsSelected.includes(item.element?.id),
      isLoading: d.isLoading,
      component:
        item.element?.componentId &&
        d.components.find((c) => c.id() === item.element?.componentId),
      components: d.components,
    }),
  );
  return (
    <li className="cursor-pointer">
      <span
        className={`hover:bg-gray-100 transition flex flex-row gap-3  p-2 truncate ${isSelected && "bg-indigo-100"}`}
        onClick={() => item.id && selectItem(item.id)}
      >
        <div className="w-5 h-5">
          {item.isComponent ? (
            isLoading && isSelected ? (
              <Spinner />
            ) : (
              <Puzzle color={component && component.color()} />
            )
          ) : isLoading && isSelected ? (
            <Spinner />
          ) : item.isBackground ? (
            <Image />
          ) : (
            <FileIcon />
          )}
        </div>
        {item.name}
      </span>
    </li>
  );
};
