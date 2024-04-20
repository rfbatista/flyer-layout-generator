import { ElementTree } from "../../entities/photoshop";
import { usePhotoshopFiles } from "../../store/files";
import { FileIcon } from "../Icons/File";
import { Directory } from "./Directory";
import { Item } from "./Item";

type Props = {
  tree?: ElementTree;
};
export const TreeView: React.FC<Props> = ({ tree }) => {
  return (
    <>
      <ul
        style={{ borderLeftColor: `black`, borderLeftWidth: 2 }}
        className="p-2 pt-0 ml-2 mb-0 mt-0 pb-0 menu bg-default text-content-700"
      >
        {tree &&
          tree.children &&
          tree.children.map((item, index) => {
            if (item.isDir)
              return <Directory key={item.element?.id} item={item} />;
            return <Item key={item.element?.id} item={item} />;
          })}
      </ul>
    </>
  );
};
