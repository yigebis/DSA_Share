import React from 'react'
import { FaGithub, FaLinkedin, FaWhatsapp } from 'react-icons/fa';
import {FaCode} from 'react-icons/fa'
import { Link } from 'react-router-dom';

function Footer(){
return(
    <>
    <div className="flex  flex-col justify-center items-center absoulte bg-black bottom-0 w-full">
    <ul className="flex space-x-8 justify-center  bg-black text-white ">
    <Link to="https://github.com/naol16" target="_blank"><li className="text-2xl"><FaGithub/></li></Link>
    <Link to="https://www.linkedin.com/in/yigerem-bisrat-73bb92254/" target=""><li className="text-2xl"><FaLinkedin/></li></Link>
    <li className="text-2xl"><FaWhatsapp/></li>
    <Link to="https://leetcode.com/u/yigerem_bisrat/" target="_blank">
    <li>
    <div>
    <h1>LeetCode</h1>
    <FaCode size={30} color="green" />
    </div>
    </li>
    </Link>
    </ul>
    <p className="text-white">&copy; 2024 NY tech</p>
    </div>

    </>
)
}
export default Footer