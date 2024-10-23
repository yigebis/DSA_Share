import React from "react"
import Footer from "../footer/footer"
import Header from "../header/header"
function Layout({children}){

    return (
    <>
    <div className="flex flex-col min-h-screen">
    <div className="flex-grow">
    <Header/>
    {children}
    <Footer/>
    </div>
    </div>
    </>
    )
}
export default Layout