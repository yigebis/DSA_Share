import React from 'react'
import ShareIcon from '@mui/icons-material/Share';
import VideoLibraryIcon from '@mui/icons-material/VideoLibrary';
import MenuBookIcon from '@mui/icons-material/MenuBook';
import EditIcon from '@mui/icons-material/Edit';
import { Link } from 'react-router-dom';


function Dashboard() {
  return (
    <div className="flex flex-col    justify-center items-center flex-grow bg-black min-h-screen ">
      <h1 className="text-4xl text-white mb-4">what is included</h1>
      <p className="text-lg text-white">here you  can access rated videos and blogs  from students and also you can share your thought </p>
     <div className= "flex  space-x-6 flex-grow justify-center items-center ">
      <div className= " flex  w-5/6 h-64 bg-[#2F3136] border border-white justify-center items-center">
      <Link to="/videoupload">
     <button className="justify-center items-center ">
    <ShareIcon className=" text-white bg-[#2F3136] margin-auto "/>
 </button>
      </Link>
      <p className="text-xl text-white">share your idea</p>
      </div>
      
      <div className= " flex  w-5/6 h-64 bg-[#2F3136] border border-white justify-center items-center">
     <button className="justify-center items-center ">
     <VideoLibraryIcon className="text-white text-4xl"/>
      </button>
      <p className="text-xl text-white">watch videos</p>
      </div>
   </div>

   <div className= "flex  space-x-6 flex-grow justify-center items-center ">
      <div className= " flex  w-5/6 h-64 bg-[#2F3136] border border-white justify-center items-center">
     <button className="justify-center items-center ">
      <MenuBookIcon className=" text-white bg-[#2F3136] margin-auto "/>
      </button>
      <p className="text-xl text-white">read blogs</p>
      </div>
      
      <div className= " flex  w-5/6 h-64 bg-[#2F3136] border border-white justify-center items-center">
     <button className="justify-center items-center ">
     <EditIcon className="text-white text-4xl"/>
      </button>
      <p className="text-xl text-white">edit your lectures</p>
      </div>
      
      
      
      
   </div>
    </div>
  )
}

export default Dashboard
