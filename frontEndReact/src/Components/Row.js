import React from 'react'

function Row({val, key, handleDelete}) {

    const langMap = {"en" : "English"}
    console.log(val, "val")

    return (
        <tr key =  {key}>
            <td> {val.title} </td>
            <td> {val.pageCount} </td>
            <td> {val.authors} </td>
            <td> {langMap[val.language]} </td>
            <td> <button onClick={() => handleDelete(val)} className="del-button"> Delete </button> </td>
        </tr>
    )
}

export default Row