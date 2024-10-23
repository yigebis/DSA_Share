// import React from 'react'
// import { useEffect } from 'react'
// import { useState } from 'react'
// import axiosurl from '../../axios'

// async function Watchvideos() {
// const [videos,Setvideos]=useState[[]]
// try{
//  useEffect(
//  ()=>{
// const data= axiosurl.get('/video/data')
// Setvideos(data.response)
// },[]
// )
// }
// catch(error){}

//   return (
// //     <div>
// //     <ul>
// //     {/* {videos.map((value,i)=>(
// //         <li key={i}>
// //         <h1></h1>
// //         <video>

// //         </video>
// //         </li>
    
// // ))} */}
// // </ul>
// //     </div>
//   )
// }

// export default Watchvideos
// 1. Using Spread Operator (...)
// This is a common and clean way to add elements to an array state.

// Concept:

// Create a new array by spreading the existing elements and appending the new one(s).
// Example Flow:

// jsx
// Copy code
// const addElement = (newElement) => {
//   setArrayState(prevArray => [...prevArray, newElement]);
// };