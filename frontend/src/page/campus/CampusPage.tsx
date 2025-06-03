import CampusGrid from "@/components/campus/CampusGrid";
import CampusHero from "@/components/campus/CampusHero";
import JoinCampusDialog, {
  JoinCampusRef,
} from "@/components/dialog/JoinCampusDialog";
import LoginModal, { LoginModalRef } from "@/components/dialog/LoginDialog";
import { useRef } from "react";

export default function CampusPage() {
  const loginModalRef = useRef<LoginModalRef>(null);
  const joinCampusDialogRef = useRef<JoinCampusRef>(null);

  const handleUnauthorized = () => {
    if (loginModalRef.current) {
      loginModalRef.current.open();
    }
  };

  const onNoCampusUserClick = () => {
    if (joinCampusDialogRef.current) {
      joinCampusDialogRef.current.open();
    }
  };

  return (
    <>
      <CampusHero
        onUnauthorized={handleUnauthorized}
        onNoCampusUserClick={onNoCampusUserClick}
      />
      <CampusGrid />

      <LoginModal ref={loginModalRef} />
      <JoinCampusDialog ref={joinCampusDialogRef} />
    </>
  );
}
