'use client'

import Link from "next/link";
import { ColumnDef } from "@tanstack/react-table";
import { MoreHorizontal } from "lucide-react"

import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

export type URL = {
  id: number;
  alias: string;
  link: string;
  url: string;
  createdAt: string;
  expiresAt?: string | null;
  clickCount: number;
}

export const columns: ColumnDef<URL>[] = [
  {
    accessorKey: "Alias",
    header: () => <div className="text-center">Alias</div>,
    cell: (info) => {
      const alias = info.getValue() as string
      return (
        <div className="flex justify-center">
          <div className="font-medium">{alias}</div>
        </div>
      )
    }
  },
  {
    accessorKey: "URL",
    header: () => <div className="text-left ml-4">URL</div>,
    cell: (info) => {
      const originalURL = info.getValue() as string
      return <Link href={originalURL} target="_blank" className="text-center ml-4 font-medium hover:underline">{originalURL}</Link>
    },
  },
  {
    accessorKey: "ClickCount",
    header: () => <div className="text-center font-medium">Click Count</div>,
    cell: (info) => {
      const count = info.getValue() as number
      return (
        <div className="flex justify-center">
          <div className="font-medium">{count}</div>
        </div>
      )
    },
  },
  {
    accessorKey: "CreatedAt",
    header: () => <div className="text-right">Date Created</div>,
    cell: info => {
      const value = info.getValue() as string;
      const date = new Date(value);
      const formatted = date.toLocaleDateString('en-US', {
        month: 'numeric',
        day: 'numeric',
        year: 'numeric',
      });

      return <div className="text-right font-medium">{formatted}</div>
    },
  },
]