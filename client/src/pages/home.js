import React, { Fragment, useState } from "react"

import Input from "../components/input";


const Home = () => {

    const [itemName, setItemName] = useState("");
    const [itemPrice, setItemPrice] = useState("");
    const [alert, setAlert] = useState('')

    function handleName(evt) {
        setItemName(evt.target.value)
    }

    function handlePrice(evt) {
        setItemPrice(evt.target.value)
    }

    function handleSubmit(evt) {
        evt.preventDefault();
        const payload = JSON.stringify({
            "name": itemName,
            "price": itemPrice
        })

        const requestOption = {
            method: "POST",
            header: {
                "Content-Type": "application/json"
            },
            body: payload
        }
        console.log(requestOption)

        fetch(`${process.env.REACT_APP_BACKEND}/v1/add`, requestOption)
            .then((res) => res.json())
            .then(data => {
                if (data.status === "OK") {
                    console.log("OK")
                    setAlert("Added successfully")
                } else {
                    console.log("ERROR")
                    setAlert("Failed")
                }
            })

    }

    return (
        <div className="container mt-3">
            {alert.length > 0 && (
                <div className="alert alert-primary" role="alert">
                    {alert}
                </div>
            )}

            <Fragment>
                <div className="h1">
                    Add Record
                </div>
                <form onSubmit={handleSubmit}>
                    <Input
                        title={"Record Name"}
                        type={"text"}
                        name={"recordName"}
                        handleChange={handleName}
                        placeholder={'Record name'}
                    />

                    <Input
                        title={"Record Price"}
                        type={"number"}
                        name={"price"}
                        handleChange={handlePrice}
                        placeholder={'Price'}
                    />
                    <button className="btn btn-primary">Submit</button>
                </form>
            </Fragment>


        </div>
    );
}

export default Home;