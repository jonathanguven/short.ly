"use client"; 

import { ChevronDown } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"

interface LogoutButtonProps {
  user: string | null;
}

const LogoutButton = ({ user }: LogoutButtonProps) => {
  const logout = async () => {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_BASE_URL}/logout`, {
        method: "POST",
        credentials: "include", 
      });

      if (response.ok) {
        const data = await response.text();
        console.log(data); 

        window.location.href = "/";
      } else {
        console.error("Failed to log out");
      }
    } catch (error) {
      console.error("Error occurred during logout:", error);
    }
  };

  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="ghost" className="ml-auto font-normal text-base">
            {user} <ChevronDown className="ml-2 h-4 w-4" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent>
          <Button
            onClick={logout}
            variant={"ghost"}
            className="w-full text-base text-gray-600 font-normal"
          >
            Logout
          </Button>
        </DropdownMenuContent>
      </DropdownMenu>
    </>
  );
};

export default LogoutButton;
