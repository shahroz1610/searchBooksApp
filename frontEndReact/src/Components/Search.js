import React, {useState} from 'react'

function Search({fetchBooks}) {
    const [search, setSearch] = useState("");
    return (
        <div className='search'>
            <input className='search-box'
                value={search}
                placeholder="Search Books"
                onChange={(e) => {
                    setSearch(e.target.value);
                }}
            />
            {console.log(search, "search")}
            <button className='search-button' onClick = {() => fetchBooks(search)}> Search </button>
        </div>
    )
}

export default Search