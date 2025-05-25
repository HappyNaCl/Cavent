import { toast } from "sonner";
import { Button } from "../ui/button";
import { useNavigate } from "react-router";

export default function LogoutButton() {
  const nav = useNavigate();

  const handleClick = async () => {
    try {

    } catch (error) {
      toast.error(`Error: ${error}`);
    }
  };
  return <Button onClick={handleClick}>Logout</Button>;
}
