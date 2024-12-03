import React from 'react'
import axiosurl from '../../axios'
import { useLocation,useNavigate} from 'react-router-dom'

async function solution() {
    
const navigate= useNavigate()
const location =useLocation()
const params= new URLSearchParams(location.search);
const element= params.get('element')
try{
const response=  axios.get(`solution/${element}`)
// the data should  contain the rating and user account 
const  data=JSON.parse(response.data)
}
catch(error){
    console.error("there is error",error)
}
const  handlenext = (element)=>{
    navigate(`soultion/element/?userid=${element}`)
}
    
  return (
    <div>
    {data.map((element)=>{
        <div  onClick={handlenext(element.userid)} className="flex  h-[10vh] w-full">
            <div classname="text-white text-2xl">
                {element.rating}
            </div>
            <div>
                {element.useracount}
            </div>

        </div>

    })
    }
    </div>
  )
}

export default solution
