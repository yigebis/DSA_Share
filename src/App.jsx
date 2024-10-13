import { BrowserRouter as Router,Routes,Route, RouterProvider } from "react-router-dom"
import Home from "./pages/home/home"
import Signup2 from "./pages/signup/signup"
import Signin from "./pages/signin/signin"
import Signin2 from "./pages/signin/signin"

function App() {

  return (
    <>
    <Router>
      <Routes>
        <Route  path="/" element={<Home/>}></Route>
        <Route path="/signup" element={<Signup2/>}></Route>
        <Route path="/signin" element={<Signin2/>}></Route>
      </Routes>
     </Router>
  </>
  )
}

export default App
