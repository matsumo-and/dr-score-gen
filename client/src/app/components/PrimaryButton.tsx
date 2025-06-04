import Link from "next/link";
import React from "react";

type PrimaryButtonProps = {
  title: string;
  href?: string;
  onClick?: () => void;
  className?: string;
  type?: "button" | "submit" | "reset";
};

export default function PrimaryButton({
  title,
  href,
  onClick,
  className = "",
  type = "button",
}: PrimaryButtonProps) {
  const baseClass =
    "bg-gradient-to-r from-indigo-500 to-blue-500 text-white px-8 py-3 rounded-full text-lg font-semibold shadow-lg hover:scale-105 transition " +
    className;

  if (href) {
    return (
      <Link href={href} className={baseClass}>
        {title}
      </Link>
    );
  }
  return (
    <button type={type} className={baseClass} onClick={onClick}>
      {title}
    </button>
  );
} 