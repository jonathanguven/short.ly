'use client'

import Link from "next/link";
import { ColumnDef } from "@tanstack/react-table";
import { ArrowUpDown, MoreHorizontal } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Checkbox } from "@/components/ui/checkbox";
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
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && "indeterminate")
        }
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
      />
    ),
    cell: ({ row }) => {
      const alias = row.original.alias;
  
      return (
        <div>
          <Checkbox
            checked={row.getIsSelected()}
            onCheckedChange={(value) => row.toggleSelected(!!value)}
            aria-label={`Select row for alias ${alias}`}
          />
        </div>
      );
    },
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "Alias",
    header: () => <div className="text-left ml-4">Alias</div>,
    cell: (info) => {
      const alias = info.getValue() as string
      return (
        <div className="ml-4">
          <Link 
            href={`${process.env.NEXT_PUBLIC_BASE_URL}/s/${alias}`} 
            target="_blank" 
            className="font-medium hover:underline"
          >
            {alias}
          </Link>
        </div>
      )
    }
  },
  {
    accessorKey: "URL",
    header: () => <div className="text-left ml-4">URL</div>,
    cell: (info) => {
      const originalURL = info.getValue() as string
      return (
        <div
          className="text-left ml-4 font-medium truncate md:max-w-56 lg:max-w-md"
        >
          {originalURL}
        </div>
      );
    },
  },
  {
    accessorKey: "ClickCount",
    header: ({ column }) => {
      return (
        <Button
          variant="ghost"
          className="font-medium text-normal"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Clicks
          <ArrowUpDown className="ml-2 h-4 w-4" />
        </Button>
      )
    },
    cell: (info) => {
      const count = info.getValue() as number
      return (
        <div>
          <div className="font-medium ml-10">{count}</div>
        </div>
      )
    },
  },
  {
    accessorKey: "CreatedAt",
    header: ({ column }) => {
      return (
        <Button
          variant="ghost"
          className="font-medium text-normal"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Date Created
          <ArrowUpDown className="ml-2 h-4 w-4" />
        </Button>
      )
    },
    cell: info => {
      const value = info.getValue() as string;
      const date = new Date(value);
      const formatted = date.toLocaleDateString('en-US', {
        month: 'numeric',
        day: 'numeric',
        year: 'numeric',
      });

      return <div className="text-left font-medium ml-8">{formatted}</div>
    },
  },
  {
    id: "actions",
    cell: ({ row }) => {
      const payment = row.original
 
      return (
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="ghost" className="h-8 w-8 p-0">
              <span className="sr-only">Open menu</span>
              <MoreHorizontal className="h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuItem>Edit URL</DropdownMenuItem>
            <DropdownMenuItem>Edit Alias</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      )
    },
  },
]