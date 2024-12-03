import React, { Children } from "react"
import { useState,useEffect } from "react";
import { Link, Navigate } from "react-router-dom";
import  SearchIcon from '@mui/icons-material/Search';
import Brightness4Icon from '@mui/icons-material/Brightness4'; // Moon Icon for Dark Mode
import Brightness7Icon from '@mui/icons-material/Brightness7';
import MenuIcon from '@mui/icons-material/Menu';
import { Button } from "@mui/material";
import { useContext,createContext } from "react";
import { Darkcontext } from "../../App";
import axiosurl from "../../axios";
import { useNavigate } from "react-router-dom";

function Header(){
  const navigate=useNavigate()

  // useEffect(
  //  async ()=>{
  //   try{

  //this  api is going to  give us  the title  all dsa's there 
  //    const value= await  axiosurl.get('our api')
      // here the value.data  should be   object 
     //   const data=JSON.parse(value.data)
     //
  // we will the value for searching purpose
  // Setobject(data)

  //   }
  //   catch(error){
  //     console.log(error);
  //   }
  //   }
  //   ,[]
  // )
  const [fetcheddata,Setfetchdata]=useState()
   const [toggleMenu, setToggleMenu] = useState(false);
   const {show,Setshow}=useState(false)
   const {dark,toggler}  = useContext(Darkcontext)
   const [searchitem,Setsearchietm]=useState('')
   const items=[]
   const [filtered,Setfiltered]= useState([])
   const [listappear,Setlistappear]=useState(false)
   const content=createContext()
   const [object,Setobject]=useState([ {
      name:"naola",
      id:1329/14
    },
    {
      name:"misgana",
      id:1324/14
    },
    {
      name:"sudeyis",
      id:1524/14
    }
   ]
  )


   const handlesolutions=async (element)=>{
    try{
      navigate(`/solution/?element=${element}`)
 }
    catch(error){
      console.error("here is the error",error)
    }
  }
    const shower=()=>{
       Setshow(!show)
    
    }
    const handleinput=(e)=>{
      const value=e.target.value
      Setlistappear(true)
      Setsearchietm(value)
      const filteredlist =object.filter((object)=>(object.name.toLowerCase().includes(searchitem.toLowerCase())))
      Setfiltered(filteredlist)

    }
    const handlevariable =()=>{
      const newvalue=!listappear
      Setlistappear(newvalue)
    }

    return(
    <>
   
  
  <div className={`hidden md:block   bg-[#2F3136] flex  h-20 rounded-sm items-center ${ dark? 'bg-[#2F3136] text-white ': ' bg-[#ffff] text-black  '}px-4`}>
        <ul className="flex space-x-8   items-center justify-around w-full">
        <li className="text-lg">DSA SHARING</li>
        <div className="flex space-x-8 justify-center">
          <li><Link to="/" >home</Link></li>
          <li><Link to ="">About</Link></li>
          <li onClick={handlevariable}><SearchIcon/>
          <input
          type="text"
          value={searchitem}
          onChange={handleinput}
          placeholder="type the title of dsa"
          className="text-black w-[20vw] "
          >
          </input>
          { listappear && <ul  className="bg-white  h-[35vh]  w-[50vw] fixed z-[100]">
          {
            filtered.map((element,index)=>(
              <li className="text-black pl-5" key={index}  onclick={handlesolutions}>
                {element.name}
              </li>
            ))
          }
          
          </ul>
}
          </li>
          
       
          </div>
          {dark ? <button onClick={toggler}><li className={`${ dark? 'text-white': 'text-black'}`}><Brightness4Icon/></li></button>:<button onClick={toggler}><li className={`${ dark? 'text-white': 'text-black'}`}><Brightness7Icon/></li></button>
  }

<div className={`flex justify-center items-center mt-3   w-[125px] h-[35px] rounded-2 ${dark?'bg-[#7DC400]':'bg-[#4F46E5]'}`}>
            <Link to="/signin"><p className=" text-xl text-white  text-center">sign in</p></Link>
          </div>

   </ul>
      </div>

  <div className={`md:hidden flex bg-[#2F3136] h-20 rounded-sm items-center ${ dark? 'bg-[#4F46E5] text-white ': ' bg-[#ffff] text-black  '}px-4`}>
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