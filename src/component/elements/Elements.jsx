import React from 'react'
import axiosurl from '../../axios'
import { useLocation} from 'react-router-dom'

async function Elements() {
    const location=useLocation();
    const params= new URLSearchParams(location);
    const element= params.get('element')
    try{
        const response = await axiosurl.get(`element/soution/${element}`)
        // the data might contain text ,image and videos  and the response  will contain the  data we will use 
        // it is json file with type,file
        const data= JSON.parse(response.data)

    }
    catch(error){
        console.error("error",error);

    }
    

  return (
    <div>
      
    </div>
  )
}

export default Elements
