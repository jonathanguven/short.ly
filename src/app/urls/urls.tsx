'use client'

import { useState } from 'react';
import { DataTable } from './data-table';
import { columns } from './columns';

type URL = {
  ID: number;
  Alias: string;
  Link: string;
  URL: string;
  CreatedAt: string;
  ExpiresAt?: string | null;
  ClickCount: number;
}

interface URLsClientProps {
  initialUrls: URL[];
}

export default function URLsClient({ initialUrls }: URLsClientProps) {
  const [urls, setUrls] = useState<URL[]>(initialUrls);

  const updateRow = (updatedRow: URL) => {
    setUrls((prevUrls) =>
      prevUrls.map((url) => (url.ID === updatedRow.ID ? updatedRow : url))
    );
  };

  return <DataTable columns={columns} data={urls} updateRow={updateRow} />;
}