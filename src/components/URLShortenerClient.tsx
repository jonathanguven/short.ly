'use client';

import { useState } from "react";
import { Copy, Check } from 'lucide-react';
import { Card, CardContent, CardFooter, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Label } from "@/components/ui/label";

interface URLShortenerClientProps {
  isAuthenticated: boolean;
}

export default function URLShortenerClient({ isAuthenticated }: URLShortenerClientProps) {
  const [URL, setURL] = useState('');
  const [alias, setAlias] = useState('');
  const [response, setResponse] = useState('');
  const [error, setError] = useState(false);
  const [copied, setCopied] = useState(false);

  const handleSubmit = async () => {
    if (!URL) {
      setResponse("URL is required.");
      setError(true);
      return;
    }
  
    try {
      const serverURL = process.env.NEXT_PUBLIC_BASE_URL;
      const res = await fetch(`${serverURL}/shorten`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify({ URL, alias }),
      });
  
      if (!res.ok) {
        const errorMessage = await res.text();
        throw new Error(errorMessage);
      }
  
      const data = await res.json();
  
      if (data.shortened_url) {
        setResponse(data.shortened_url); 
        setError(false);
        setCopied(false);
      } else {
        setResponse("Error: Failed to retrieve shortened URL.");
        setError(true);
      }
    } catch (err) {
      if (err instanceof Error) {
        setResponse(err.message);
      } else {
        setResponse("An unexpected error occurred.");
      }
      setError(true);
    }
  };

  const handleClear = () => {
    setURL("");
    setAlias("");
    setResponse("");
    setError(false);
    setCopied(false);
  };

  const handleCopy = () => {
    if (!error && response) {
      navigator.clipboard.writeText(response);
      setCopied(true);
    }
  };

  return (
    <div className="max-w-3xl w-full mx-auto space-y-6">
      <Card>
        <CardHeader>
          <CardTitle>Shrink</CardTitle>
          <CardDescription className="text-md">Enter a long URL to get a short, shareable link.</CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="url">URL</Label>
            <Input 
              id="url" 
              className="text-md"
              value={URL}
              placeholder="https://example.com/very/long/url"
              onChange={(e) => setURL(e.target.value)}
            />
          </div>
          <div className="space-y-2">
            <Label htmlFor="alias">Custom Alias</Label>
            <Input 
              id="alias" 
              {...!isAuthenticated && { disabled: true }}
              placeholder={isAuthenticated ? 'Enter a custom alias (optional)' : 'Enter a custom alias (must be logged in)'}
              value={alias} 
              maxLength={8}
              className="text-md"
              onChange={(e) => setAlias(e.target.value)} 
            />
          </div>
        </CardContent>
        <CardFooter className="flex justify-between">
          <Button variant="outline" onClick={handleClear}>Clear</Button>
          <Button onClick={handleSubmit} type="submit">Shrink</Button>
        </CardFooter>
      </Card>
      <div className="relative w-full">
        <div
          className={`text-center flex justify-center items-center overflow-x-auto whitespace-nowrap shadow-sm text-3xl rounded-lg cursor-pointer ${error ? 'text-red-500' : 'text-black'} ${response ? 'text-zinc-600' : 'text-zinc-400'} bg-white border-2`} 
          style={{
            fontSize: 'clamp(1rem, 3vw, 2rem)',
            height: 'clamp(4rem, 10vw, 6rem)',
          }}
          onClick={handleCopy}
        >
          {response || 'Shortened URL will appear here, click to copy'}
        </div>
        {response && !error && (
          <div 
            className="absolute right-3 top-3 cursor-pointer flex items-center justify-center"
            onClick={handleCopy}
          >
            {copied ? (
              <Check color="#52525b"/>
            ) : (
              <Copy color="#52525b"/>
            )}
          </div>
        )}
      </div>
    </div>
  );
}
