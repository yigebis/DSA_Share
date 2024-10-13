import React from "react"
import { useState } from "react";
import { Link } from "react-router-dom";
import  SearchIcon from '@mui/icons-material/Search';
import Brightness4Icon from '@mui/icons-material/Brightness4'; // Moon Icon for Dark Mode
import Brightness7Icon from '@mui/icons-material/Brightness7';
import MenuIcon from '@mui/icons-material/Menu';

import { Button } from "@mui/material";
function Header(){
  // const [toggleMenu, setToggleMenu] = useState(false);
     const [dark,setdark]=useState(true)
     const [show,Setshow]=useState(false)
     const toggler=()=>{
      if(dark){
        setdark(!dark)
      }
      else{
        setdark(true)
      }

     }
    const shower=()=>{
       Setshow(!show)
    
    }

    return(
    <>
   
  
  <div className={`hidden md:block   bg-[#2F3136] flex  h-20 rounded-sm items-center ${ dark? 'bg-[#2F3136] text-white ': ' bg-[#ffff] text-black  '}px-4`}>
        <ul className="flex space-x-8   items-center justify-around w-full">
        <li className="text-lg">DSA SHARING</li>
        <div className="flex space-x-8 justify-center">
          <li><Link to="/" >home</Link></li>
          <li><Link to ="/signup">About</Link></li>
          <li><SearchIcon/></li>
          
       
          </div>
          {dark ? <button onClick={toggler}><li className={`${ dark? 'text-white': 'text-black'}`}><Brightness4Icon/></li></button>:<button onClick={toggler}><li className={`${ dark? 'text-white': 'text-black'}`}><Brightness7Icon/></li></button>
  }

<div className=" flex justify-center items-center mt-3   w-[125px] h-[35px] rounded-2 bg-[#7DC400]">
            <Link to="/signin"><p className=" text-xl text-white  text-center">sign in</p></Link>
          </div>

   </ul>
      </div>

  <div className={`md:hidden flex bg-[#2F3136] h-20 rounded-sm items-center ${ dark? 'bg-[#2F3136] text-white ': ' bg-[#ffff] text-black  '}px-4`}>
  <ul className="flex space-x-4   items-center justify-around w-full">
        <li className="text-lg">DSA SHARING</li>
        <li><button onClick={shower}><MenuIcon/> </button></li>
   </ul>
  </div>
  {
  show && 
  <div className="bg-black text-white">
  <ul className="list-none">
  <li className="style-none"><Link to="/" >home</Link></li>
  <li><Link to ="/signup">About</Link></li>
  <li><Link to ="#" >Services</Link></li>
  <li><Link to="#" >Contact</Link></li>
  </ul>
  </div>
  }
      </>

    )
}
export default  Header