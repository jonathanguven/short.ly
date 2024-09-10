import { Card, CardContent, CardFooter, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"
import { Label } from "@/components/ui/label"

export default function URLShortenerForm() {
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
            <Input id="url" placeholder="https://example.com/very/long/url"  />
          </div>
          <div className="space-y-2">
            <Label htmlFor="alias">Custom Alias</Label>
            <Input id="alias" placeholder="Enter a custom alias (optional)" />
          </div>
        </CardContent>
        <CardFooter className="flex justify-between">
          <Button variant="outline">Clear</Button>
          <Button>Shorten URL</Button>
        </CardFooter>
      </Card>
      <Textarea 
        placeholder="Shortened URL will appear here" 
        className="resize-none h-14 justify-center pt-4" 
        readOnly 
      />
    </div>
  )
}
