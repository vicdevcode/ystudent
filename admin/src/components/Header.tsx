import { useAuth } from "@/providers/auth";
import { Button } from "./ui/button";

const Header = () => {
  const { logout } = useAuth();
  return (
    <header>
      <div className="container mx-auto my-0">
        <div className="py-3 flex justify-between items-center">
          <ul className="flex gap-10">
            <li>
              <a href="/">Главная</a>
            </li>
            <li>
              <a href="/faculties">Факультеты</a>
            </li>
            <li>
              <a href="/departments">Кафедры</a>
            </li>
            <li>
              <a href="/groups">Группы</a>
            </li>
            <li>
              <a href="/users">Пользователи</a>
            </li>
          </ul>
          <Button onClick={logout}>Выйти</Button>
        </div>
      </div>
    </header>
  );
};

export default Header;
