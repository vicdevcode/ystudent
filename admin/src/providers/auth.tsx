import {
  Dispatch,
  FC,
  PropsWithChildren,
  SetStateAction,
  createContext,
  useContext,
  useEffect,
  useState,
} from "react";
import { useNavigate } from "react-router-dom";

interface AuthContext {
  isAuth: boolean;
  setAuth: Dispatch<SetStateAction<boolean>>;
  token: string;
  setToken: Dispatch<SetStateAction<string>>;
  logout: () => void;
  refresh: () => void;
}

const defaultValues: AuthContext = {
  isAuth: false,
  setAuth: () => null,
  token: "",
  setToken: () => null,
  logout: () => null,
  refresh: () => null,
};

const Context = createContext(defaultValues);

interface AuthProvidersProps extends PropsWithChildren { }

export const AuthProvider: FC<AuthProvidersProps> = (props) => {
  const { children } = props;
  const navigate = useNavigate();

  const [isAuth, setAuth] = useState(defaultValues.isAuth);
  const [token, setToken] = useState(defaultValues.token);

  const logout = () => {
    localStorage.removeItem("token");
    setToken("");
    setAuth(false);
    navigate("/auth");
  };

  const refresh = async () => {
    return fetch(import.meta.env.VITE_AUTH_API + "/refresh-token", {
      method: "GET",
    }).then(async (res) => {
      const json = await res.json();
      if (res.status !== 200) {
        setToken("");
        setAuth(false);
        navigate("/auth");
        return;
      }
      setToken(json["access_token"]);
      setAuth(true);
    });
  };

  useEffect(() => {
    if (window) {
      const t = localStorage.getItem("token");
      if (t) {
        setToken(t);
        setAuth(true);
        return;
      }
    }
    setToken("");
    setAuth(false);
    navigate("/auth");
  }, [token, isAuth, navigate]);

  const exposed = {
    isAuth,
    setAuth,
    token,
    setToken,
    logout,
    refresh,
  };

  return <Context.Provider value={exposed}>{children}</Context.Provider>;
};

export const useAuth = () => useContext(Context);
