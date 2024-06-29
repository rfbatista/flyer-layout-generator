import React, { ReactNode, useEffect, useRef } from "react";
import "./modal.css";
import { useModal } from "./store";

export function Modal({
  children,
  title = "",
}: {
  children?: ReactNode;
  title: string;
}) {
  const modalRef = useRef<HTMLDialogElement | null>(null);
  const { isOpen, close, ch, title: modalTitle } = useModal();

  const handleKeyDown = (event: React.KeyboardEvent<HTMLDialogElement>) => {
    if (event.key === "Escape") {
      close();
    }
  };

  useEffect(() => {
    const modalElement = modalRef.current;
    if (modalElement) {
      if (isOpen) {
        modalElement.showModal();
      } else {
        modalElement.close();
      }
    }
  }, [isOpen]);
  return (
    <dialog
      data-size="md"
      ref={modalRef}
      onKeyDown={handleKeyDown}
      className="modal"
    >
      <div className="modal__header">
        <h3
          id="defaultModalTitle"
          className="font-semibold tracking-wide text-black dark:text-white"
        >
          {modalTitle !== "" ? modalTitle : title}
        </h3>
        <span className="modal__close" onClick={close} />
      </div>
      <div className="modal__body">{ch ? ch : children}</div>
    </dialog>
  );
}
