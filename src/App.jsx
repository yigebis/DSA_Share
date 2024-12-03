import { BrowserRouter as Router,Routes,Route, RouterProvider } from "react-router-dom"
import Home from "./pages/home/home"
import Signup2 from "./pages/signup/signup"
import Signin from "./pages/signin/signin"
import Signin2 from "./pages/signin/signin"
import Footer from "./component/footer/footer"
import Dashboardpage from "./pages/dashboard/dashboardpage"
import { createContext,useState } from "react"
import Share from "./pages/share/share"
import Elements from "./component/elements/elements"

const Darkcontext=createContext();

function   DarkMode  ({children}){
 const [dark,setdark]=useState(true);

 const toggler=()=>{
  if(dark){
    setdark(!dark)
  }
  else{
    setdark(true)
  }
 }
 return (
   <Darkcontext.Provider value={{dark , toggler}}>
    {children}
   </Darkcontext.Provider>
 )
}

function App() {
    // const [toggleMenu, setToggleMenu] = useState(false);
  return (
    <>
    <DarkMode>
    <Router>
      <Routes>
        <Route  path="/" element={<Home/>}></Route>
        <Route path="/signup" element={<Signup2/>}></Route>
        <Route path="/signin" element={<Signin2/>}></Route>
        <Route path="/about" element={<Footer/>}></Route>
        <Route path="/videoupload"   element={<Share/>}></Route>
        <Route path="/dashboard"  element={<Dashboardpage/>}></Route>
        <Route path="/soultion"    element={<Shareolutions/>}></Route>
        <Route path="/soultion/element" element={Elements}></Route>

        {/* <Route  path="/mycourses"  element={}></Route> */}
      </Routes>
     </Router>
    </DarkMode>
    </>

  );
}
export  {App,Darkcontext}
