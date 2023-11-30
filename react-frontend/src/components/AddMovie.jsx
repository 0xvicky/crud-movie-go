import React, {useState} from "react";

const AddMovie = () => {
  const [movie, setMovie] = useState({
    title: "",
    isbn: "",
    director: {
      firstname: "",
      lastName: ""
    }
  });

  const addMovie = async () => {
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(movie)
    };
    try {
      const newMovie = await fetch("http://localhost:8080/movies", options);
      const data = await newMovie.json();
      console.log(data);
    } catch (error) {
      console.log(`Error occured:${error}`);
    }
  };
  console.log(movie);
  return (
    <div className='flex flex-col'>
      <input
        type='text'
        placeholder='Title'
        onChange={e => {
          setMovie(res => {
            return {...res, title: e.target.value};
          });
        }}
      />
      <input
        type='text'
        placeholder='ISBN'
        onChange={e => {
          setMovie(res => {
            return {...res, isbn: e.target.value};
          });
        }}
      />
      <input
        type='text'
        placeholder='Director First Name'
        onChange={e => {
          setMovie(res => {
            return {...res, director: {...res.director, firstname: e.target.value}};
          });
        }}
      />
      <input
        type='text'
        placeholder='Director Last Name'
        onChange={e => {
          setMovie(res => {
            return {...res, director: {...res.director, lastName: e.target.value}};
          });
        }}
      />
      <button
        className='bg-black text-purple-100 p-1 rounded-md'
        onClick={addMovie}>
        Add Movie
      </button>
    </div>
  );
};

export default AddMovie;
