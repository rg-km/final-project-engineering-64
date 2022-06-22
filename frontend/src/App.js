
import Navbar from './component/NavBar';
//import Header from './pages/Header';
import Content from './pages/Content';
import About from './pages/About';
import UploadBook from "./pages/UploadBook";
import Register from "./pages/Register";
import Login from "./pages/Login";
import { Route, Routes } from "react-router-dom";
import DescriptionBook from "./pages/DescriptionBook";

function App() {
  return (
    <main aria-label="App" className="App">
      <div className="routes" aria-label="routes">
        <Navbar />
        <Routes>
          <Route path="/" element={<Content />} />
          <Route path="/uploadbook" element={<UploadBook />} />
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/about" element={<About />} />
          <Route path="/description" element={<DescriptionBook />} />
        </Routes>
      </div>
    </main>
  );
}

export default App;
