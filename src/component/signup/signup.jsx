import React from 'react'
import { useNavigate } from 'react-router-dom';
import { useRef } from 'react';
import { Link } from 'react-router-dom'; 
import { useContext } from 'react';
import { Darkcontext } from '../../App';
import axios from 'axios';
import axiosurl from '../../axios';

function Signup() {
    const {dark}=useContext(Darkcontext)
    const navigate=useNavigate();
    const username = useRef();
    const firstname = useRef();
    const lastname = useRef();
    const useremail = useRef();
    const userpassword = useRef();
    const  handleSubmit=(e)=>{
     e.preventDefault();
      const usernameValue = username.current.value;
      const firstnameValue = firstname.current.value;
      const lastnameValue = lastname.current.value;
      const useremailValue = useremail.current.value;
      const userpasswordValue = userpassword.current.value;
  
      if (
        !usernameValue ||
        !firstnameValue ||
        !lastnameValue ||
        !useremailValue ||
        !userpasswordValue
      ) {
        alert("Please provide all information");
  
        return;
      } 
      else {

      alert("successfully registered");
      try{
        axiosurl.post('/register',{usernameValue,
          firstnameValue,
          lastnameValue,
          useremailValue,
          userpasswordValue


        })

      }
      catch(error){
        console.log(error.message)
      }
      navigate("/");


    

        console.log("Mock registration data:", {
          username: usernameValue,
          firstname: firstnameValue,
          lastname: lastnameValue,
          email: useremailValue,
          password: userpasswordValue,
      });
        
     }

    }
  return (

    <div className={`flex flex-grow   ${dark ?  'bg-black': 'bg-white'} flex-1 flex-col justify-center px-6 py-12 lg:px-8`}>
    <div className="sm:mx-auto sm:w-full sm:max-w-sm">
      
      <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-white">
        Sign up to your account
      </h2>
    </div>

    <div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
      <form   className="space-y-6"  onSubmit={handleSubmit}>

      <div>
          <label className={`block text-sm font-medium leading-6 ${dark?'text-white':'text-black'}`}>
            user name
          </label>
          <div className="mt-2">
            <input
              id="username"
              name="username"
              type="text"
              required
              autoComplete="user name"
              ref={username}
              className={`px-4 block w-full rounded-md border-0 py-1.5 text-white shadow-sm ring-1 ring-inset ring-gray-300 placeholder: focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 ${dark?'text-white bg-black':'text-black bg-white'}`}
            />
          </div>
        </div>
        <div>
          <label  className={`block text-sm font-medium leading-6 ${dark?'text-white':'text-black'}`}>
            first name
          </label>
          <div className="mt-2">
            <input
              id="firstname"
              name="firstname"
              type="text"
              required
              autoComplete="firstname"
              ref={firstname}
              className= {`px-4 block w-full rounded-md border-0 py-1.5 text-white shadow-sm ring-1 ring-inset ring-gray-300 placeholder:focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 ${dark?'text-white bg-black':'text-black bg-white'}`}
            />
          </div>
        </div>

        <div>
          <label  className={`block text-sm font-medium leading-6 ${dark?'text-white bg-black':'text-black bg-white'}`}>
            lastname
          </label>
          <div className="mt-2">
            <input
              id="lastname"
              name="lastname"
              type="text"
              required
      
              ref={lastname}
              className={`px-4 block w-full rounded-md border-0 py-1.5 text-white shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-white focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 ${dark?'text-white bg-black':'text-black bg-white'}`}
            />
          </div>
        </div>

        <div>
          <label htmlFor="email" className={`block text-sm font-medium leading-6   ${dark? 'text-white':'text-black'}`}>
            Email address
          </label>
          <div className="mt-2">
            <input
              id="email"
              name="email"
              type="email"
              required
              autoComplete="email"
              ref={useremail}
              className={ `px-4 block w-full rounded-md border-0 py-1.5 text-white shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-white focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 ${dark?'text-white bg-black':'text-black bg-white'}`}
            />
          </div>
        </div>

        <div>
          <div className="flex items-center justify-between">
            <label htmlFor="password" className={`block text-sm font-medium leading-6 ${dark? 'text-white':'text-black'}`}>
              Password
            </label>
          </div>
          <div className="mt-2 flex items-center justify-center">
            <input
              id="password"
              name="password"
              type="password"
              required
              autoComplete="current-password"
              ref={userpassword}
              className={`px-4 block w-full rounded-md border-0 py-1.5 text-white shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-black focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6 ${dark?'text-white bg-black':'text-black bg-white'}`}
            />
          </div>
        </div>

        <div>
          <button
            type="submit"
            className={`flex w-full justify-center rounded-md  px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-#6366F1 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-#6366F1 ${ dark? 'bg-[#7DC400] text-white': 'bg-[#4F46E5] text-black'}}`}
          >
            join us
          </button>
        </div>
      </form>

     
    </div>
  </div>
)
}
export default Signup;
