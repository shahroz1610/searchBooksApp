import React, { useState, useEffect } from 'react'
import axios from "axios"
import Row from './Row';
import Search from './Search';

function Table() {
    const [books, setBooks] = useState([]);

    const fetchBooks = async (bookName) => {
        console.log("BookName", bookName)
        const response = await axios.get("http://localhost:8090/getBooks", {
            params: {
                "bookName": bookName
            }
        }).catch(err => console.log(err))
        if (response) {
            console.log("Response received", response.data)
            var data = response.data["items"]
            var book = []
            for (let i = 0; i < data.length; i++) {
                var temp = {}
                for (var key in data[i]["volumeInfo"]) {
                    if (key == "authors" && data[i]["volumeInfo"][key]!=null) {
                        var s = ""
                        console.log(data[i]["volumeInfo"][key])
                        for (let j = 0; j < data[i]["volumeInfo"][key].length; j++) {
                            s += data[i]["volumeInfo"][key][j]
                            if (j < data[i]["volumeInfo"][key].length - 1) {
                                s += ", "
                            }
                        }
                        temp[key] = s
                    } else {
                        temp[key] = data[i]["volumeInfo"][key]
                    }
                }
                book.push(temp)
            }
            setBooks(book)
            console.log("first", books)
        }
    }

    function handleDelete(val) {
        const idx = books.indexOf(val)
        if (idx > -1) {
            books.splice(idx, 1)
            const newBooks = books.slice()
            setBooks(newBooks)
        }
    }

    return (
        <div className='outer-div'>
            <Search fetchBooks = {fetchBooks} />
           <table id='books'>
                <tr>
                    <th> Title </th>
                    <th> Pages </th>
                    <th> Authors </th>
                    <th> Language </th>
                    <th> Remove </th>
                </tr>

                {books.map((val, key) => {
                    return (
                        <Row val = {val} key = {key} handleDelete = {handleDelete} />
                    )
                })}
           </table>
        </div>
    )
}

export default Table