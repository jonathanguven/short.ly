'use client'

import { useState } from "react"

import { Copy, Check } from 'lucide-react';
import { Card, CardContent, CardFooter, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"
import { Label } from "@/components/ui/label"

export default function URLShortenerForm() {
  const [URL, setURL] = useState('');
  const [alias, setAlias] = useState('');
  const [response, setResponse] = useState('');
  const [error, setError] = useState(false);
  const [copied, setCopied] = useState(false);

  const handleSubmit = () => {
    if (!URL) {
      setResponse("Error: URL is required.");
      setError(true);
      return;
    }

    const shortLink = `https://shrink.lol/s/${alias || "generated-alias"}`;
    setResponse(shortLink);
    setError(false);
    setCopied(false);
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
          <CardDescription>Enter a long URL to get a short, shareable link.</CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="space-y-2">
            <Label htmlFor="url">URL</Label>
            <Input 
              id="url" 
              value={URL}
              placeholder="https://example.com/very/long/url"
              onChange={(e) => setURL(e.target.value)}
            />
          </div>
          <div className="space-y-2">
            <Label htmlFor="alias">Custom Alias</Label>
            <Input 
              id="alias" 
              placeholder="Enter a custom alias (optional)" 
              value={alias} 
              onChange={(e) => setAlias(e.target.value)} 
            />
          </div>
        </CardContent>
        <CardFooter className="flex justify-between">
          <Button variant="outline" onClick={handleClear}>Clear</Button>
          <Button onClick={handleSubmit}>Shrink</Button>
        </CardFooter>
      </Card>
      <div className="relative w-full">
        <Textarea 
          value={response} 
          placeholder="Shortened URL will appear here, click to copy"
          className={`resize-none h-28 justify-center pt-10 text-center text-3xl cursor-pointer ${error ? 'text-red-500' : 'text-black'}`} 
          readOnly 
          onClick={handleCopy}
        />  
        {response && !error && (
          <div 
            className="absolute right-3 top-3 cursor-pointer flex items-center justify-center"
            onClick={handleCopy}
          >
            {copied ? (
              <Check />
            ) : (
              <Copy />
            )}
          </div>
        )}
      </div>
    </div>
  )
}
