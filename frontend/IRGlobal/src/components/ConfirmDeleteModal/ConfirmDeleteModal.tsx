import React from "react";
import { Modal } from "../Modal/Modal";
import { Button } from "../Button/Button";

interface ConfirmDeleteModalProps {
  isOpen: boolean;
  onClose: () => void;
  onConfirm: () => void;
  title?: string;
  message?: string;
  itemName?: string;
  isLoading?: boolean;
}

export const ConfirmDeleteModal: React.FC<ConfirmDeleteModalProps> = ({
  isOpen,
  onClose,
  onConfirm,
  title = "Confirmar Exclusão",
  message,
  itemName,
  isLoading = false,
}) => {
  const defaultMessage = itemName
    ? `Tem certeza que deseja deletar "${itemName}"?`
    : "Tem certeza que deseja deletar este item?";

  return (
    <Modal isOpen={isOpen} onClose={onClose} title={title} size="sm">
      <div className="text-center">
        {/* Icon */}
        <div className="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-red-100 mb-4 transform transition-transform duration-200 hover:scale-105">
          <svg
            className="h-6 w-6 text-red-600"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
            />
          </svg>
        </div>

        {/* Message */}
        <div className="mb-6">
          <p className="text-sm text-gray-600">{message || defaultMessage}</p>
          <p className="text-xs text-gray-500 mt-2">
            Esta ação não pode ser desfeita.
          </p>
        </div>

        {/* Actions */}
        <div className="flex gap-3 justify-center">
          <Button
            variant="secondary"
            onClick={onClose}
            disabled={isLoading}
            size="sm"
          >
            Cancelar
          </Button>
          <Button
            variant="danger"
            onClick={onConfirm}
            disabled={isLoading}
            size="sm"
          >
            {isLoading ? "Deletando..." : "Deletar"}
          </Button>
        </div>
      </div>
    </Modal>
  );
};
 