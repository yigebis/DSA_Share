import React from 'react'
import Layout from '../../component/layout/layout'
import { Link } from 'react-router-dom'
import myimage  from '../../assets/dsa share.png'

function Home(){
    return(
        <>
        
          <Layout> 
          <div className="flex  flex-col flex-grow bg-black  w-full min-h-screen">
            <div>
          <h1 className="text-4xl text-white pt-10 text-center">DSA  SHARING AND LEARNING SITE</h1><br/>
          </div>
          <div>
          <p className=" text-xl  text-white pt-5 text-center"> the best  site  to share your understanding of DSA and getting understanding  from other peer a2svians</p>
          </div>
          <div className=" flex justify-center mt-10">
          <div className=" flex justify-center items-center mt-10   w-[200px] h-[50px] rounded bg-[#7DC400]">
            <Link to="/signup"><p className=" text-xl text-white  text-center">create account</p></Link>
          </div>
          </div>
          </div>
         </Layout>
         
     
        </>
    )
}
export default Home