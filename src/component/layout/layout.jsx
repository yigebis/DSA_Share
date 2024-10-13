import React from "react"
import Footer from "../footer/footer"
import Header from "../header/header"
function Layout({children}){

    return (
    <>
    <Header/>
    {children}
    <Footer/>
    </>
    )
}
export default Layout