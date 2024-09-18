"use client"; 

import { Button } from "@/components/ui/button";

const LogoutButton = () => {
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
    <Button
      onClick={logout}
      variant={"ghost"}
      className="text-base text-gray-600 font-normal"
    >
      Logout
    </Button>
  );
};

export default LogoutButton;
