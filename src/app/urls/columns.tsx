'use client'
import React from "react";
import Link from "next/link";
import { CellContext, ColumnDef } from "@tanstack/react-table";
import { ArrowUpDown, MoreHorizontal } from "lucide-react"
import { useToast } from "@/hooks/use-toast";

import { Button } from "@/components/ui/button"
import { Checkbox } from "@/components/ui/checkbox";
import { Input } from "@/components/ui/input";

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Label } from "@/components/ui/label"


export type URL = {
  ID: number;
  Alias: string;
  Link: string;
  URL: string;
  CreatedAt: string;
  ExpiresAt?: string | null;
  ClickCount: number;
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
      const alias = row.original.Alias;
  
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
            href={`${process.env.NEXT_PUBLIC_DOMAIN}/s/${alias}`} 
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
          className="text-left ml-4 font-medium truncate md:max-w-52 lg:max-w-md"
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
    cell: ({ row, table }: CellContext<URL, unknown>) => {
      const originalAlias = row.original.Alias;
      const originalURL = row.original.URL;
      const [alias, setAlias] = React.useState(originalAlias);
      const [url, setUrl] = React.useState(originalURL);
      const { toast } = useToast();
      const [isOpen, setIsOpen] = React.useState(false);
  
      const handleSubmit = async () => {
        try {
          const response = await fetch(
            `${process.env.NEXT_PUBLIC_BASE_URL}/urls/${originalAlias}`, 
            {
              method: 'PUT',
              headers: {
                'Content-Type': 'application/json',
              },
              credentials: 'include',
              body: JSON.stringify({
                new_alias: alias,
                new_url: url,
              }),
            }
          );
  
          if (response.ok) {
            toast({
              title: "Success!",
              description: "URL updated successfully.",
            });
            setIsOpen(false);
            const updateRow = (table.options.meta as any)?.updateRow;
            if (updateRow) {
              updateRow({
                ...row.original,
                Alias: alias,
                URL: url,
              });
            }
          } else {
            throw new Error("Failed to update URL.");
          }
        } catch (error) {
          toast({
            title: "Error",
            description: "An error occurred while updating the URL.",
          });
        }
      };
  
      return (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
          <DialogTrigger asChild>
            <Button variant="ghost" className="h-8 w-8 p-0">
              <span className="sr-only">Open menu</span>
              <MoreHorizontal className="h-4 w-4" />
            </Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[425px]">
            <DialogHeader>
              <DialogTitle>Edit URL</DialogTitle>
              <DialogDescription>
                Make changes to your shortened URL here.
              </DialogDescription>
            </DialogHeader>
            <div className="grid gap-4 py-4">
              <div className="grid grid-cols-4 items-center gap-4">
                <Label htmlFor="alias" className="text-right">
                  Alias
                </Label>
                <Input
                  id="alias"
                  value={alias}
                  onChange={(e) => setAlias(e.target.value)}
                  className="col-span-3"
                />
              </div>
              <div className="grid grid-cols-4 items-center gap-4">
                <Label htmlFor="url" className="text-right">
                  URL
                </Label>
                <Input
                  id="url"
                  value={url}
                  onChange={(e) => setUrl(e.target.value)}
                  className="col-span-3"
                />
              </div>
            </div>
            <DialogFooter>
              <Button type="button" onClick={handleSubmit}>
                Save changes
              </Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      );
    },
  },
]