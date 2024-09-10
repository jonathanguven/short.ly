'use client'

import { useState } from "react"

import { Card, CardContent, CardFooter, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"
import { Label } from "@/components/ui/label"

export default function URLShortenerForm() {
  const [URL, setURL] = useState('');
  const [alias, setAlias] = useState('');
  const [response, setResponse] = useState('');

  const handleSubmit = () => {
    if (!URL) {
      setResponse("Error: URL is required.");
      return;
    }

    const shortLink = `https://shrink.lol/s/${alias || "generated-alias"}`;
    setResponse(shortLink);
  };

  // Handle clear button click
  const handleClear = () => {
    setURL("");
    setAlias("");
    setResponse("");
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
      <Textarea 
        value={response} 
        placeholder="Shortened URL will appear here"
        className="resize-none h-28 justify-center pt-10 text-center text-3xl" 
        readOnly 
      />
    </div>
  )
}
